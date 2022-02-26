package main

// Problem description URL https://atcoder.jp/contests/abc236/tasks/abc240_a

import (
	"bufio"
	"container/list"
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
	n   = 0
	a   []int
	res []int
	s   *list.List
)

func solve() {
	scan(&n)
	a = make([]int, n+1)
	res = make([]int, 0)
	for i := 0; i < n; i++ {
		scan(&a[i])
	}
	s = list.New()
	for i := 0; i < n; i++ {
		if s.Len() != 0 {
			// 判断与下一个是否相同
			if i+1 < n && a[i] == a[i+1] {
				s.PushBack(a[i])
			} else {
				// 判断与之前一个是否相同
				b := s.Back()
				if b.Value != a[i] {
					s.PushBack(a[i])
				}
				for b.Value == a[i] {
					s.Remove(b)
					if s.Len() != 0 {
						b = s.Back()
					} else {
						break
					}
				}

			}
		} else {
			s.PushBack(a[i])
		}
		res = append(res, s.Len())
	}
	for _, v := range res {
		printf("%d\n", v)
	}

}

func main() {
	defer out.Flush()
	solve()
}
