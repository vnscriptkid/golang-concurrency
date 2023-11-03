package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://educative.io",
	"https://educative.io/teach",
	"https://educative.io/assessments",
	"https://educative.io/projects",
	"https://educative.io/paths",
	"https://educative.io/learning-plans",
	"https://educative.io/learn",
	"https://educative.io/edpresso",
	"https://educative.io/explore",
	"https://educative.io/efer-a-friend",
	"https://google.com",
	"https://twitter.com",
}

func fetchUrl(url string, wg *sync.WaitGroup) {
	// Write your code here
	if wg != nil {
		defer wg.Done()
	}

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s => status=%s\n", url, res.Status)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Write your code here
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, &wg)
	}

	wg.Wait()

	fmt.Fprintf(w, "Welcome")
}

func handleRequests() {
	// Write your code here

	http.HandleFunc("/", homeHandler)

	if err := http.ListenAndServe(":7079", nil); err != nil {
		fmt.Println("failed to listen and serve: ", err)
	}

}

func main() {

	handleRequests()
}
