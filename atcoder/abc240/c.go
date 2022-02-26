package abc240

// Problem description URL https://atcoder.jp/contests/abc236/tasks/abc240_a

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func scan(a ...interface{})             { fmt.Fscan(in, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(out, f, a...) }

type pair struct {
	A int
	B int
}

var (
	n     = 0
	x     = 0
	steps []pair
	dp    [][]bool
)

func solve() {
	scan(&n, &x)
	tmp := pair{
		A: 0,
		B: 0,
	}
	for i := 0; i < n; i++ {
		scan(&tmp.A, &tmp.B)
		steps = append(steps, tmp)
	}

	dp = make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, 10001)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < 10001; j++ {
			if j+steps[i].A < 10001 && i+1 < n+1 {
				dp[i+1][j+steps[i].A] = dp[i][j] || dp[i+1][j+steps[i].A]
			}
			if j+steps[i].B < 10001 && i+1 < n+1 {
				dp[i+1][j+steps[i].B] = dp[i][j] || dp[i+1][j+steps[i].B]
			}
		}
	}
	if dp[n][x] == true {
		printf("Yes")
	} else {
		printf("No")
	}
}

func main() {
	defer out.Flush()
	solve()
}
