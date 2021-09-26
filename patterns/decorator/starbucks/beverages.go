package starbucks

type Beverage interface {
	GetDescription() string
	Cost()
}

type espresso struct {
}

type americano struct {
}

type cappuccino struct {
}

type mochaccino struct {
}

type latte struct {
}

func NewEspresso() *espresso {
	return new(espresso)
}

func NewAmericano() *americano {
	return new(americano)
}

func NewCappuccino() *cappuccino {
	return new(cappuccino)
}

func NewMochaccino() *mochaccino {
	return new(mochaccino)
}

func NewLate() *latte {
	return new(latte)
}
