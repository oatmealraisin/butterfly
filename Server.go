/* * *
 * Server will handle the outbound and inbound requests, maybe our main driver?
 * Gets a database right now...
 */

package bf

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
*/


//declare a handler, adds functionality to DefaultServeMux
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

//log.Fatal(http.ListenAndServe(":8080", nil))


// Takes a database for putting the cards
// Returns the error
func StartServer(db Database, comms chan string) {

	s := &http.Server{
		Addr:           ":1337",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
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
