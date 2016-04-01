/*
TODO Description
*/

package bf

import (
	"time"
)

type Facebook struct {
}

func (plug Facebook) Cards() Card {
	return Card{Username: "facebook", Message: "interface", TimeStamp: time.Now()}
}

func (plug Facebook) PostCards(c Card) bool {
	return true
}

func (plug Facebook) GetCards() bool {

	return true
}
