package miner

import "errors"

var ErrNotEnoughBalance error = errors.New("We don't have enough money to hire this miner type")
var ClassWasNotFound error = errors.New("This class doesn't exist")
var ToolWasNotFound error = errors.New("This tool doesn't exist")
var GameCanNotEnd error = errors.New("You still didn't purchased all items")
