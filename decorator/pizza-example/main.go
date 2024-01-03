package main

import "fmt"

// Wearing clothes is an example of using decorators.
// When you’re cold, you wrap yourself in a sweater. If you’re still cold with a sweater, you can wear a jacket on top. If it’s raining, you can put on a raincoat.
// All of these garments “extend” your basic behavior but aren’t part of you,
//  and you can easily take off any piece of clothing whenever you don’t need it.

// But each class should implement the same interface

func main() {
	pizza := &VeggieMania{}
	pizzaWithCheese := &CheeseTopping{pizza: pizza}
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
