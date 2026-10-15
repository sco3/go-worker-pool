# Go Worker Pool

A lightweight implementation of a worker pool pattern in Go. 
This repository demonstrates concurrent task processing using Go channels, 
worker goroutines, and synchronization via `sync.WaitGroup`.

The project showcases distributing a list of tasks 
across a constrained number of concurrent workers, 
safely collecting and returning the results. 

To run the application locally, simply execute:

```bash
go run main.go
```