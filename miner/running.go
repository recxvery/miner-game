package miner

import (
	"context"
	"time"
)

func (m *Mine) Collector(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case coal := <-m.inComeCh:
			m.Mtx.Lock()
			m.CoalSum += coal
			m.Mtx.Unlock()
		}
	}
}

func (m *SomeMiner) Run(ctx context.Context) <-chan Coal {
	ticker := time.NewTicker(m.interval)

	defer func() {
		ticker.Stop()
	}()

	stam := m.stamina.Load()
	for range stam {
		m.earning += m.bonus
		select {
		case <-ctx.Done():
			return m.coalChan
		default:
		}

		select {
		case <-ctx.Done():
			return m.coalChan
		case <-ticker.C:
			m.coalChan <- m.earning
		}

		m.stamina.Add(-1)
	}

	return m.coalChan
}

func (m *Mine) BackgroundMine(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			m.Mtx.Lock()
			m.CoalSum++
			m.Mtx.Unlock()
		}
	}
}
