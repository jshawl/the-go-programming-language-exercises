package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	_, err = os.Stat("output.txt")
	if os.IsNotExist(err) {
		_, err := os.Create("output.txt")
		if err != nil {
			panic(err)
		}
	}
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0600)
	s := fmt.Sprintf("%.2fs %7d %s \n", secs, nbytes, url)
	if _, err = f.WriteString(s); err != nil {
		panic(err)
	}
}
