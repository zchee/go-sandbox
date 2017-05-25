package main

import "time"

func test() {
	t := time.NewTicker(time.Second)
	var ok bool
	select {
	case _, ok = <-t.C:
		
	}
}
