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
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
		{"hit", "cog", []string{"hot", "dot", "lot", "cog"}, 0},
	}

	implementations = []struct {
		name string
		f    func(beginWord string, endWord string, wordList []string) int
	}{
		{"ladderLength", ladderLength},
		{"ladderLength_leet1", ladderLength_leet1},
		{"ladderLength_leet2", ladderLength_leet2},
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
