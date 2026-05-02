package main

import (
	"fmt"
)

type GameBoard struct {
	dice                   *Dice
	nextTurn               []*Player
	snakes, ladders        []*Jumper
	playersCurrentPosition map[string]int
	boardSize              int
}

func NewGameBoard(dice *Dice, nextTurn []*Player, snakes, ladders []*Jumper, playersCurrentPosition map[string]int, boardSize int) *GameBoard {
	return &GameBoard{
		dice:                   dice,
		nextTurn:               nextTurn,
		snakes:                 snakes,
		ladders:                ladders,
		playersCurrentPosition: playersCurrentPosition,
		boardSize:              boardSize,
	}
}

func (gb *GameBoard) startGame() {
	for len(gb.nextTurn) > 1 {
		player := gb.nextTurn[0]
		gb.nextTurn = gb.nextTurn[1:]
		currentPosition := gb.playersCurrentPosition[player.getPlayerName()]
		diceValue := gb.dice.rollDice()
		nextCell := currentPosition + diceValue

		if nextCell > gb.boardSize {
			gb.nextTurn = append(gb.nextTurn, player)
		} else if nextCell == gb.boardSize {
			fmt.Println(player.getPlayerName(), "won the game")
		} else {
			nextPosition := nextCell
			b := false

			for _, v := range gb.snakes {
				if v.startPoint == nextCell {
					nextPosition = v.endPoint
				}
			}

			if nextPosition != nextCell {
				fmt.Println(player.getPlayerName(), "Bitten by Snake present at:", nextCell)
			}

			for _, v := range gb.ladders {
				if v.startPoint == nextCell {
					nextPosition = v.endPoint
					b = true
				}
			}

			if nextPosition != nextCell && b {
				fmt.Println(player.getPlayerName(), "Got ladder present at:", nextCell)
			}

			if nextPosition == gb.boardSize {
				fmt.Println(player.getPlayerName(), "won the game")
			} else {
				gb.playersCurrentPosition[player.getPlayerName()] = nextPosition
				fmt.Println(player.getPlayerName(), "is at position", nextPosition)
				gb.nextTurn = append(gb.nextTurn, player)
			}
		}
	}
}

// func main() {
// 	// Assuming you have initialized all the required variables (dice, nextTurn, snakes, ladders, playersCurrentPosition, and boardSize)
// 	gameBoard := NewGameBoard(dice, nextTurn, snakes, ladders, playersCurrentPosition, boardSize)
// 	gameBoard.startGame()
// }

// You'll need to define the corresponding Dice, Player, and Jumper types and methods separately.
