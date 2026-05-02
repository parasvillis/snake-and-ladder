package main

type Player struct {
	playerName string
	id         int
}

func NewPlayer(playerName string, id int) *Player {
	return &Player{
		playerName: playerName,
		id:         id,
	}
}

func (p *Player) getPlayerName() string {
	return p.playerName
}

// func main() {
// 	// You can use the NewPlayer function to create instances of the Player struct.
// }
