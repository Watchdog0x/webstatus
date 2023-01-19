package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	red := "\033[31m"
	green := "\033[32m"
	reset := "\033[0m"
	if err != nil {
		fmt.Printf("%s%-40s DOWN %s%s\n", red, url, err, reset)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s%-40s %-10d%s\n", green, url, resp.StatusCode, reset)
	}
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	urls := []string{"https://mail.d-fender.net", "https://psono.d-fender.net", "https://d-fender.net"}

	wg.Add((len(urls)))

	fmt.Printf("URL %s Status code\n", strings.Repeat(" ", 36))
	fmt.Println(strings.Repeat("\u2500", 55))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}
	wg.Wait()
}
