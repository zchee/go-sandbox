// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
)

var (
	enc = flag.Bool("encode", true, "encoding arg")
	dec = flag.Bool("decode", false, "decoding arg")
)

func main() {
	flag.Parse()
	s := flag.Arg(0)

	switch {
	case *dec:
		u, err := url.QueryUnescape(s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(u)
	case *enc:
		u := url.QueryEscape(s)
		fmt.Print(u)
	}
}
