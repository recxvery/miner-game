package miner

import (
	"miner-game/miner/equipment"
	"time"
)

type InfoMine struct {
	Balance Coal
	Count   int
}

func (m *Mine) GetMineInfo() InfoMine {
	m.Mtx.RLock()
	defer m.Mtx.RUnlock()
	inf := InfoMine{
		Balance: m.CoalSum,
		Count:   len(m.Miners),
	}

	return inf
}

func (m *Mine) HireMiner(kind string) (MinerInfo, error) {
	var minerDto MinerInfo

	switch kind {
	case "strong":
		m.Mtx.Lock()
		defer m.Mtx.Unlock()
		if m.CoalSum-450 >= 0 {
			m.CoalSum -= 450
			newMiner := &SomeMiner{
				kind:     kind,
				earning:  10,
				coalChan: m.inComeCh,
				price:    450,
				interval: 1 * time.Second,
				bonus:    3,
			}

			newMiner.stamina.Store(60)
			minerDto.Class, minerDto.Stamina = newMiner.kind, newMiner.stamina.Load()

			m.Miners[m.MinersCount] = newMiner
			m.MinersCount++

			go newMiner.Run(m.mineCtx)

		} else {
			return MinerInfo{}, ErrNotEnoughBalance
		}

	case "normal":
		m.Mtx.Lock()
		defer m.Mtx.Unlock()
		if m.CoalSum-50 >= 0 {
			m.CoalSum -= 50
			newMiner := &SomeMiner{
				kind:     kind,
				earning:  3,
				coalChan: m.inComeCh,
				price:    50,
				interval: 2 * time.Second,
				bonus:    0,
			}
			newMiner.stamina.Store(45)
			minerDto.Class, minerDto.Stamina = newMiner.kind, newMiner.stamina.Load()

			m.Miners[m.MinersCount] = newMiner
			m.MinersCount++

			go newMiner.Run(m.mineCtx)
		} else {
			return MinerInfo{}, ErrNotEnoughBalance
		}

	case "little":
		m.Mtx.Lock()
		defer m.Mtx.Unlock()
		if m.CoalSum-5 >= 0 {
			m.CoalSum -= 5
			newMiner := &SomeMiner{
				kind:     kind,
				earning:  1,
				coalChan: m.inComeCh,
				price:    5,
				interval: 3 * time.Second,
				bonus:    0,
			}
			newMiner.stamina.Store(30)
			minerDto.Class, minerDto.Stamina = newMiner.kind, newMiner.stamina.Load()

			m.Miners[m.MinersCount] = newMiner
			m.MinersCount++

			go newMiner.Run(m.mineCtx)
		} else {
			return MinerInfo{}, ErrNotEnoughBalance
		}
	}

	return minerDto, nil
}
func (m *Mine) GetAllMiners() map[int64]MinerInfo {
	tmp := make(map[int64]MinerInfo)

	m.Mtx.RLock()
	defer m.Mtx.RUnlock()
	for k, miner := range m.Miners {
		info := miner.Info()
		tmp[k] = info
	}

	return tmp
}

func (m *Mine) GetWorkingMiners() map[int64]MinerInfo {
	tmp := make(map[int64]MinerInfo)

	m.Mtx.RLock()
	defer m.Mtx.RUnlock()
	for k, miner := range m.Miners {
		info := miner.Info()
		if info.Stamina > 0 {
			tmp[k] = info
		}
	}

	return tmp
}

func (m *Mine) GetMinersWorkingClass(class string) map[int64]MinerInfo {
	tmp := make(map[int64]MinerInfo)

	m.Mtx.RLock()
	defer m.Mtx.RUnlock()
	for k, miner := range m.Miners {
		info := miner.Info()
		if info.Class == class {
			tmp[k] = info
		}
	}

	return tmp
}

func (m *Mine) GetBalance() int64 {
	m.Mtx.RLock()
	defer m.Mtx.RUnlock()

	balance := m.CoalSum
	return int64(balance)
}

func (m *Mine) GetBoughtEquipmentList() map[string]equipment.Equip {
	tmp := make(map[string]equipment.Equip)
	m.Mtx.RLock()
	defer m.Mtx.RUnlock()
	for k, v := range m.Equipmentlist {
		if v.IsBought == true {
			tmp[k] = *v
		}
	}

	return tmp
}

func (m *Mine) GetUnBoughtEquipmentList() map[string]equipment.Equip {
	tmp := make(map[string]equipment.Equip)
	m.Mtx.RLock()
	defer m.Mtx.RUnlock()
	for k, v := range m.Equipmentlist {
		if v.IsBought == false {
			tmp[k] = *v
		}
	}

	return tmp
}

func (m *Mine) BuyEquipment(tool string) (equipment.Equip, error) {
	m.Mtx.Lock()
	defer m.Mtx.Unlock()

	equip, ok := m.Equipmentlist[tool]
	if !ok {
		return equipment.Equip{}, ToolWasNotFound
	}

	if m.CoalSum-Coal(equip.Price) < 0 {
		return equipment.Equip{}, ErrNotEnoughBalance
	}

	m.CoalSum -= Coal(equip.Price)
	equip.Buy()

	return *equip, nil
}

func (m *Mine) GetEquipmentList() map[string]equipment.Equip {
	tmp := make(map[string]equipment.Equip)

	m.Mtx.Lock()
	defer m.Mtx.Unlock()

	for k, v := range m.Equipmentlist {
		tmp[k] = *v
	}

	return tmp
}

func (m *Mine) StopGame() error {
	m.Mtx.RLock()
	defer m.Mtx.RUnlock()

	c := 0

	for _, v := range m.Equipmentlist {
		if v.IsBought == true {
			c++
		}
	}

	if c == 3 {
		if m.cancelCtx != nil {
			m.cancelCtx()
		}
	} else {
		return GameCanNotEnd
	}

	return nil
}
