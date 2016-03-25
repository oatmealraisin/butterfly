package main

import (
	"fmt"
	"time"
)

type Card struct {
	Username string
	id       uint
	Message  string
}

func getNewPosts(c chan Card) {
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
