package main

import "fmt"

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
	mark       int // 0  - white, 1 - black
	parent     *graphVertex
	comp       int
	name       int
	graphEdges []int
}

func DFS1(listIncidence *[]graphVertex, q *queue, bridges *int) {
	for i, v := range *listIncidence {
		if v.mark == 0 {
			*bridges--
			VisitVertex1(listIncidence, q, &(*listIncidence)[i])
		}
	}
}
func VisitVertex1(listIncidence *[]graphVertex, q *queue, vertex *graphVertex) {
	(*vertex).mark = 1
	(*q).enqueue(vertex)
	for _, e := range (*vertex).graphEdges {
		if (*listIncidence)[e].mark == 0 {
			(*listIncidence)[e].parent = vertex
			VisitVertex1(listIncidence, q, &((*listIncidence)[e]))
		}
	}
}
func DFS2(listIncidence *[]graphVertex, q *queue, bridges *int) {
	var s *graphVertex
	component := 0
	for (*q).count > 0 {
		s = (*q).dequeue()
		if (*s).comp == -1 {
			VisitVertex2(listIncidence, s, component)
			*bridges++
			component++
		}
	}

}
func VisitVertex2(listIncidence *[]graphVertex, vertex *graphVertex, component int) {
	(*vertex).comp = component
	for _, e := range (*vertex).graphEdges {
		if ((*listIncidence)[e].comp == -1) && ((*listIncidence)[e].parent != vertex) {
			VisitVertex2(listIncidence, &((*listIncidence)[e]), component)
		}
	}
}

func main() {
	var n, m, v1, v2 int
	fmt.Scan(&n)
	fmt.Scan(&m)
	listIncidence := make([]graphVertex, 0)
	var q queue
	q.init(n)
	for i := 0; i < n; i++ {
		var s graphVertex
		s.mark = 0
		s.name = i
		s.comp = -1
		s.graphEdges = make([]int, 0)
		listIncidence = append(listIncidence, s)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&v1, &v2)
		listIncidence[v1].graphEdges = append(listIncidence[v1].graphEdges, v2)
		listIncidence[v2].graphEdges = append(listIncidence[v2].graphEdges, v1)
	}
	bridges := 0
	DFS1(&listIncidence, &q, &bridges)
	DFS2(&listIncidence, &q, &bridges)
	fmt.Println(bridges)
}
