// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1931/problem/D
// https://codeforces.com/problemset/status/1931/problem/D
func Test_cf1931D(t *testing.T) {
	testCases := [][2]string{
		{
			`7
6 5 2
1 2 7 4 9 6
7 9 5
1 10 15 3 8 12 15
9 4 10
14 10 2 2 11 11 13 5 6
9 5 6
10 7 6 7 9 7 7 10 10
9 6 2
4 9 7 1 2 2 13 3 15
9 2 3
14 6 1 15 12 15 8 2 15
10 5 7
13 3 3 2 12 11 3 7 13 14`,
			`2
0
1
3
5
7
0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1931D)
}
