// Generated by copypasta/template/generator_test.go
package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1732/D2
// https://codeforces.com/problemset/status/1732/problem/D2
func Test_cf1732D2(t *testing.T) {
	testCases := [][2]string{
		{
			`18
+ 1
+ 2
? 1
+ 4
? 2
+ 6
? 3
+ 7
+ 8
? 1
? 2
+ 5
? 1
+ 1000000000000000000
? 1000000000000000000
- 4
? 1
? 2`,
			`3
6
3
3
10
3
2000000000000000000
3
4`,
		},
		{
			`10
+ 100
? 100
+ 200
? 100
- 100
? 100
+ 50
? 50
- 50
? 50`,
			`200
300
100
100
50`,
		},
		{
			`7
? 4
? 4
+ 4
- 4
+ 2
+ 4
? 4`,
			`4
4
8`,
		},
		{
			`7
+ 2
? 2
- 2
+ 2
- 2
+ 2
? 2`,
			`4
4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1732D2)
}

func TestCompare_cf1732D2(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		q := rg.Int(1, 7)
		rg.NewLine()
		const U = 2
		has := [U+1]bool{}
		for ; q > 0; q-- {
			op := rg.IntOnly(0, 2)
			var x int
			if op < 2 {
				for {
					op = rg.IntOnly(0, 1)
					x = rg.IntOnly(1, U)
					if op == 0 != has[x] {
						break
					}
				}
				has[x] = op == 0
			} else {
				x = rg.IntOnly(1, U)
			}
			rg.Byte("+-?"[op])
			rg.Space()
			rg.Int(x, x)
			rg.NewLine()
		}
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var q, x int
		var op string
		Fscan(in, &q)
		has := map[int]bool{}
		for ; q > 0; q-- {
			Fscan(in, &op, &x)
			if op == "+" {
				has[x] = true
			} else if op == "-" {
				delete(has, x)
			} else {
				mex := x
				for has[mex] {
					mex += x
				}
				Fprintln(out, mex)
			}
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, cf1732D2)
}
