// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/23/C
// https://codeforces.com/problemset/status/23/problem/C
func Test_cf23C(t *testing.T) {
	testCases := [][2]string{
		{
			`2
2
10 15
5 7
20 18
1
0 0`,
			`YES
1 3
YES
1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf23C)
}
