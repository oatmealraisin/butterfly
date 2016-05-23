/*
This will connect to the server as a client when all is said and done.
*/

package main

import "net/http"

func main() {
	if resp, err := client.Get("127.0.0.1:1337"); err != nil {
		// TODO: Handle Error
	}

	if req, err := http.NewRequest("GET", "127.0.0.1:1337", nil); err == nil {

	} else {
		// TODO: Handle Error
	}

	//	if resp, err := client.Do(req); err == nil {
	//		// The client must close the response body when finished with it
	//		resp.Body.Close()
	//		// TODO: Handle response
	//	} else {
	//		// TODO: Handle Error
	//	}
}
