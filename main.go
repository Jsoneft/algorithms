package main

// https://atcoder.jp/contests/abc236/tasks/abc236_a

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
	str := ""
	scan(&str)

	l, r := 0, 0
	scan(&l, &r)
	sBytes := []byte(str)
	sBytes[l-1], sBytes[r-1] = sBytes[r-1], sBytes[l-1]
	//printf("str: %s, l = %c, r = %c",sBytes, sBytes[l-1], sBytes[r-1])
	printf(string(sBytes))

}
func main() {
	defer out.Flush()
	solve()
}
