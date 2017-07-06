// Copyright 2017 The hiromifight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	for {
		cmd := exec.Command("say", "hiromi, fight!")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}
}
