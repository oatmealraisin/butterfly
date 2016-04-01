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
	cards = append(cards, Card{Username: "Twitter", Message: "interface", TimeStamp: time.Now()})
	return cards
}

func (plug Twitter) PostCards(c Card) bool {
	return true
}

func (plug Twitter) GetCards() bool {
	return true
}
