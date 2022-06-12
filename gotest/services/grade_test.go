package services_test

import (
	"testing"

	"github.com/sing3demons/gotest/services"
)

func TestCheckGrade(t *testing.T) {
	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "success grade A", score: 85, expected: "A"},
		{name: "success grade B", score: 75, expected: "B"},
		{name: "success grade C", score: 66, expected: "C"},
		{name: "success grade D", score: 59, expected: "D"},
		{name: "success grade F", score: 49, expected: "F"},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)
			expected := c.expected

			if grade != expected {
				t.Errorf("got %v expected %v", grade, expected)
			}
		})
	}

}

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}

}
