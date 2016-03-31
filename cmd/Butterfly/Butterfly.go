/*
	TODO: package comment
*/

package main

import (
	"fmt"
	"time"

	"github.com/oatmealraisin/bf"
)

func displayPosts(c chan bf.Card) {
	for {
		select {
		case card := <-c:
			fmt.Println(card.Username)
			fmt.Println(card.Message)
		default:
			bf.Refresh(p)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	comms := make(chan string, 100)
	db := Database{MaxSize: 100}
	go bf.StartServer(db, comms)
	var p bf.Plugin
	p = bf.Twitter{}
	bf.AddModule(p)

	go displayPosts(c)
}
