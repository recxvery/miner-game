package equipment

type Equip struct {
	Price    int64
	IsBought bool
}

var Picxaxe = Equip{
	Price:    3000,
	IsBought: false,
}
var AirShaft = Equip{
	Price:    15000,
	IsBought: false,
}

var Trolley = Equip{
	Price:    50000,
	IsBought: false,
}

func (e *Equip) Buy() {
	e.IsBought = true
}
