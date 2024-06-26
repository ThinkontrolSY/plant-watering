//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3/sysfs"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	croner     *cron.Cron
	weather    *Weather
	waterIOs   map[string]*WaterIO
	statistics map[string]*WaterStatistic
	baseTime   int32
}

type WaterStatistic struct {
	AutoWatering   int32
	ManualWatering int32
}

func (r *Resolver) Start() {
	r.croner = cron.New()
	// Add a task to cron to run every 10 minutes
	// 早上 6 点到 9 点每 10 分钟执行一次
	_, err := r.croner.AddFunc("*/10 6-7 * * *", r.Task)
	if err != nil {
		log.Println("Error scheduling morning job:", err)
	}
	// 每天早上 6 点执行一次
	_, err = r.croner.AddFunc("0 6 * * *", func() {
		for _, w := range r.statistics {
			w.AutoWatering = 0
			w.ManualWatering = 0
		}
	})
	if err != nil {
		log.Println("Error scheduling weather job:", err)
	}

	r.croner.Start()
	log.Println("Cron started")
	go r.GetWeatherInfo()
}

func NewSchema() graphql.ExecutableSchema {
	pinN1 := sysfs.Pins[198]
	pinN2 := sysfs.Pins[199]
	if err := pinN1.Out(gpio.Low); err != nil {
		log.Println("Failed to set pin N1 to low:", err)
	}
	if err := pinN2.Out(gpio.Low); err != nil {
		log.Println("Failed to set pin N2 to low:", err)
	}
	baseTime := viper.GetInt32("baseTime")
	if baseTime == 0 {
		baseTime = 60
	}
	resolver := &Resolver{
		baseTime: baseTime,
		waterIOs: map[string]*WaterIO{
			"N1": {
				Pin: pinN1,
			},
			"N2": {
				Pin: pinN2,
			},
		},
		statistics: map[string]*WaterStatistic{
			"N1": {},
			"N2": {},
		},
	}
	config := Config{
		Resolvers: resolver,
	}
	resolver.Start()
	return NewExecutableSchema(config)
}
