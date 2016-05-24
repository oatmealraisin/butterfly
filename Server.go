package bf

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Takes a database for putting the cards
// Returns the error
func StartServer(port string) error {
	fmt.Println(port)

	http.HandleFunc("/", helpHandler)
	http.HandleFunc("/help/", helpHandler)
	http.HandleFunc("/exit/", exitHandler)
	http.HandleFunc("/cards/", cardHandler)
	log.Fatal(http.ListenAndServe(port, nil))

	// Wrote this to create a custom server, abandoned due to difficulties adding
	// a handler function
	//server := &http.Server{
	//	Addr:           port,
	//	Handler:        handler,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//	// TODO: Check out connstate, pretty cool looking
	//}
	//
	//server.HandleFunc("/", handler)
	//log.Fatal(server.ListenAndServe())

	return nil
}

/**
Handler Section
*/

// Default handler for root requests.  Maybe prints a help page?
// TODO: Print a help page
// TODO: Create a help html page
func helpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: Help page")
}

// Handler for when we get an http request asking for cards.  Will return
// json of each n card, where n is /card/n if n is NaN, return error
func cardHandler(w http.ResponseWriter, r *http.Request) {

	if number, err := strconv.Atoi(r.URL.Path[len("/cards/"):]); err != nil {
		fmt.Fprintf(w, "Please specify, with an int, how many cards you need.")
		return
	} else {
		// TODO: Grab top n cards from database
		fmt.Fprintf(w, "Hello! %d", number)
	}
}

// TODO: Description
func exitHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
