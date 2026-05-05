package miner

import (
	"context"
	"time"
)

type MinerInfo struct {
}

type Miner struct {
	earningInInterval Coal
	stamina        int64
	earning        chan Coal
	workInterval   time.Duration
	producingBouns Coal
}

func (m *MineInfo) NewLittleMiner() (*Miner, error) {
	m.mtx.Lock()
	if m.coalSum <= 5 {
		m.mtx.Unlock()
		return nil, errNotEnoughBalance
	}
	m.MinersCount++
	m.coalSum -= 5
	m.mtx.Unlock()
	return &Miner{
		earningInInterval: 1,
		stamina: 30,
		earning: make(chan Coal, 30),
		workInterval: time.Second * 3,
		producingBouns: 0,
	}, nil
}


func (m *MineInfo) NewNormalMiner() (*Miner, error) {
	m.mtx.Lock()
	if m.coalSum <= 50 {
		m.mtx.Unlock()
		return nil, errNotEnoughBalance
	}
	m.MinersCount++
	m.coalSum -= 50
	m.mtx.Unlock()
	return &Miner{
		earningInInterval : 3,
		stamina: 45,
		earning: make(chan Coal, 45),
		workInterval: time.Second * 2,
		producingBouns: 0,
	}, nil
}

func (m *MineInfo) NewStrongMiner() (*Miner, error) {
	m.mtx.Lock()
	if m.coalSum <= 450 {
		m.mtx.Unlock()
		return nil, errNotEnoughBalance
	}
	m.MinersCount++
	m.coalSum -= 450
	m.mtx.Unlock()
	return &Miner{
		earningInInterval: 10,
		stamina: 60,
		earning: make(chan Coal, 60),
		workInterval: time.Second * 1,
		producingBouns: 3,
	}, nil
}


func (m *Miner) Run(ctx context.Context) <-chan Coal {
	ticker := time.NewTicker(m.workInterval)
	defer func() {
		ticker.Stop()
		close(m.earning)
	}()
	for range m.stamina {
		select {
		case <- ctx.Done():
			return m.earning
		case <- ticker.C:
				m.earning <- m.earningInInterval + m.producingBouns
		}
	}

	return m.earning
}