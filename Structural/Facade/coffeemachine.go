package main

import "fmt"

// Low-level coffee machine API
type CoffeeMachine struct {
	beanAmount		float32 // how much coffee beans to use per cup
	grinderLevel	int // Desired coarseness of ground beans
	waterTemp		int // Water heater temp readout
	waterAmount		float32 // Amount of water per cup
	milkAmount		float32 // Amount of milk per cup
	addFoam			bool	// Top off with foam?
}

func (c *CoffeeMachine) startCoffee(beans float32, grind int) {
	c.beanAmount = beans
	c.grinderLevel = grind
	fmt.Println("Order for coffee with", beans, "grams of beans and grind level", grind, ".")
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Order processed.")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding", c.beanAmount, "grams of beans at level", c.grinderLevel, ".")
	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding", amount, "mL of milk.")
	c.milkAmount = amount
	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp int) int {
	fmt.Println("Setting water temperature to", temp, "°C.")
	c.waterTemp = temp
	return temp
}

func (c *CoffeeMachine) useWater(amount float32) float32 {
	fmt.Println("Adding", amount, "mL of hot water.")
	c.waterAmount = amount
	return amount
}

func (c *CoffeeMachine) addMilkFoam(foam bool) bool {
	if foam == true {
		fmt.Println("Adding milk foam.")
	} else {
		fmt.Println("Not adding milk foam.")
	}
	c.addFoam = foam
	return foam
}