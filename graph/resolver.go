//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"log"
	"plant-watering/ent"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/robfig/cron/v3"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type WeatherResponse struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Lives    []struct {
		Province      string `json:"province"`
		City          string `json:"city"`
		Adcode        string `json:"adcode"`
		Weather       string `json:"weather"`
		Temperature   string `json:"temperature_float"`
		WindDirection string `json:"winddirection"`
		WindPower     string `json:"windpower"`
		Humidity      string `json:"humidity_float"`
		ReportTime    string `json:"reporttime"`
	} `json:"lives"`
}

type Weather struct {
	ReportTime    time.Time
	Temperature   float64
	Humidity      float64
	WindDirection string
	WindPower     string
	Weather       string
}
type Resolver struct {
	entClient *ent.Client
	croner    *cron.Cron
	weather   *Weather
}

func (r *Resolver) Start() {
	r.croner = cron.New()
	// Add a task to cron to run every 10 minutes
	// 早上 7 点到 10 点每 10 分钟执行一次
	_, err := r.croner.AddFunc("*/10 7-9 * * *", r.Task)
	if err != nil {
		log.Println("Error scheduling morning job:", err)
	}
	// 下午 16 点到 19 点每 10 分钟执行一次
	_, err = r.croner.AddFunc("*/10 16-18 * * *", r.Task)
	if err != nil {
		log.Println("Error scheduling evening job:", err)
	}
	r.croner.Start()
}

func NewSchema(
	entClient *ent.Client,
) graphql.ExecutableSchema {
	resolver := &Resolver{
		entClient: entClient,
	}
	config := Config{
		Resolvers: resolver,
	}
	resolver.Start()
	return NewExecutableSchema(config)
}
