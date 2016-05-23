/*
	TODO: package comment
*/

package main

import (
	"flag"

	"github.com/oatmealraisin/butterfly"
)

func main() {
	// TODO: Initialize XDG_CONFIG_HOME directory, look for plugins, parse config

	// TODO: Move to init function
	var port = flag.String("p", ":1337",
		"Specify the port that the server will listen on.")

	bf.StartServer(*port)
}

func init() {

}
