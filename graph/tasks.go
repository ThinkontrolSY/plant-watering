package graph

import (
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
			Daytemp      string `json:"daytemp"`
			Nighttemp    string `json:"nighttemp"`
			Daywind      string `json:"daywind"`
			Nightwind    string `json:"nightwind"`
			Daypower     string `json:"daypower"`
			Nightpower   string `json:"nightpower"`
		} `json:"casts"`
	} `json:"forecasts"`
}

type Weather struct {
	ReportTime       time.Time
	DayTemperature   int32
	NightTemperature int32
	WindDirection    string
	WindPower        string
	Weather          string
	WaterPlanSec     int32
}

func (w *Weather) CalculateWateringSeconds(baseTime int32) {

	// 如果温度太低，不进行浇水
	if w.DayTemperature < 0 || w.NightTemperature < 0 {
		w.WaterPlanSec = 0
	}

	// 如果天气包含“雨”或“雪”，不进行浇水
	if strings.Contains(w.Weather, "雨") || strings.Contains(w.Weather, "雪") {
		w.WaterPlanSec = 0
	}

	// 定义气温对需水量的影响
	tempAdjustment := (w.DayTemperature + w.NightTemperature - 40) * 2

	// 定义天气对需水量的影响
	weatherAdjustment := int32(0)
	if strings.Contains(w.Weather, "霾") || strings.Contains(w.Weather, "雾") {
		weatherAdjustment = 5
	} else if strings.Contains(w.Weather, "风") {
		weatherAdjustment = 10
	}

	// 计算最终的浇水时长
	wateringTime := baseTime + tempAdjustment + weatherAdjustment

	// 确保浇水时间不小于基础时间，也不过长
	if wateringTime < 0 {
		wateringTime = 0 // 设定一个下限，避免过度浇水
	} else if wateringTime > 120 {
		wateringTime = 120 // 设定一个上限，避免过度浇水
	}

	w.WaterPlanSec = wateringTime
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
	params.Add("extensions", "all")
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

	t1, err := strconv.ParseInt(live.Daytemp, 10, 32)
	if err != nil {
		log.Printf("Failed to parse day temperature: %v", err)
		return err
	}

	t2, err := strconv.ParseInt(live.Nighttemp, 10, 32)
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
		DayTemperature:   int32(t1),
		NightTemperature: int32(t2),
		WindDirection:    live.Daywind,
		WindPower:        live.Daypower,
		Weather:          live.Dayweather,
	}
	r.weather.CalculateWateringSeconds(r.baseTime)
	return nil
}

func (r *Resolver) Task() {
	if err := r.GetWeatherInfo(); err == nil {
		log.Println("Watering duration:", r.weather.WaterPlanSec)
		pulse := r.weather.WaterPlanSec / 12
		if pulse > 0 {
			log.Println("Pulse duration:", pulse)
			for c, w := range r.waterIOs {
				if err := w.Watering(pulse); err != nil {
					log.Println("Failed to water:", err)
				} else {
					log.Println("Watering:", pulse)
					if s, ok := r.statictics[c]; ok {
						s.AutoWatering += pulse
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

func (w *WaterIO) Watering(seconds int32) error {
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
		log.Println("Pin set to low after watering duration")
	}()

	return nil
}
