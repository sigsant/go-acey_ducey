package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sigsant/acey_ducey/pkg/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\n\tAcey-Ducey is played in the following manner: ")
	fmt.Println("\n\tThe dealer (Computer) deals two cards face up.")
	fmt.Println("\tYou have an option to bet or not to bet depending on whether or not you feel the card will have a value between the  first two.")
	fmt.Println(("\tIf you do not want to bet, input an 0"))

	game.StartGame()

}
