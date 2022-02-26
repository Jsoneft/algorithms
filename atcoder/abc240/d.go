package abc240

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
	Color int
	Size  int
}

var (
	n   = 0
	a   []int
	res []int
	s   *list.List
	cur = 0
)

func solve() {
	scan(&n)
	a = make([]int, n+1)
	res = make([]int, 0)
	for i := 0; i < n; i++ {
		scan(&a[i])
	}
	s = list.New()
	cur = 0
	for i := 0; i < n; i++ {
		if s.Len() != 0 {
			lstE := s.Back()
			lst := lstE.Value.(pair)
			// a,b -> a,b+1
			if lst.Color == a[i] {
				lst.Size++
				cur++
				lstE.Value = lst
				if lst.Size >= lst.Color {
					cur -= lst.Size
					s.Remove(lstE)
				}
			} else {
				s.PushBack(pair{
					Color: a[i],
					Size:  1,
				})
				cur++
			}

		} else {
			s.PushBack(pair{
				Color: a[i],
				Size:  1,
			})
			cur++
		}
		res = append(res, cur)
	}
	for _, v := range res {
		printf("%d\n", v)
	}

}

func main() {
	defer out.Flush()
	solve()
}
