package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	multiThreadeddl("https://google.com", "threaded-google.txt")
}
func multiThreadeddl(url string, file string) {
	bufferChan := make(chan []byte)
	go downloadThreaded(bufferChan, url)
	f, err := os.Create("threaded-" + file)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	for data := range bufferChan {
		f.Write(data)
	}
}

// download a page and returns a byte array
func downloadThreaded(bufferChan chan []byte, url string) {

	resp, err := http.Get(url)
	if err != nil {
		bufferChan <- nil
	}

	defer resp.Body.Close()

	bodyBuff, err := io.ReadAll(resp.Body)
	if err != nil {
		bufferChan <- nil
	}
	bufferChan <- bodyBuff
	close(bufferChan)
}
