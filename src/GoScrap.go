package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly" // import colly
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments")
	}

	fmt.Println("GoScrap v1.3 by ghosface-engineer")
	c := colly.NewCollector()

	website := os.Args[1]

	if len(os.Args) == 3 {
		timeout, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Example: ./GoScrap <website> <proxy> <timeout>")
			log.Fatal(err)
		}
		c.SetRequestTimeout(time.Duration(timeout) * time.Second)
	}

	if len(os.Args) == 4 {
		proxy := os.Args[3]
		c.SetProxy(proxy)
	}

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Page Title:", e.Text)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit(website)

}
