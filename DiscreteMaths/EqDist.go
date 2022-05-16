package main

import (
	"fmt"
	"sort"
)

type queue struct {
	data  []*graphVertex
	cap   int
	count int
	head  int
	tail  int
}

func (q *queue) init(n int) {
	(*q).data = make([]*graphVertex, n)
	(*q).cap = n
	(*q).count = 0
	(*q).head = 0
	(*q).tail = 0
}

func (q *queue) enqueue(g *graphVertex) {
	(*q).data[(*q).tail] = g
	(*q).tail++
	if (*q).tail == (*q).cap {
		(*q).tail = 0
	}
	(*q).count++
}

func (q *queue) dequeue() *graphVertex {
	g := (*q).data[(*q).head]
	(*q).head++
	if (*q).head == (*q).cap {
		(*q).head = 0
	}
	(*q).count--
	return g
}

type graphVertex struct {
	mark       bool
	name       int
	graphEdges []int
}

func BFS(listIncidence []graphVertex, q queue, start int, distList *[]int) {
	for i := range listIncidence {
		listIncidence[i].mark = false
	}
	listIncidence[start].mark = true
	q.enqueue(&listIncidence[start])
	for q.count > 0 {
		v := q.dequeue()
		for _, e := range (*v).graphEdges {
			if !listIncidence[e].mark {
				(*distList)[e] = (*distList)[v.name] + 1
				listIncidence[e].mark = true
				q.enqueue(&listIncidence[e])
			}
		}
	}
}

func main() {
	var n, m, k, v1, v2 int
	fmt.Scan(&n)
	fmt.Scan(&m)
	listIncidence := make([]graphVertex, 0)
	var q queue
	q.init(n)
	for i := 0; i < n; i++ {
		var s graphVertex
		s.name = i
		s.mark = false
		s.graphEdges = make([]int, 0)
		listIncidence = append(listIncidence, s)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&v1, &v2)
		listIncidence[v1].graphEdges = append(listIncidence[v1].graphEdges, v2)
		listIncidence[v2].graphEdges = append(listIncidence[v2].graphEdges, v1)
	}
	fmt.Scan(&k)
	distances := make([][]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&v1)
		distances[i] = make([]int, n)
		BFS(listIncidence, q, v1, &distances[i])
	}
	Eq := make([]int, 0)
	Yep := true
	for i := 0; i < n; i++ {
		Yep = true
		for j := 0; j < k-1; j++ {
			if distances[j][i] != distances[j+1][i] || distances[j][i] == 0 {
				Yep = false
				break
			}
		}
		if Yep {
			Eq = append(Eq, i)
		}
	}
	sort.Ints(Eq)
	if len(Eq) == 0 {
		fmt.Println("-")
	} else {
		for _, e := range Eq {
			fmt.Printf("%d ", e)
		}
	}
}
