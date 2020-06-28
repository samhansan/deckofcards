package card

import (
    "fmt"
    "math/rand"
    "sort"
    "time"
)

type Suit uint8

// defining the suits as constants, with successive integer values
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

// defining the ranks as constants, with successive integer values
// we don't want Ace to start at 0, so we do the +1
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

// create a new deck of cards, iterating through suits and ranks
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

// take a []Card and randomly order the cards and return the new []Card
func Shuffle(cards []Card) []Card {
    shuffled := make([]Card, len(cards))
    r := rand.New(rand.NewSource(time.Now().Unix()))
    perm := r.Perm(len(cards))
    for i, j := range perm {
        shuffled[i] = cards[j]
    }
    return shuffled
}

func RankSort(cards []Card) []Card {
    sort.Slice(cards, Less(cards))
    return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
    return func(cards []Card) []Card {
        sort.Slice(cards, less(cards))
        return cards
    }
}

//go through []Card, return index of largest value Card
func CardMax(cards []Card) int {
    if len(cards) == 0 {
            return -1
    }
    // start with first index, compare against all cards in []Card
    var max_index int = 0

    for index, card := range cards {
        if Smaller(cards[max_index], card) {
            max_index = index
        }
    }
    return max_index
}


func Less(cards []Card) func(i, j int) bool {
    return func(i, j int) bool {
        return absRank(cards[i]) < absRank(cards[j])
    }
}

//compare two cards, returns True if i < j
func Smaller(card_1, card_2 Card) bool {
    return absRank(card_1) < absRank(card_2)
}

// normally could use both rank and suit to make an absRank value, but I don't care about suit for now
func absRank(c Card) int {
    return int(c.Rank)
}

// remove a card from []Card at the specified index
func Remove(cards []Card, s int) []Card {
    return append(cards[:s], cards[s+1:]...)
}

// print all cards in []Card
func PrintCards(cards []Card) {
    for _, card := range cards {
        fmt.Println("  ", card)
    }
    return
}