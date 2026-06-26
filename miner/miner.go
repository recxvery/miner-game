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

type MinerPrice struct {
	Kind     string
	Earning  int64
	Stamina  int64
	Price    int64
	Interval string
	Bonus    int64
}

var LittleMiner = MinerPrice{
	Kind:     "little miner",
	Earning:  1,
	Stamina: 30,
	Price:    5,
	Interval: "3 seconds",
	Bonus:    0,
}

var NormalMiner = MinerPrice{
	Kind:     "normal miner",
	Earning:  3,
	Stamina: 45,
	Price:    50,
	Interval: "2 seconds",
	Bonus:    0,
}

var StrongMiner = MinerPrice{
	Kind:     "strong miner",
	Earning:  10,
	Stamina: 60,
	Price:    450,
	Interval: "1 second",
	Bonus:    3,
}
