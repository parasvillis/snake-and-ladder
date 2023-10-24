package main

type Jumper struct {
	startPoint int
	endPoint   int
}

func NewJumper(startPoint, endPoint int) *Jumper {
	return &Jumper{
		startPoint: startPoint,
		endPoint:   endPoint,
	}
}

// func main() {
// 	// You can use the NewJumper function to create instances of the Jumper struct.
// }
