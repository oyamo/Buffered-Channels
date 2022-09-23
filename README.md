# Buffered Channels Golang

`go 1.19`

This is implementation of lightweight threads (Goroutines) in Golang. The scripts download webpages concurrently and a benchmark is determined between a scrapper that implements Goroutines and one that does not.

## Benchmarking
```shell
go test -bench=.
```

## Building

```bash
make
```

## Cleaning

```bash
make clean
```