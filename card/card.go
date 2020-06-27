package card

import (
    "fmt"
    "math/rand"
    "time"
)

type Suit uint8

const (
    Spade Suit = iota
    Club
    Diamond
    Heart
)

var suits = [...]Suit{Spade, Club, Diamond, Heart}

func (s Suit) String() string {
    return [...]string{"Spade", "Club", "Diamond", "Heart"}[s]
}

type Rank uint8

const (
    Ace Rank = iota + 1
    Two
    Three
    Four
    Five
    Six
    Seven
    Eight
    Nine
    Ten
    Jack
    Queen
    King
)

func (r Rank) String() string {
    return [...]string{"", "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}[r]
}

const (
    minRank = Ace
    maxRank = King
)

type Card struct {
    Suit
    Rank
}

func New() []Card {
    var cards []Card
    for _, suit := range suits {
        for rank := minRank; rank <= maxRank; rank++ {
            cards = append(cards, Card{Suit: suit, Rank: rank})
        }
    }
    return cards
}

func (c Card) String() string {
    return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func Shuffle(cards []Card) []Card {
    shuffled := make([]Card, len(cards))
    r := rand.New(rand.NewSource(time.Now().Unix()))
    perm := r.Perm(len(cards))
    for i, j := range perm {
        shuffled[i] = cards[j]
    }
    return shuffled
}

func PrintCards(cards []Card) {
    for _, card := range cards {
        fmt.Println(card)
    }
    return
}