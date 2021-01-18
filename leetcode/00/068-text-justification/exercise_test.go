package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		words    []string
		maxWidth int
		output   []string
	}{
		{
			[]string{"This", "is", "an", "example", "of", "text", "justification."},
			16,
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
			16,
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	implementations = []struct {
		name string
		f    func(words []string, maxWidth int) []string
	}{
		{"fullJustify", fullJustify},
		{"fullJustify2", fullJustify2},
		{"fullJustify3", fullJustify3},
		{"fullJustify_leet", fullJustify_leet},
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
						result := f(ca.words, ca.maxWidth)
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
					f(c.words, c.maxWidth)
				}
			}
		})
	}
}
