package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fileName := *flag.String("file", "url.txt", "source url file")
	flag.Parse()
	b, e := os.ReadFile(fileName)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	slice := strings.Split(string(b), "\n")
	resCh := make(chan int)
	errCh := make(chan error)
	for _, url := range slice {
		go ping(url, resCh, errCh)
	}
	for range len(slice) {
		select {
		case err := <-errCh:
			fmt.Println(err)
		case res := <-resCh:
			fmt.Println(res)
		}
	}
}

func ping(url string, respCh chan int, errCh chan error) {
	res, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	respCh <- res.StatusCode
}
