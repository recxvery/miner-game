package input

import "errors"

type MinerDTO struct {
	Class string
}

func (m MinerDTO) ValidateToCreate() error {
	if m.Class != "strong" && m.Class != "normal" && m.Class != "little" {
		return errors.New("Incorrect Miner class")
	}

	return nil
}




