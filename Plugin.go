/*
TODO: Comment on package
*/
package bf

type Plugin interface {
	Cards() Card // TODO Implement this as a slice
	GetCards() bool
	PostCards(Card) bool
}

func Refresh(p []Plugin) {
	for index, plug := range p {
		go p.GetCards()
	}
}
