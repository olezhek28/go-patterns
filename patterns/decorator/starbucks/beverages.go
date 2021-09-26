package starbucks

type Beverage interface {
	GetDescription() string
	Cost() float64
}

type espresso struct {
	description string
	price       float64
}

type americano struct {
	description string
	price       float64
}

type cappuccino struct {
	description string
	price       float64
}

type mochaccino struct {
	description string
	price       float64
}

type latte struct {
	description string
	price       float64
}

func NewEspresso() *espresso {
	return &espresso{
		description: "Espresso",
		price:       2.00,
	}
}

func (e *espresso) GetDescription() string {
	return e.description
}

func (e *espresso) Cost() float64 {
	return e.price
}

func NewAmericano() *americano {
	return &americano{
		description: "Americano",
		price:       2.50,
	}
}

func (a *americano) GetDescription() string {
	return a.description
}

func (a *americano) Cost() float64 {
	return a.price
}

func NewCappuccino() *cappuccino {
	return &cappuccino{
		description: "Cappuccino",
		price:       2.30,
	}
}

func (c *cappuccino) GetDescription() string {
	return c.description
}

func (c *cappuccino) Cost() float64 {
	return c.price
}

func NewMochaccino() *mochaccino {
	return &mochaccino{
		description: "Mochaccino",
		price:       2.80,
	}
}

func (m *mochaccino) GetDescription() string {
	return m.description
}

func (m *mochaccino) Cost() float64 {
	return m.price
}

func NewLate() *latte {
	return &latte{
		description: "Latte",
		price:       3.00,
	}
}

func (l *latte) GetDescription() string {
	return l.description
}

func (l *latte) Cost() float64 {
	return l.price
}
