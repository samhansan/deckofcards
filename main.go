package main

import (
    "fmt"
    "card/card"
)

func main() {
    new_deck := card.New()
    shuffled := card.Shuffle(new_deck)
}