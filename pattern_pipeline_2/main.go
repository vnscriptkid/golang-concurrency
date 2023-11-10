package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func displayData(ic <-chan CapText) <-chan string {
	c := make(chan string)

	go func() {
		var wg sync.WaitGroup
		for x := range ic {
			wg.Add(1)
			go concatenateValue(x, c, &wg)
		}

		wg.Wait()

		close(c)
	}()

	return c
}

func concatenateValue(input CapText, c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- fmt.Sprintf("%s, %s, %s", input.Text, input.CapFirst, input.CapAll)
}

func prepareData(ic <-chan string) <-chan CapText {
	c := make(chan CapText)

	go func() {
		for s := range ic {
			c <- CapText{Text: s, CapFirst: capitalizeFirst(s), CapAll: capitalizeAll(s)}
		}

		close(c)
	}()

	return c
}

func generateData() <-chan string {
	c := make(chan string)

	go func() {
		file, err := os.Open("text.txt")

		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')

			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")

			c <- line

			if err == io.EOF {
				break
			}
		}

		close(c)
	}()

	return c
}

func main() {

	c := displayData(prepareData(generateData()))
	for data := range c {
		log.Printf("Items: %+v", data)
	}

}

// hello world
// CapFirst: Hello world
// CapAll: HELLO WORLD

type CapText struct {
	Text     string
	CapFirst string
	CapAll   string
}

func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func capitalizeAll(s string) string {
	return strings.ToUpper(s)
}
