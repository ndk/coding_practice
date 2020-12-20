package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		s      string
		output string
	}{
		{"abc", "abc"},
		{"abcd", "abcd"},
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
		{"F R  I   E    N     D      S      ", "S D N E I R F"},
	}

	implementations = []struct {
		name string
		f    func(s string) string
	}{
		{"reverseWords space O(1)", reverseWords},
		{"reverseWords_split", reverseWords_split},
		{"reverseWords_double", reverseWords_double},
		{"reverseWords_deque", reverseWords_deque},
	}
)

func TestExercise(t *testing.T) {
	Convey("Test", t, func() {
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				for _, impl := range implementations {
					f := impl.f
					Convey(impl.name, func() {
						result := f(ca.s)
						So(result, ShouldResemble, ca.output)
					})
				}
			})
		}
	})
}

func BenchmarkExercise(b *testing.B) {
	for _, implementation := range implementations {
		f := implementation.f
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range cases {
					f(c.s)
				}
			}
		})
	}
}
