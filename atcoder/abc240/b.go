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
func solve() {
	n := 0
	scan(&n)
	set := make(map[int]bool, 0)
	for i := 0; i < n; i++ {
		val := 0
		scan(&val)
		set[val] = true
	}
	printf("%d", len(set))
}

func main() {
	defer out.Flush()
	solve()
}
