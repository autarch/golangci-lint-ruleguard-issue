package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func SomeRule(m dsl.Matcher) {
	m.Match("$foo").
		Where(m["foo"].Text.Matches("x")).
		Where(m["foo"].Text.Matches("y")).
		Report("x and y")
}
