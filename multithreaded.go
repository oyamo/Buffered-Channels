package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//	func main() {
//		multiThreadeddl("https://google.com", "threaded-google.txt")
//	}
func multiThreadeddl(url string, file string) {
	bufferChan := make(chan []byte)
	go downloadThreaded(bufferChan, url)
	f, err := os.Create("threaded-" + file)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	for data, ok := <-bufferChan; ok; {
		fmt.Println(data)
		f.Write(data)
	}
}

// download a page and returns a byte array
func downloadThreaded(bufferChan chan []byte, url string) {
	var bodybuff []byte
	resp, err := http.Get(url)
	if err != nil {
		bufferChan <- nil
	}

	defer resp.Body.Close()

	for b, r := io.ReadAtLeast(resp.Body, bodybuff, 1024); b == 0 && r != nil; {
		fmt.Println(" sending buffer")
		bufferChan <- bodybuff
		fmt.Println("sent buffer")
	}

	close(bufferChan)
}
