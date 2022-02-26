package abc236

import (
	"bufio"
	"fmt"
	"os"
)

// Problem description URL https://atcoder.jp/contests/abc236/tasks/abc236_b

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func scan(a ...interface{}) {
	fmt.Fscan(in, a...)
}
func printf(f string, a ...interface{}) {
	fmt.Fprintf(out, f, a...)
}
func solve() {
	n := 0
	scan(&n)
	res := 0
	val := 0
	for i := 0; i < 4*n-1; i++ {
		scan(&val)
		res ^= val
	}
	printf("%d", res)

}
func main() {
	defer out.Flush()
	solve()
}
