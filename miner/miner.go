package miner

import (
	"context"
	"sync/atomic"
	"time"
)

type Miner interface {
	Run(ctx context.Context) <-chan Coal
	Info() MinerInfo
}

type SomeMiner struct {
	kind     string
	earning  Coal
	stamina  atomic.Int64
	coalChan chan Coal
	price    Coal
	interval time.Duration
	bonus    Coal
}
