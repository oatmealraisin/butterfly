/*
TODO
*/
package bf

type Twitter struct {
}

func (plug Twitter) Cards() Card {
	return Card{Username: "Twitter", Message: "interface", TimeStamp: time.Now()}
}

func (plug Twitter) PostCards(c Card) bool {
	return true
}

func (plug Twitter) getCards() bool {
	return true
}
