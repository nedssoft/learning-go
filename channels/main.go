package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	var c = make(chan string)

	 for _, link := range links {
		go checkLink(link, c)
	}

	 for l := range c {
		go func(lnk string) {
			time.Sleep(5 * time.Second)
			checkLink(lnk, c)
		}(l)
	}
	
}


func checkLink(link string, c chan string) {
   _, err := http.Get(link)
	 if  err != nil {
	  fmt.Printf("%s seems down", link)
		return
	}
	c <- link
	fmt.Printf("%s is up\n", link)
}
