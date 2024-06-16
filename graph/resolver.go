//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"log"
	"plant-watering/ent"

	"github.com/99designs/gqlgen/graphql"
	"github.com/robfig/cron/v3"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3/rpi"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	entClient *ent.Client
	croner    *cron.Cron
	weather   *Weather
	waterIOs  map[string]*WaterIO
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
	pinN1 := rpi.P1_11
	pinN2 := rpi.P1_12
	if err := pinN1.Out(gpio.Low); err != nil {
		log.Println("Failed to set pin N1 to low:", err)
	}
	if err := pinN2.Out(gpio.Low); err != nil {
		log.Println("Failed to set pin N2 to low:", err)
	}
	resolver := &Resolver{
		entClient: entClient,
		waterIOs: map[string]*WaterIO{
			"N1": {
				Pin: pinN1,
			},
			"N2": {
				Pin: pinN2,
			},
		},
	}
	config := Config{
		Resolvers: resolver,
	}
	resolver.Start()
	return NewExecutableSchema(config)
}
