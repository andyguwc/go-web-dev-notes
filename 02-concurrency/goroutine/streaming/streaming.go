/* 

Streaming data using a single channel 

Using for <elemem> := range <channel>{...} to receive data 
*/


package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink from its water bowl",
		"The dark bird of prey lands on a small tree after hunting for fish",
	}

	histogram := make(map[string]int)
	wordsCh := make(chan string)

	// splits line and emit words to channel
	// The sender goroutine loops through the text line and sends a word at a time. 
	// It then blocks until the word is received by the receiving (main) goroutine.
	go func() {
		defer close(wordsCh) // closes channel upon fn return
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				wordsCh <- word
			}
		}
	}()

	// process word stream and count words
	// loop until wordsCh is closed

	for {
		word, opened := <-wordsCh // Pulls the data from the channel
		if !opened { // Checks the open status of the channel
			break // If closed, break out of the loop
		}
		histogram[word]++ // Otherwise record histogram
	}

	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}
