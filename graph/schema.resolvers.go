package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"log"
	"plant-watering/ent"
	"plant-watering/ent/waterlog"
	"plant-watering/graph/model"
	"time"
)

// Water is the resolver for the water field.
func (r *mutationResolver) Water(ctx context.Context, input model.WateringInput) (bool, error) {
	log.Printf("Watering plant %s", input.Channel)
	if w, ok := r.waterIOs[input.Channel]; ok {
		if err := w.Watering(input.Seconds); err != nil {
			return false, fmt.Errorf("failed to water: %v", err)
		} else {
			log.Printf("Watering %s: %d", input.Channel, input.Seconds)
			if _, err := r.entClient.WaterLog.Create().SetChannel(input.Channel).SetSeconds(input.Seconds).SetManual(true).Save(context.Background()); err != nil {
				log.Printf("Failed to save water log: %v", err)
			}
			return true, nil
		}
	} else {
		return false, fmt.Errorf("channel %s not found", input.Channel)
	}
}

// Channels is the resolver for the channels field.
func (r *queryResolver) Channels(ctx context.Context) ([]string, error) {
	return []string{"N1", "N2"}, nil
}

// Weather is the resolver for the weather field.
func (r *queryResolver) Weather(ctx context.Context) (*Weather, error) {
	return r.weather, nil
}

// WaterStatistic is the resolver for the waterStatistic field.
func (r *queryResolver) WaterStatistic(ctx context.Context, channel string) (*model.WaterStatistic, error) {
	// 设置当前日期的早上6点时间
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location())

	// 查询
	var results []struct {
		Manual  bool `json:"manual"`
		Seconds int  `json:"seconds"`
	}
	err := r.entClient.WaterLog.Query().
		Where(
			waterlog.TimeGTE(startOfDay),
			waterlog.ChannelEQ("channel"),
		).
		GroupBy(waterlog.FieldManual).
		Aggregate(ent.Sum(waterlog.FieldSeconds)).
		Scan(ctx, &results)

	if err != nil {
		fmt.Printf("failed querying waterlog: %v", err)
		return nil, err
	}

	// 计算
	var autoSeconds, manualSeconds int
	for _, result := range results {
		if result.Manual {
			manualSeconds = result.Seconds
		} else {
			autoSeconds = result.Seconds
		}
	}
	return &model.WaterStatistic{
		AutoWatering:   autoSeconds,
		ManualWatering: manualSeconds,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
