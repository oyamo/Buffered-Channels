package main

import "testing"

type testItem struct {
	url  string
	file string
}

var tests = []testItem{
	{"https://google.com", "google.txt"},
	{"https://facebook.com", "facebook.txt"},
	{"https://airbnb.com", "airbnb.txt"},
}

func BenchmarkSingleThread(t *testing.B) {

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.B) {
			singleThread(tt.url, tt.file)
		})
	}
}

func BenchmarkMultiThreadeddl(t *testing.B) {

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.B) {
			multiThreadeddl(tt.url, tt.file)
		})
	}
}
