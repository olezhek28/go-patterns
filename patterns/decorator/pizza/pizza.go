package pizza

type Pizza interface {
	GetPrice() float64
	GetCalories() int
}

type cheesePizza struct {
	price    float64
	calories int
}

func NewCheesePizza() *cheesePizza {
	return &cheesePizza{
		price:    10.00,
		calories: 500,
	}
}

func (cp *cheesePizza) GetPrice() float64 {
	return cp.price
}

func (cp *cheesePizza) GetCalories() int {
	return cp.calories
}
