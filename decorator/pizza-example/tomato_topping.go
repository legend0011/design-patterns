package main

type TomatoTopping struct {
	pizza IPizza
}

// Implement the same interface as base object
func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
