package main

import (
	"math/rand"
	"time"

	"github.com/sigsant/acey_ducey/pkg/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	game.StartGame()

}
