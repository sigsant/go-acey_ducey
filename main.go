package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sigsant/acey_ducey/pkg/card"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	firstCard, secondCard := card.ReturnCard()
	fmt.Println("Datos de la primera carta", firstCard.ToString())
	fmt.Println("Datos de la segunda carta", secondCard.ToString())
}
