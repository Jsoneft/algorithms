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
	A int
	B int
}

type Node struct {
	CurPos int
	I      int
}

var (
	steps []pair
	n     = 0
	x     = 0
	res   = false
)

func solve() {
	steps = make([]pair, 0)
	scan(&n, &x)
	tmp := pair{
		A: 0,
		B: 0,
	}
	for i := 0; i < n; i++ {
		scan(&tmp.A, &tmp.B)
		steps = append(steps, tmp)
	}

	res = bfs(0)
	if res == false {
		printf("No")
	} else {
		printf("Yes")
	}
}

func bfs(s int) bool {
	q := list.New()
	q.PushBack(Node{
		CurPos: s,
		I:      0,
	})
	for q.Len() != 0 {
		cur := q.Remove(q.Front())
		node := cur.(Node)
		if node.CurPos == x && node.I == n {
			return true
		}
		if node.I != n {
			q.PushBack(Node{
				CurPos: node.CurPos + steps[node.I].A,
				I:      node.I + 1,
			})
			q.PushBack(Node{
				CurPos: node.CurPos + steps[node.I].B,
				I:      node.I + 1,
			})
		}
	}
	return false
}

func dfs(curPos int, idx int) bool {
	if res == true {
		return true
	}
	if idx == n {
		if curPos == x {
			res = true
			return res
		}
		return false
	}
	return dfs(curPos+steps[idx].A, idx+1) || dfs(curPos+steps[idx].B, idx+1)
}

func main() {
	defer out.Flush()
	solve()
}
