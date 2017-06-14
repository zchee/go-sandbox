// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fake sw_vers command for Intel Parallel Studio XE 2018 does not support macOS 10.13 (High Sierra).

package main

import "fmt"

func main() {
	fmt.Print("10.12")
}
