// http://qiita.com/ryurock/items/b0466ad144f5e0555a95

// scheme: a bluemonday sample with sanitize of li without scp scheme.
package main

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

func main() {
	p := bluemonday.NewPolicy()

	p.AllowElements("li").AllowElements("ul")
	p.AllowStandardURLs()
	p.AllowAttrs("href").OnElements("a")
	html := p.Sanitize(
		`<ul>
<li class="toclevel-2 tocsection-2"><a href="scp://wiki.jp/hoge"><span class="tocnumber">1.1</span> <span class="toctext">字源</span></a></li>
<li class="toclevel-2 tocsection-3"><a href="scp://wiki.jphoge"><span class="tocnumber">1.2</span> <span class="toctext">音</span></a></li>
<li class="toclevel-2 tocsection-4"><a href="https://google.com"><span class="tocnumber">1.3</span> <span class="toctext">意義</span></a></li>
</ul>
`,
	)

	fmt.Println(html)
}
