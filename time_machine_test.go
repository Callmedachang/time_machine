package time_machine

import (
	"context"
	"testing"
	"time"
	"time_machine/dao"
	"time_machine/model"

	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func TestMachine_Put(t *testing.T) {
	machine := NewTimeMachine(&TimeMachineConf{
		RedisConf: &dao.ConnectionConf{RedisAddress: "127.0.0.1:6379"},
		TTL:       time.Hour,
	})
	err := machine.Put(ctx, &model.Event{
		Metadata:  []byte("testDemo1"),
		Timestamp: time.Now(),
	})
	_ = assert.Nil(t, err, nil)
}
