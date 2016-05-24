/*
	TODO: Description
*/

package bf

import (
	"encoding/json"
	"time"
)

type Card struct {
	Attachment string    // Filename for any attachments like pictures
	Username   string    // Username of the key user in the post
	Message    string    // Text body of the post
	TimeStamp  time.Time // Time when the post was created
	Type       string    // Website of origin for this post
	id         int       // Unique ID for the post.  Unique only to the website
}

// Returns a JSON string (or object?) representing the card
func ToJSON(card Card) []byte {
	if result, err := json.Marshal(card); err != nil {
		// TODO: Error
		return nil
	} else {
		return result
	}
}

func FromJSON(card []byte) (Card, error) {
	var result Card

	if err := json.Unmarshal(card, &result); err != nil {
		// TODO: Handle error
		return result, err
	}

	return result, nil
}
