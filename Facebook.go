/*
TODO Description
*/

package bf

type Facebook struct {
}

func (plug Facebook) Cards() Card {
	return Card{
		Username:  "facebook",
		Message:   "interface",
		TimeStamp: time.Now()
	}
}

func (plug Facebook) PostCards(c Card) bool {
	return true
}

func (plug Facebook) getCards() bool {

	return true
}
