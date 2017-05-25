package main

import "time"

func test() {
	t := time.NewTicker(time.Second)
	select {
	case _, ok := <-t.C:
		
	}
}
