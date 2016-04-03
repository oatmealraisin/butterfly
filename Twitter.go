/*
TODO
*/
package bf

import (
	"time"
)

type Twitter struct {
}

func (plug Twitter) Cards() []Card {
	var cards []Card

	card := Card{
		Username:  "Twitter",
		Message:   "interface",
		TimeStamp: time.Now(),
	}
	cards = append(cards, card)

	return cards
}

func (plug Twitter) PostCards(c Card) bool {
	return true
}

func (plug Twitter) GetCards() bool {
	return true
}
