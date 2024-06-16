package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"periph.io/x/conn/v3/gpio"
)

type WeatherResponse struct {
	Status    string `json:"status"`
	Count     string `json:"count"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Forecasts []struct {
		Province   string `json:"province"`
		City       string `json:"city"`
		Adcode     string `json:"adcode"`
		ReportTime string `json:"reporttime"`
		Casts      []struct {
			Date         string `json:"date"`
			Week         string `json:"week"`
			Dayweather   string `json:"dayweather"`
			Nightweather string `json:"nightweather"`
			Daytemp      string `json:"daytemp_float"`
			Nighttemp    string `json:"nighttemp_float"`
			Daywind      string `json:"daywind"`
			Nightwind    string `json:"nightwind"`
			Daypower     string `json:"daypower"`
			Nightpower   string `json:"nightpower"`
		} `json:"casts"`
	} `json:"forecasts"`
}

type Weather struct {
	ReportTime       time.Time
	DayTemperature   float64
	NightTemperature float64
	WindDirection    string
	WindPower        string
	Weather          string
	WaterPlanSec     int
}

// 根据气温和湿度计算浇水时长的函数
func (w *Weather) calculateWateringSeconds() {
	// 定义绣球花基础需水量（分钟）
	baseTime := 20.0

	// 如果温度太低，不进行浇水
	if w.DayTemperature < 0 || w.NightTemperature < 0 {
		w.WaterPlanSec = 0
	}

	// 如果天气包含“雨”或“雪”，不进行浇水
	if strings.Contains(w.Weather, "雨") || strings.Contains(w.Weather, "雪") {
		w.WaterPlanSec = 0
	}

	// 定义气温对需水量的影响
	// 假设每增加1摄氏度，增加1分钟
	tempAdjustment := (w.DayTemperature + w.NightTemperature - 40) / 2

	// 定义天气对需水量的影响
	weatherAdjustment := 0.0
	if strings.Contains(w.Weather, "霾") || strings.Contains(w.Weather, "雾") {
		weatherAdjustment = 5.0
	} else if strings.Contains(w.Weather, "风") {
		weatherAdjustment = 2.0
	}

	// 计算最终的浇水时长
	wateringTime := baseTime + tempAdjustment + weatherAdjustment

	// 确保浇水时间不小于基础时间，也不过长
	if wateringTime < baseTime/2.0 {
		wateringTime = baseTime / 2.0 // 设定一个下限，避免过度浇水
	} else if wateringTime > 60 {
		wateringTime = 60 // 设定一个上限，避免过度浇水
	}

	w.WaterPlanSec = int(wateringTime * 60)
}

func (r *Resolver) GetWeatherInfo() error {
	if r.weather != nil && time.Since(r.weather.ReportTime) < 1*time.Hour {
		return nil
	}
	// Get weather data from the weather API, API address: https://restapi.amap.com/v3/weather/weatherInfo
	// API parameters: key, city
	// The key is the key of the weather API, and the city is the city name
	key := "805685a37870fd471eeb75db48fb3f2b"
	city := "330106"
	baseURL := "https://restapi.amap.com/v3/weather/weatherInfo"
	params := url.Values{}
	params.Add("key", key)
	params.Add("city", city)
	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", response.StatusCode)
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return err
	}

	var weatherResponse WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		return err
	}

	if weatherResponse.Status != "1" || weatherResponse.Infocode != "10000" {
		log.Printf("Unexpected status: %s, infocode: %s", weatherResponse.Status, weatherResponse.Infocode)
		return fmt.Errorf("unexpected status: %s, infocode: %s", weatherResponse.Status, weatherResponse.Infocode)
	}

	if len(weatherResponse.Forecasts) == 0 {
		log.Println("No weather data")
		return fmt.Errorf("no weather data")
	}

	forecast := weatherResponse.Forecasts[0]

	if len(forecast.Casts) == 0 {
		log.Println("No weather data")
		return fmt.Errorf("no weather data")
	}

	live := forecast.Casts[0]

	t1, err := strconv.ParseFloat(live.Daytemp, 64)
	if err != nil {
		log.Printf("Failed to parse day temperature: %v", err)
		return err
	}

	t2, err := strconv.ParseFloat(live.Nighttemp, 64)
	if err != nil {
		log.Printf("Failed to parse night temperature: %v", err)
		return err
	}

	t, err := time.Parse("2006-01-02 15:04:05", forecast.ReportTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return err
	}

	r.weather = &Weather{
		ReportTime:       t,
		DayTemperature:   t1,
		NightTemperature: t2,
		WindDirection:    live.Daywind,
		WindPower:        live.Daypower,
		Weather:          live.Dayweather,
	}
	r.weather.calculateWateringSeconds()
	return nil
}

func (r *Resolver) Task() {
	if err := r.GetWeatherInfo(); err == nil {
		log.Println("Watering duration:", r.weather.WaterPlanSec)
		pulse := r.weather.WaterPlanSec / 24
		if pulse > 0 {
			log.Println("Pulse duration:", pulse)
			for c, w := range r.waterIOs {
				if err := w.Watering(pulse); err != nil {
					log.Println("Failed to water:", err)
				} else {
					log.Println("Watering:", pulse)
					if _, err := r.entClient.WaterLog.Create().SetChannel(c).SetSeconds(pulse).SetManual(false).Save(context.Background()); err != nil {
						log.Println("Failed to save water log:", err)
					}
				}
			}
		}
	} else {
		log.Println("Failed to get weather info")
	}
}

type WaterIO struct {
	Pin   gpio.PinIO
	mutex sync.Mutex
}

func (w *WaterIO) Watering(seconds int) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// Check if the pin is already high
	if w.Pin.Read() == gpio.High {
		return fmt.Errorf("watering")
	}

	// Set the pin high
	if err := w.Pin.Out(gpio.High); err != nil {
		return fmt.Errorf("failed to set pin high: %v", err)
	}

	// Use a goroutine and a timer to reset the pin to low after the duration
	go func() {
		timer := time.NewTimer(time.Duration(seconds) * time.Second)
		<-timer.C

		w.mutex.Lock()
		defer w.mutex.Unlock()
		w.Pin.Out(gpio.Low)
		fmt.Println("Pin set to low after watering duration")
	}()

	return nil
}
