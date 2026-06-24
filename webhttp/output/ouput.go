package output

import (
	"encoding/json"
	"log"
	"miner-game/miner"
	"miner-game/miner/equipment"
	"time"
)

type ErrorDTO struct {
	Message string
	Time    time.Time
}

type MineDTO struct {
	Miners        map[int64]miner.MinerInfo
	RunningTime   string
	MinersCount   int64
	CoalSum       miner.Coal
	Equipmentlist map[string]equipment.Equip
}

type MinersDTObyType map[int64]miner.MinerInfo

func NewMinerDTOByType(m *miner.Mine, inputClass string) MinersDTObyType {
	m.Mtx.RLock()
	defer m.Mtx.RUnlock()

	tmp := make(map[int64]miner.MinerInfo)
	

	for id, miner := range m.Miners {
		if miner.Info().Class == inputClass {
			tmp[id] = miner.Info()
		}
	}

	return tmp
}

func NewAllMinerDTO(m *miner.Mine) map[int64]miner.MinerInfo {
	m.Mtx.RLock()
	defer m.Mtx.RUnlock()

	tmp := make(map[int64]miner.MinerInfo)

	for id, miner := range m.Miners {
		tmp[id] = miner.Info()
	}

	return tmp
}

func (e *ErrorDTO) ErrToJson() string {
	data, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return string(data)
}

func NewEquipmentDTO(m *miner.Mine) map[string]equipment.Equip {
	m.Mtx.RLock()
	defer m.Mtx.RUnlock()

	tmp := make(map[string]equipment.Equip)

	for key, tool := range m.Equipmentlist {
		tmp[key] = *tool
	}

	return tmp
}
