// http://qiita.com/ryurock/items/b0466ad144f5e0555a95

// alert: a bluemonday sample with sanitize the alert.
package main

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

func main() {
	p := bluemonday.UGCPolicy()
	html := p.Sanitize(
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
	)

	// Output:
	// <a href="http://www.google.com" rel="nofollow">Google</a>
	fmt.Println(html)
}
