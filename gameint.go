package main

import "fmt"

type GameCharacter interface {
	Attack() // must have an attack type
}

type Knight struct{} // Knight struct
type Wizard struct{} // Wizard struct
type Dragon struct{} // Dragon struct

// Knight attack method
func (k Knight) Attack() {
	fmt.Println("Knight is attacking")
}

func (w Wizard) Attack() {
	fmt.Println("Wizard is attacking")
}

func (d Dragon) Attack() {
	fmt.Println("Dragon is attacking")
}

func StartGame(char GameCharacter) {
	fmt.Println("Game is starting")
	char.Attack()
}

func main() {
	knight := Knight{}
	wizard := Wizard{}
	dragon := Dragon{}
	StartGame(knight)
	StartGame(wizard)
	StartGame(dragon)
}
