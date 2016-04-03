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
