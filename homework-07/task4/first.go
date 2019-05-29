// task4
// Author: Karpov A. mailto:art.frela@gmail.com
// Date: 2019-05-22
// first response from site mirrors

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	sites := []string{"https://www.google.cz/", "https://www.google.ru/", "https://www.google.ua/"}
	fmt.Println(mirroredQuery(sites))
	time.Sleep(3 * time.Second) // pause to wait for execution goroutines
}

// mirroredQuery - returns fastest response
func mirroredQuery(sites []string) string {
	responses := make(chan string, 3)
	for _, site := range sites {
		go func(site string) {
			responses <- request(site)
		}(site)
	}
	return <-responses // возврат самого быстрого ответа
}

// request - make request to the http and returns some data
func request(url string) string {
	t := time.Now()
	r, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	if r.StatusCode == http.StatusOK {
		cooks := r.Cookies()
		domain := strings.Split(fmt.Sprintf("%v", cooks[0]), ";")
		log.Printf("%v [%v]", domain[2], time.Since(t)) // logging for approve
		return fmt.Sprintf("%v [%v]", domain[2], time.Since(t))
	}
	return fmt.Sprintf("fault request, %d %s", r.StatusCode, r.Status)
}
