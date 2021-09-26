package starbucks

type Milk struct {
	Beverage Beverage
}

type Cinnamon struct {
	Beverage Beverage
}

type Chocolate struct {
	Beverage Beverage
}

type Rum struct {
	Beverage Beverage
}

func (m *Milk) GetDescription() string {
	return m.Beverage.GetDescription() + ", milk"
}

func (m *Milk) Cost() float64 {
	return m.Beverage.Cost() + 0.50
}

func (c *Cinnamon) GetDescription() string {
	return c.Beverage.GetDescription() + ", cinnamon"
}

func (c *Cinnamon) Cost() float64 {
	return c.Beverage.Cost() + 0.20
}

func (ch *Chocolate) GetDescription() string {
	return ch.Beverage.GetDescription() + ", chocolate"
}

func (ch *Chocolate) Cost() float64 {
	return ch.Beverage.Cost() + 0.50
}

func (r *Rum) GetDescription() string {
	return r.Beverage.GetDescription() + ", rum"
}

func (r *Rum) Cost() float64 {
	return r.Beverage.Cost() + 0.50
}
