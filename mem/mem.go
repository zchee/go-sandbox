// Copyright 2016 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mem

import "C"
import "unsafe"

func Cpy(dest, src []byte, count int) int {
	if count == 0 {
		return 0
	}
	C.memcpy(unsafe.Pointer(&dest[0]), unsafe.Pointer(&src[0]), C.size_t(count))
	return count
}

func Move(dest, src []byte, count int) int {
	if count == 0 {
		return 0
	}
	C.memmove(unsafe.Pointer(&dest[0]), unsafe.Pointer(&src[0]), C.size_t(count))
	return count
}
