package time_machine

import (
	"time"
	"time_machine/dao"
)

type TimeMachineConf struct {
	RedisConf *dao.ConnectionConf
	TTL       time.Duration
}
