package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		emails []string
		output int
	}{
		{
			[]string{
				"test.email+alex@leetcode.com",
				"test.e.mail+bob.cathy@leetcode.com",
				"testemail+david@lee.tcode.com",
			},
			2,
		},
		{
			[]string{
				"testemail@leetcode.com",
				"testemail1@leetcode.com",
				"testemail+david@lee.tcode.com",
			},
			3,
		},
		{
			[]string{
				"test.email+alex@leetcode.com",
				"test.email.leet+alex@code.com",
			},
			2,
		},
	}

	implementations = []struct {
		name string
		f    func(emails []string) int
	}{
		{"numUniqueEmails", numUniqueEmails},
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
						result := f(ca.emails)
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
					f(c.emails)
				}
			}
		})
	}
}
