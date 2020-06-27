package main

import (
    "github/deckofcards/card"
)

func main() {
    new_deck := card.New()
    shuffled := card.Shuffle(new_deck)
    card.PrintCards(shuffled)
}