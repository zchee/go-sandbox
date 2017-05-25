package main

import "time"

func test() {
	t := time.NewTicker(time.Second)
	select {
	case how, ok := <-t.C:
		
	}
}
