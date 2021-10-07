package main

import "fmt"

// User friendly facade to cover the coffeMachine API

func makeAmericano(size float32) {
	fmt.Println("\nPreparing an Americano\n----------------------")
	// Connect to the low-level machine API
	americano := &CoffeeMachine{}
	// Use 5/8 the amount of beans for 'size' mL hot water and set grinder to level 6
	waterAmount := size
	beanAmount := (size / 8.0) * 5.0
	grinderLevel := 6

	americano.startCoffee(beanAmount, grinderLevel)
	americano.grindBeans()
	americano.useWater(waterAmount)
	americano.addMilkFoam(false)

	americano.endCoffee()

	fmt.Println("Americano is ready")
}

func makeLatte(size float32) {
	fmt.Println("\nPreparing a Latte\n----------------------")
	// Connect to the low-level machine API
	latte := &CoffeeMachine{}
	// Use 3/8 the amount of beans for 'size' mL hot water and set grinder to level 2
	waterAmount := size
	beanAmount := (size / 8.0) * 3.0
	grinderLevel := 2

	latte.startCoffee(beanAmount, grinderLevel)
	latte.grindBeans()
	latte.useWater(waterAmount)

	// Add milk amount 2/8th of 'size'
	milk := (size / 8.0) * 2.0
	latte.useMilk(milk)
	latte.addMilkFoam(true)

	latte.endCoffee()

	fmt.Println("Latte is ready")
}
