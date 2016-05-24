/*
	TODO: package comment
*/

// TODO: Use a logger

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var (
	XDG_CONFIG_HOME = os.Getenv("XDG_CONFIG_HOME")

	/* cl flags */

	port = flag.String("p", ":1337",
		"Specify the port that the server will listen on.")
	max_size = flag.Int("max_size", 1000,
		"The maximum amount of cards to keep in the database.")
)

func init() {
	// TODO: look for plugins
	// TODO: parse config
}

func main() {

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		// TODO: Handle error
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Using port", *port)

	http.HandleFunc("/", helpHandler)
	http.HandleFunc("/help/", helpHandler)
	http.HandleFunc("/exit/", exitHandler)
	http.HandleFunc("/cards/", cardHandler)
	log.Fatal(http.ListenAndServe(*port, nil))

}

/* Handler Section */

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
