/* * *
 * Server will handle the outbound and inbound requests, maybe our main driver?
 * Gets a database right now...
 */

package bf

// Takes a database for putting the cards
// Returns the error
func StartServer(db Database, comms chan string) {

}

func AddModule(p Plugin) {
	// Add the plugin/module to our cycle thing, so that we check it
}

/*

*/
func getNewPosts(c chan Card, p Plugin) {
	for _, card := range p.Cards() {
		c <- card
	}
}
