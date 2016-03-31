/*
TODO: Comment on package
*/
package bf

type Plugin interface {
	Cards() []Card
	getCards() bool
	PostCards(Card) bool
}

func Refresh(p []Plugin) {
	for index, plug := range p {
		go p.getCards()
	}
}
