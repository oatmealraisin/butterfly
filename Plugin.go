/*
TODO: Comment on package
*/
package bf

type Plugin interface {
	Cards() []Card
	GetCards() bool
	PostCards(Card) bool
}

func Refresh(p []Plugin) {
	for _, plug := range p {
		go plug.GetCards()
	}
}
