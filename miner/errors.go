package miner

import "errors"

var errNotEnoughBalance error = errors.New("We don't have enough money to hire this miner type")