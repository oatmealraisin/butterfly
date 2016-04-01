/* * *
 * Server will handle the outbound and inbound requests, maybe our main driver?
 * Gets a database right now...
 */

package bf


pimport(
	"fmt"
	"net/http"
	"time"
)


// Takes a database for putting the cards
// Returns the error
func StartServer(db Database, comms chan string) {

	s := &http.Server{
		Addr:		":1337",
		Handler:	myHandler,
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
		MaxHeaderBytes:	1 << 20,
	}
	log.Fatal(s.ListenAndServe())
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
