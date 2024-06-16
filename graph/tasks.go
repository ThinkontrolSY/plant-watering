package graph

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// 根据气温和湿度计算浇水时长的函数
func calculateWateringSeconds(temperature, humidity float64) int {
	// 定义绣球花基础需水量（分钟）
	baseTime := 20.0

	// 定义气温和湿度对需水量的影响
	// 假设每增加1摄氏度，增加1分钟，每减少1%湿度，增加0.5分钟
	tempAdjustment := (temperature - 20) * 1.0
	humidityAdjustment := (60 - humidity) * 0.5

	// 计算最终的浇水时长
	wateringTime := baseTime + tempAdjustment + humidityAdjustment

	// 确保浇水时间不小于基础时间，也不过长
	if wateringTime < baseTime/2 {
		wateringTime = baseTime
	} else if wateringTime > 90 {
		wateringTime = 90 // 设定一个上限，避免过度浇水
	}

	return int(wateringTime * 60) // 转换为秒
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

	if len(weatherResponse.Lives) == 0 {
		log.Println("No weather data")
		return fmt.Errorf("no weather data")
	}

	live := weatherResponse.Lives[0]

	temperature, err := strconv.ParseFloat(live.Temperature, 64)
	if err != nil {
		log.Printf("Failed to parse temperature: %v", err)
		return err
	}

	humidity, err := strconv.ParseFloat(live.Humidity, 64)
	if err != nil {
		log.Printf("Failed to parse humidity: %v", err)
		return err
	}

	t, err := time.Parse("2006-01-02 15:04:05", live.ReportTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return err
	}

	r.weather = &Weather{
		ReportTime:    t,
		Temperature:   temperature,
		Humidity:      humidity,
		WindDirection: live.WindDirection,
		WindPower:     live.WindPower,
		Weather:       live.Weather,
	}
	return nil
}

func (r *Resolver) Task() {
	if err := r.GetWeatherInfo(); err == nil {
		waterDuration := calculateWateringSeconds(r.weather.Temperature, r.weather.Humidity)
		log.Println("Watering duration:", waterDuration)
	} else {
		log.Println("Failed to get weather info")
	}
}
