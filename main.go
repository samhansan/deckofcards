package main

import (
    "bufio"
    "fmt"
    "github/deckofcards/card"
    "github/deckofcards/bad_ai"
    "os"
    "strconv"
)

func main() {
    fmt.Println("  ~~ War ~~ ")
    // Create new deck and shuffle it
    new_deck := card.New()
    shuffled := card.Shuffle(new_deck)

    // Player gets first 5 cards, AI gets next 5 cards
    var player_hand = shuffled[:5]
    var bad_ai_hand = shuffled[5:10]
    var player_score int = 0
    var bad_ai_score int = 0

    // Keep playing until all cards are gone from players' hands
    for len(player_hand) > 0 {
        fmt.Println("Score   ", player_score, " : ", bad_ai_score)
        fmt.Println("~~~~~~~~~~~~~~~~~~~~")
        fmt.Println("Your Hand:")
        card.PrintCards(player_hand)
        fmt.Println("\nBad_AI's Hand:")
        card.PrintCards(bad_ai_hand)
        fmt.Println("")

        player_input := PlayerChoose(player_hand)
        player_choice := player_hand[player_input]
        bad_ai_choice_index := bad_ai.Choose(bad_ai_hand)
        bad_ai_choice := bad_ai_hand[bad_ai_choice_index]

        fmt.Println("You chose:")
        fmt.Println("  ", player_choice)
        fmt.Println("Bad_AI chose:")
        fmt.Println("  ", bad_ai_choice)
        fmt.Println("")

        player_point, bad_ai_point := GivePoint(player_choice, bad_ai_choice)
        player_score += player_point
        bad_ai_score += bad_ai_point

        player_hand = card.Remove(player_hand, player_input)
        bad_ai_hand = card.Remove(bad_ai_hand, bad_ai_choice_index)
    }

    fmt.Println("~~~~~~~~~~~~~~~~~~~~\n~~~~~~~~~~~~~~~~~~~~")
    if player_score > bad_ai_score {
        fmt.Println("You beat Bad_AI!  You won the War!")
    } else {
        fmt.Println("You lost to Bad_AI :( Better luck next time!")
    }
}

// Player inputs an integer between 1 and len(player_hand), returns appropriate index of card in player_hand
func PlayerChoose(player_hand []card.Card) int {
    for {
        fmt.Print("What card do you choose? (1-", len(player_hand), "):")
        input := bufio.NewScanner(os.Stdin)
        for input.Scan() {
            choice, err := strconv.Atoi(input.Text())
            if err != nil {
                fmt.Println(err)
                continue
            }
            if (choice > len(player_hand)) || (choice < 1) {
                fmt.Println("Not a valid choice!")
                continue
            }
            return choice - 1
        }
    }
}

// compares player and bad_ai cards, returns point accordingly
func GivePoint(player_card, bad_ai_card card.Card) (int, int) {
    // tie goes to bad_ai
    if card.Smaller(bad_ai_card, player_card) {
        fmt.Println("You Win!\n")
        return 1, 0
    }
    fmt.Println("Bad_AI Wins!\n")
    return 0, 1
}