package dao

import (
	"context"
	"time_machine/model"
)

type Storage interface {
	PutEvent(ctx context.Context, event *model.Event) error
	ListEventsByTime(ctx context.Context, start, end int64) ([]*model.Event, error)
	CountEventsByTime(ctx context.Context, start, end int64) (int64, error)
	ListTypeEventsByTime(ctx context.Context, eventType string, start, end int64) ([]*model.Event, error)
	CountTypeEventsByTime(ctx context.Context, eventType string, start, end int64) (int64, error)
}
