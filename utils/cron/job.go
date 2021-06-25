package cron

import (
	"github.com/liov/tiga/utils/log"
)

type RedisTo struct {
}

func (RedisTo) Run() {
	log.Info("定时任务执行")
}
