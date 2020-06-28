package bad_ai

import (
    "github/deckofcards/card"
)

// bad_ai currently just chooses the largest card from []Card
func Choose(cards []card.Card) (int) {
    return card.CardMax(cards)
}