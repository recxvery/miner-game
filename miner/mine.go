package miner

import (
	"context"
	"sync"
)

type Coal int64

type Miners interface {
	Run(ctx context.Context) <-chan Coal
	Info() MinerInfo
}

type MineInfo struct {
	MinersCount int64
	coalSum     int64
	mtx         sync.Mutex
}
