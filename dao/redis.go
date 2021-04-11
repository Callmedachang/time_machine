package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"time_machine/model"

	"github.com/go-redis/redis/v8"
)

const MachineSortedSetKey = "MachineSortedSetKey"

type RedisStorage struct {
	redis *redis.Client
	ttl   time.Duration
}

type ConnectionConf struct {
	RedisAddress string
	RedisPWD     string
	RedisDB      int
}

func InitRedisStorage(conf *ConnectionConf, ttl time.Duration) *RedisStorage {
	r := &RedisStorage{
		redis: newRedisClient(conf),
		ttl:   ttl,
	}
	r.clearData(context.Background())
	return r
}

func newRedisClient(conf *ConnectionConf) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPWD, // no password set
		DB:       conf.RedisDB,  // use default DB
	})
}

func (r *RedisStorage) setEventData(ctx context.Context, id int64, event *model.Event) error {
	data, _ := json.Marshal(event)
	return r.redis.Set(ctx, fmt.Sprint(id), data, r.ttl).Err()
}

func (r *RedisStorage) getMemberID(ctx context.Context) int64 {
	return r.redis.Incr(ctx, "TimeMachineIdKey").Val()
}

func (r *RedisStorage) zAddEventData(ctx context.Context, id int64, event *model.Event) error {
	return r.redis.ZAdd(ctx, MachineSortedSetKey, &redis.Z{Score: float64(event.Timestamp.Unix()), Member: id}).Err()
}

func (r *RedisStorage) PutEvent(ctx context.Context, event *model.Event) error {
	mid := r.getMemberID(ctx)
	if err := r.setEventData(ctx, mid, event); err != nil {
		return err
	}
	if err := r.zAddEventData(ctx, mid, event); err != nil {
		return err
	}
	return nil
}

func (r *RedisStorage) ListEventsByTime(ctx context.Context, start, end int64) ([]*model.Event, error) {
	//TODO:IMPL
	return nil, nil
}

func (r *RedisStorage) CountEventsByTime(ctx context.Context, start, end int64) (int64, error) {
	//TODO:IMPL
	return 0, nil
}

func (r *RedisStorage) ListTypeEventsByTime(ctx context.Context, eventType string, start, end int64) ([]*model.Event, error) {
	//TODO:IMPL
	return nil, nil
}

func (r *RedisStorage) CountTypeEventsByTime(ctx context.Context, eventType string, start, end int64) (int64, error) {
	//TODO:IMPL
	return 0, nil
}

func (r *RedisStorage) clearData(ctx context.Context) {
	go func() {
		timeTicker := time.NewTicker(time.Second * 20)
		for {
			select {
			case <-timeTicker.C:
				if err := r.doClear(ctx); err != nil {
					log.Printf("TimeMachine.RedisStorage.doClearErr:(%+v)", err)
				}
			}
		}
	}()
}

func (r *RedisStorage) doClear(ctx context.Context) error {
	return r.redis.ZRemRangeByScore(ctx, MachineSortedSetKey,
		fmt.Sprint(time.Now().Add(-1*r.ttl).Unix()),
		fmt.Sprint(time.Now().Unix())).Err()
}
