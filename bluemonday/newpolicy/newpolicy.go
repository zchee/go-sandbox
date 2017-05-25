// http://qiita.com/ryurock/items/b0466ad144f5e0555a95

// newpolicy: a bluemonday sample with NewPolicy.
package main

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

func main() {
	p := bluemonday.NewPolicy()

	p.AllowElements("li").AllowElements("ul")

	html := p.Sanitize(
		`<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E5.AD.97.E6.BA.90"><span class="tocnumber">1.1</span> <span class="toctext">字源</span></a></li>
<li class="toclevel-2 tocsection-3"><a href="#.E9.9F.B3"><span class="tocnumber">1.2</span> <span class="toctext">音</span></a></li>
<li class="toclevel-2 tocsection-4"><a href="#.E6.84.8F.E7.BE.A9"><span class="tocnumber">1.3</span> <span class="toctext">意義</span></a>
</ul>
`,
	)

	fmt.Println(html)
}
