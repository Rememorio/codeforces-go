// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/55/D
// https://codeforces.com/problemset/status/55/problem/D
func Test_cf55D(t *testing.T) {
	testCases := [][2]string{
		{
			`1
1 9`,
			`9`,
		},
		{
			`1
12 15`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf55D)
}