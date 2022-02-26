package abc240

// Problem description URL https://atcoder.jp/contests/abc236/tasks/abc240_a

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func scan(a ...interface{})             { fmt.Fscan(in, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(out, f, a...) }
func solve() {
	a, b := 0, 0
	scan(&a, &b)
	abs := math.Abs(float64(a - b))
	if abs == 9 || abs == 1 {
		printf("Yes")
	} else {
		printf("No")
	}
}

func main() {
	defer out.Flush()
	solve()
}
