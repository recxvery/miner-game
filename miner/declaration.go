package miner

import (
	"context"
	"miner-game/miner/equipment"
	"sync"
	"time"
)

type Coal int64

type MinerInfo struct {
	Class   string
	Stamina int64
}

func (m *SomeMiner) Info() MinerInfo {
	return MinerInfo{
		Class:   m.kind,
		Stamina: m.stamina.Load(),
	}
}

type Mine struct {
	Mtx           sync.RWMutex
	Miners        map[int64]*SomeMiner
	mineCtx       context.Context
	cancelCtx     context.CancelFunc
	RunningTime   time.Time
	MinersCount   int64
	CoalSum       Coal
	inComeCh      chan Coal
	Equipmentlist map[string]*equipment.Equip
}

func NewMine(ctx context.Context, cancel context.CancelFunc) *Mine {

	return &Mine{
		mineCtx:     ctx,
		cancelCtx:   cancel,
		Miners:      make(map[int64]*SomeMiner),
		inComeCh:    make(chan Coal),
		RunningTime: time.Now(),
		Equipmentlist: map[string]*equipment.Equip{
			"Picxaxe":  &equipment.Picxaxe,
			"AirShaft": &equipment.AirShaft,
			"Trolley":  &equipment.Trolley,
		},
	}
}
