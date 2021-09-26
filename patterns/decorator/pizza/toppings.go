package pizza

const (
	pepperoniPrice = 1.75
	artichokePrice = 2.00
	anchovyPrice   = 2.50
	tomatoPrice    = 0.75

	pepperoniCalories = 200
	artichokeCalories = 150
	anchovyCalories   = 350
	tomatoCalories    = 50
)

type PepperoniTopping struct {
	Pizza Pizza
}

type ArtichokeTopping struct {
	Pizza Pizza
}

type AnchovyTopping struct {
	Pizza Pizza
}

type TomatoTopping struct {
	Pizza Pizza
}

func (pt *PepperoniTopping) GetPrice() float64 {
	return pt.Pizza.GetPrice() + pepperoniPrice
}

func (pt *PepperoniTopping) GetCalories() int {
	return pt.Pizza.GetCalories() + pepperoniCalories
}

func (art *ArtichokeTopping) GetPrice() float64 {
	return art.Pizza.GetPrice() + artichokePrice
}

func (art *ArtichokeTopping) GetCalories() int {
	return art.Pizza.GetCalories() + artichokeCalories
}

func (ant *AnchovyTopping) GetPrice() float64 {
	return ant.Pizza.GetPrice() + anchovyPrice
}

func (ant *AnchovyTopping) GetCalories() int {
	return ant.Pizza.GetCalories() + anchovyCalories
}

func (tt *TomatoTopping) GetPrice() float64 {
	return tt.Pizza.GetPrice() + tomatoPrice
}

func (tt *TomatoTopping) GetCalories() int {
	return tt.Pizza.GetCalories() + tomatoCalories
}
