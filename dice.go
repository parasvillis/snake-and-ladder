package main

import (
	"math/rand"
)

type Dice struct {
	numberOfDice int
}

func NewDice(numberOfDice int) *Dice {
	return &Dice{
		numberOfDice: numberOfDice,
	}
}

func (d *Dice) rollDice() int {
	// rand.Seed(time.Now().UnixNano())
	return rand.Intn(6*d.numberOfDice-1*d.numberOfDice) + 1
}

// func main() {
// 	numberOfDice := 2 // You can set the desired number of dice here
// 	dice := NewDice(numberOfDice)
// 	result := dice.rollDice()
// 	fmt.Println("Dice Roll Result:", result)
// }
