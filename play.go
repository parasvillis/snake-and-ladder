package main

func main() {
	dice := NewDice(1)
	p1 := NewPlayer("player1", 1)
	p2 := NewPlayer("player2", 2)
	allPlayers := []*Player{p1, p2}
	snake1 := NewJumper(10, 2)
	snake2 := NewJumper(99, 12)
	snakes := []*Jumper{snake1, snake2}
	ladder1 := NewJumper(5, 25)
	ladder2 := NewJumper(40, 89)
	ladders := []*Jumper{ladder1, ladder2}
	playersCurrentPosition := make(map[string]int)
	playersCurrentPosition["player1"] = 0
	playersCurrentPosition["player2"] = 0
	gb := NewGameBoard(dice, allPlayers, snakes, ladders, playersCurrentPosition, 100)
	gb.startGame()
}

// Define the corresponding Dice, Player, Jumper, and GameBoard types and methods separately.
