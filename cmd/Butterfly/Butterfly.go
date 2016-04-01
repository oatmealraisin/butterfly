/*
	TODO: package comment
*/

package main

import (
	"fmt"
	"time"

	"github.com/oatmealraisin/butterfly"
)

func displayPosts(c chan bf.Card) {
	for {
		select {
		case card := <-c:
			fmt.Println(card.Username)
			fmt.Println(card.Message)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	comms := make(chan string, 100)
	c := make(chan bf.Card, 100)
	db := bf.Database{MaxSize: 100}
	go bf.StartServer(db, comms)
	var p bf.Plugin
	p = bf.Twitter{}
	bf.AddModule(p)

	fmt.Println("Success!")

	go displayPosts(c)
}
