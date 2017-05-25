package main

import (
	"encoding/csv"
	"io"
	"os"
	"sync"
)

func recordImport(wg *sync.WaitGroup, q chan string) {
	defer wg.Done()
	for {
		record, ok := <-q
		if !ok {
			return
		}

		// インポート
	}
}

func main() {
	file, _ := os.Open(os.Args[0])
	defer file.Close()

	reader := csv.NewReader(file)
	var wg sync.WaitGroup

	q := make(chan string, 5)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go recordImport(&wg, q)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		q <- record
	}
	close(q)

	wg.Wait()
}
