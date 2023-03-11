package main

import (
	"fmt"
	"net/http"
)
func main() {
	links := [] string {
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}
	c := make(chan string)
	for _, l := range links {
		go checkStatus(l, c) //creating child go routines
	}

	//for sleeping go routines time.Sleep(5 * time.Second)

	// waiting for the value from the channel is blocking code
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	//for l := range c {
	//
	//}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)  //blocking code
	if err != nil {
		c <- "Down!"
		return;
	}
	c <- link + " is up!" //sending the data to channel
}