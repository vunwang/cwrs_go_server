package cwrs_goroutine

import (
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

var GlobalPool *ants.Pool
var CWRS sync.WaitGroup

var log = cwrs_zap_logger.ZapLogger

func init() {
	InitAntsPool()
}

func InitAntsPool() {
	pool, err := ants.NewPool(1000)
	if err != nil {
		log.Error(fmt.Sprintf("线程池初始化失败:%s", err.Error()))
	}
	GlobalPool = pool
}
