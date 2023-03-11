package options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions(t *testing.T) {
	pizza := NewPizza(
		WithDough("Regular"),
		WithSauce("Tomato"),
		WithCheese("Mozzarella"),
	)

	assert.Equal(t, pizza.dough, "Regular")
	assert.Equal(t, pizza.sauce, "Tomato")
	assert.Equal(t, pizza.cheese, "Mozzarella")
	assert.Equal(t, len(pizza.toppings), 0)
}
