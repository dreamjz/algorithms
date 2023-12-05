package weekly372_test

import (
	weekly372 "leetcode/contest/weekly/372"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinimumOperations(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string 
		s1 string 
		s2 string 
		s3 string
		want int
	}{
		{
			name: "case#01",
			s1: "abc",
			s2: "abb",
			s3: "ab",
			want: 2,
		},
	}{
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := weekly372.FindMinimumOperations(tc.s1, tc.s2, tc.s3)
			
			assert.Equal(t, tc.want, res)
		})
	}
}