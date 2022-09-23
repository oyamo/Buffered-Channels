package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func singleThread(url string, file string) {
	body, err := download(url)
	if err != nil {
		safepanic(err)
	}

	err = save(body, file)
	if err != nil {
		safepanic(err)
	}
}

// download a page and returns a byte array
func download(page string) ([]byte, error) {
	resp, err := http.Get(page)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Save to storage
func save(body []byte, file string) error {
	return os.WriteFile(file, body, 0644)
}

// safepanic
func safepanic(err error) {
	errStr := fmt.Errorf("error: %v", err)
	fmt.Println(errStr)
	os.Exit(-1)
}
