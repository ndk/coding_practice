package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		n      int
		k      int
		output []string
	}{
		{1, 1, []string{"0", "0"}},
		{1, 2, []string{"01", "10"}},
		{2, 2, []string{"00110", "01100"}},
		{4, 3, []string{
			"000222212122112111122201220212011202022101210211011102010102200120021001100200010000",
			"000010002001100120021002201010201110112012101220202110212022102221111211221212222000",
		}},
	}

	implementations = []struct {
		name string
		f    func(n int, k int) string
	}{
		{"crackSafe", crackSafe},
		{"crackSafe_IBWT", crackSafe_IBWT},
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
						result := f(ca.n, ca.k)
						So(result, ShouldBeIn, ca.output)
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
					f(c.n, c.k)
				}
			}
		})
	}
}
