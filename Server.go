/* * *
 * Server will handle the outbound and inbound requests, maybe our main driver?
 * Gets a database right now...
 */

package bf

// Takes a database for putting the cards
// Returns the error
func StartServer(db Database, comms chan string) uint {

}

func AddModule(p Plugin) {
	// Add the plugin/module to our cycle thing, so that we check it
}

/*

 */
func getNewPosts(c chan bf.Card, p bf.Plugin) {
	for index, card := range p.Cards {
		c <- card
	}
}
