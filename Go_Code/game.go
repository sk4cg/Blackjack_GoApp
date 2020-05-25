package main

import (
	"github.com/sk4cg/Blackjack_GoApp/Go_Code/deck"
	"fmt"
	"strings"
	"os"
)

type CardDealtTo string

const (
	dealer CardDealtTo = "dealer"
	player CardDealtTo = "player"
)

/*
Function Name : playerChoice
Funtion Purpose : Used to store player's input
Params :
	- choice ::
Return Type : string
	- userInput :: store player's input
*/
func playerChoice(choice string) string {
	//prints out the message to the user to make their decision
	fmt.Print(choice)
	var userInput string
	//stores the input in userInput string
	fmt.Scanln(&userInput)
	return userInput
}

/*
Function Name : playerInput
Funtion Purpose : act on the choice of the player to hit, stop, or abort the game (zero)
Params : None
Return Type : int
	- playNum :: keeps track of what player choice to invoke
*/
func playerInput() int {
	inputAprroved := false //checks to see if user inoput is acceptable
	input := "" //stores actual inputted characters
	playNum := 3 //keeps track of what player choice to invoke in later operations (1 = Hit and 2 = Stop)
	for inputAprroved == false {
		input = playerChoice("Hit [H] or Show [S] or Zero [0] to end :: ")
		//checks if user wants to hit
		if input == "H" || input == "h" {
			//inputAprroved shows that the user input has been accepted and put into playerDecision
			inputAprroved = true
			playNum = 1
		}
		//checks if user wants to stop
		if input == "S" || input == "s" {
			//inputAprroved shows that the user input has been accepted and put into playerDecision
			inputAprroved = true
			playNum = 2
		}
		//checks if user wants to abort/kill game
		if input == "0" {
			fmt.Printf("Got Kill signal. Aborting...\n")
			//exits game
			os.Exit(1)
		}
		if inputAprroved == false {
			//inputAprroved shows that the user input has NOT been accepted and the user is asked to input again
			fmt.Println("Player input is incorrect. Hit [H] or Show [S] or Zero [0] to end.")
		}
	}
	return playNum
}

/*
Function Name : cardsValue
Funtion Purpose : Used to store player's input
Params :
	- hand :: stores the deck of cards in current play
Return Type : int
	- cValue :: total card value
*/
func cardsValue(hand []deck.Card) int {

	cValue := 0 //card value total
	numAces := 0 //keeps track of the number of aces in the hand to decide to either appoint them as 11 points or 1 point

	//iterates through the hand to find amount of Aces
	for _, card := range hand {
		if card.Value.Name == "Ace" {
			//increments amount of Aces when one is found by 1
			numAces = numAces + 1
		} else {
			//increments total point amount when facecard is found by 10
			if card.Facecard() {
				cValue = cValue + 10
				//increments total point amount when a numbered card is found by respective value of the number
			} else {
				cValue = cValue + card.Value.Value
			}
		}
	}
	//adds up all Ace amounts until there are none left in the hand accounted for
	if numAces > 0 {
		for numAces > 1 {
			cValue = cValue + 1
			numAces = numAces - 1
		}
		//chooses whether to add 11 or 1 based on if hand might exceed 21
		if cValue + 11 <= 21 {
			cValue = cValue + 11
		} else {
			cValue = cValue + 1
		}
	}
	return cValue
}

/*
Function Name : displayPlayerCards
Funtion Purpose : Prints cards held by the player
Params :
	- hand :: stores the deck of cards of the player
Return Type : None
*/
func displayPlayerCards(hand []deck.Card) {
	fmt.Printf("\n*** Player Hand :: ***\n")
	displayCards(hand, player, false)
	fmt.Printf("\n *** Total  = %d *** \n\n", cardsValue(hand))
}

/*
Function Name : displayDealerCards
Funtion Purpose : Prints cards held by the dealer
Params :
	- hand :: stores the deck of cards of the dealer
	- isMysteryCard :: boolean value determining if the card should be hidden from the player
Return Type : None
*/
func displayDealerCards(hand []deck.Card, isMysteryCard bool) {
	if isMysteryCard {
		fmt.Printf("*** Dealer Hand :: ***\n")
		dealerHand := hand[1:]
		displayCards(dealerHand, dealer, true)
		fmt.Printf("\n")
	} else {
		fmt.Printf("\n*** Dealer Hand :: ***\n")
		displayCards(hand, dealer, false)
		fmt.Printf("\n")
	}
}

/*
Function Name : displayCards
Funtion Purpose : Prints cards in ASCII character art
Params :
	- hand :: stores the deck of cards
	- cardDealt :: enum to differentiate between dealer and player
	- firstCardHidden :: boolean value determining if the first card should be hidden from the player
Return Type : None
*/
func displayCards(hand []deck.Card, cardDealt CardDealtTo, firstCardHidden bool) {
	var name, symbol string
	spacing := " "
	var asciiL1, asciiL2, asciiL3, asciiL4, asciiL5, asciiL6, asciiL7, asciiL8, asciiL9 string

	if firstCardHidden == true && cardDealt == dealer {
		asciiL1 += "┌─────────┐"
		asciiL2 += "│░░░░░░░░░│"
		asciiL3 += "│░░░░░░░░░│"
		asciiL4 += "│░░░░░░░░░│"
		asciiL5 += "│░░░░░░░░░│"
		asciiL6 += "│░░░░░░░░░│"
		asciiL7 += "│░░░░░░░░░│"
		asciiL8 += "│░░░░░░░░░│"
		asciiL9 += "└─────────┘"
	}

	for _, card := range hand {
		if card.Facecard() {
			name = fmt.Sprintf(" %c", card.Value.Name[0])
			symbol = fmt.Sprintf("%s", card.Types.Symbol)
		} else {
			name = fmt.Sprintf("%2d", card.Value.Value)
			symbol = fmt.Sprintf("%s", card.Types.Symbol)
		}
		nameLineTop := fmt.Sprintf("│%v%v      │", name, spacing)
		symbolLine := fmt.Sprintf("│    %v    │", symbol)
		nameLineBottom := fmt.Sprintf("│      %v%v│", name, spacing)

		asciiL1 = asciiL1 + "┌─────────┐"
		asciiL2 = asciiL2 + nameLineTop
		asciiL3 = asciiL3 + "│         │"
		asciiL4 = asciiL4 + "│         │"
		asciiL5 = asciiL5 + symbolLine
		asciiL6 = asciiL6 + "│         │"
		asciiL7 = asciiL7 + "│         │"
		asciiL8 = asciiL8 + nameLineBottom
		asciiL9 = asciiL9 + "└─────────┘"
	}
	result := strings.Join([]string{asciiL1, asciiL2, asciiL3, asciiL4, asciiL5, asciiL7, asciiL7, asciiL8, asciiL9}, "\n")
	fmt.Println(result)
}

func main() {
	var d deck.Deck
	for {
		d.Initialize()
		d.Shuffle()
		playerStops := "x"

		fmt.Println("=============================================================")
		fmt.Println("\t\t\tGAME START!")
		fmt.Println("=============================================================")

		playerHand, err := d.Draw(2)
		if err != nil {
		}
		dealerHand, err := d.Draw(2)
		if err != nil {
		}

		for playerStops == "x" {
			displayPlayerCards(playerHand)
			displayDealerCards(dealerHand, true)
			playerDecision := playerInput()
			if playerDecision == 1 {
				drawnCards, err := d.Draw(1)
				if err != nil {
				}
				playerHand = append(playerHand, drawnCards[0])
				if cardsValue(playerHand) > 21 {
					displayPlayerCards(playerHand)
					fmt.Println("## Oops,you lose. Better luck next time. ##\n")
					playerStops = "s"
				}
			}
			if playerDecision == 2 {
				playerStops = "s"
			}
		}
		if cardsValue(playerHand) > 21 {
			continue
		}
		dealerDone := false
		for dealerDone == false {
			if cardsValue(dealerHand) >= 17 {
				dealerDone = true
				continue
			}
			drawnCards, err := d.Draw(1)
			if err != nil {
			}
			dealerHand = append(dealerHand, drawnCards[0])
		}
		displayDealerCards(dealerHand, false)

		if cardsValue(dealerHand) > 21 {
			fmt.Println("## Congratulations! You won. ##\n")
			continue
		}

		if cardsValue(playerHand) > cardsValue(dealerHand) {
			fmt.Println("## Congratulations! You won. ##\n")
			continue
		}

		if cardsValue(playerHand) == cardsValue(dealerHand) {
			fmt.Println("## Whoa! It's a draw. ##\n")
			continue
		}

		if cardsValue(playerHand) < cardsValue(dealerHand) {
			fmt.Println("## Oops,you lose. Better luck next time. ##\n")
			continue
		}
		fmt.Println("!! There was a problem determining the winner. !!\n")
	}
	return
}
