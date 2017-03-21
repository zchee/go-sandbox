package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"sync"
)

var (
	wg  sync.WaitGroup
	mu  sync.Mutex
	url = "http://s3.zchee.io/darwin/SDK/MacOSX10.12.sdk.Xcode8_8S193k.tar.xz"
)

func main() {
	res, _ := http.Head(url)
	maps := res.Header
	length, _ := strconv.Atoi(maps["Content-Length"][0]) // Get the content length from the header request
	limit := 10                                          // 10 Go-routines for the process so each downloads 18.7MB
	len_sub := length / limit                            // Bytes for each Go-routine
	diff := length % limit                               // Get the remaining for the last request
	buf := make([][]byte, limit+1)                       // Make up a temporary array to hold the data to be written to the file

	for i := 0; i < limit; i++ {
		wg.Add(1)

		min := len_sub * i
		max := len_sub * (i + 1)

		var rangeHeader string
		if i == limit-1 {
			max += diff // Add the remaining bytes in the last request
			rangeHeader = "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max)
		} else {
			rangeHeader = "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1)
		}

		go func(min int, max int, i int) {
			fmt.Println(i)
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)

			req.Header.Add("Range", rangeHeader)
			resp, _ := client.Do(req)
			defer resp.Body.Close()

			mu.Lock()
			buf[i], _ = ioutil.ReadAll(resp.Body)
			mu.Unlock()
			wg.Done()
		}(min, max, i)
	}

	wg.Wait()
	ioutil.WriteFile(path.Base(url), bytes.Join(buf[:], nil), 0644)
}
