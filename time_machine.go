package time_machine

import (
	"context"
	"time"
	"time_machine/dao"
	"time_machine/model"
)

type Machine struct {
	timeDuration time.Duration //存储数据的时间
	storage      dao.Storage   // 存储引擎
}

func NewTimeMachine(conf *TimeMachineConf) *Machine {
	return &Machine{storage: dao.InitRedisStorage(conf.RedisConf, conf.TTL)}
}

func (r *Machine) Put(ctx context.Context, event *model.Event) error {
	return r.storage.PutEvent(ctx, event)
}

func (r *Machine) ListByTime(ctx context.Context, start, end int64) ([]*model.Event, error) {
	return r.storage.ListEventsByTime(ctx, start, end)
}

func (r *Machine) CountsByTime(ctx context.Context, start, end int64) (int64, error) {
	return r.storage.CountEventsByTime(ctx, start, end)
}

func (r *Machine) ListTypeByTime(ctx context.Context, eventType string, start, end int64) ([]*model.Event, error) {
	return r.storage.ListTypeEventsByTime(ctx, eventType, start, end)
}

func (r *Machine) CountsTypeByTime(ctx context.Context, eventType string, start, end int64) (int64, error) {
	return r.storage.CountTypeEventsByTime(ctx, eventType, start, end)
}
