/*
	TODO: package comment
*/

package main

import (
	"fmt"
	"time"
)

type Plugin interface {
	Cards() Card
	PostCards(Card) bool
}

type Card struct {
	Username string
	id       uint
	Message  string
}

func getNewPosts(c chan Card) {
	time.Sleep(1000 * time.Millisecond)

	c <- Card{Username: "murphy", Message: "Asdf"}
	time.Sleep(300 * time.Millisecond)

	c <- Card{Username: "murphy", Message: "gjdhg"}
	time.Sleep(100 * time.Millisecond)

	c <- Card{Username: "murphy", Message: "jjhj"}
	time.Sleep(3000 * time.Millisecond)

	c <- Card{Username: "murphy", Message: "fdsa"}
	time.Sleep(1000 * time.Millisecond)

	c <- Card{Username: "murphy", Message: "hjdhgsj"}
}

func displayPosts(c chan Card) {
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
	c := make(chan Card, 100)
	go getNewPosts(c)
	go displayPosts(c)
	time.Sleep(10000 * time.Millisecond)
}
