package options

//TODO: Implement this using BuilderPattern

type Pizza struct {
	dough    string
	sauce    string
	cheese   string
	toppings []string
}

type Options struct {
	Dough    string
	Sauce    string
	Cheese   string
	Toppings []string
}

type PizzaOption func(*Options)

func WithDough(dough string) PizzaOption {
	return func(po *Options) {
		po.Dough = dough
	}
}

func WithSauce(sauce string) PizzaOption {
	return func(po *Options) {
		po.Sauce = sauce
	}
}

func WithCheese(cheese string) PizzaOption {
	return func(po *Options) {
		po.Cheese = cheese
	}
}

func WithToppings(toppings []string) PizzaOption {
	return func(po *Options) {
		po.Toppings = toppings
	}
}

func NewPizza(options ...PizzaOption) *Pizza {
	opts := &Options{}

	for _, option := range options {
		option(opts)
	}

	pizza := &Pizza{
		dough:    opts.Dough,
		sauce:    opts.Sauce,
		cheese:   opts.Cheese,
		toppings: opts.Toppings,
	}

	return pizza
}

//func main() {
//	pizza := NewPizza(
//		WithDough("Regular"),
//		WithSauce("Tomato"),
//		WithCheese("Mozzarella"),
//		WithToppings([]string{"Pepperoni", "Olives", "Mushrooms"}),
//	)
//	println(pizza.dough)
//	println(pizza.sauce)
//	println(pizza.cheese)
//	println(pizza.toppings)
//}
