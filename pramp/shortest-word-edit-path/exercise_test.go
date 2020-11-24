package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		beginWord string
		endWord   string
		wordList  []string
		output    int
	}{
		{"bit", "dog", []string{"but", "put", "big", "pot", "pog", "dog", "lot"}, 5},
		{"no", "go", []string{"to"}, -1},
	}

	implementations = []struct {
		name string
		f    func(beginWord string, endWord string, wordList []string) int
	}{
		{"shortestWordEditPath", shortestWordEditPath},
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
						result := f(ca.beginWord, ca.endWord, ca.wordList)
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
					f(c.beginWord, c.endWord, c.wordList)
				}
			}
		})
	}
}
