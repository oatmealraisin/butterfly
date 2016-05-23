/*
	TODO: Description
*/

package bf

import (
	"time"
)

type Card struct {
	Username   string
	Message    string
	TimeStamp  time.Time
	PluginType int
	id         int
}

// Returns a JSON string (or object?) representing the card
func ToJSON(card Card) {
	result string

	// TODO:

	return result
}

// Creates a new card from a JSON string (or object?)
func NewCard(card string) {
	result Card

	// TODO: 

	return result
}
