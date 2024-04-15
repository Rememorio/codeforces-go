// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1914/problem/F
// https://codeforces.com/problemset/status/1914/problem/F
func Test_cf1914F(t *testing.T) {
	testCases := [][2]string{
		{
			`6
4
1 2 1
2
1
5
5 5 5 1
7
1 2 1 1 3 3
7
1 1 3 2 2 4
7
1 2 1 1 1 3`,
			`1
0
1
3
3
3`,
		},
		{
			`1
7
1 1 3 2 2 4`,
			`3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1914F)
}
