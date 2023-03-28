package main

import (
	"fmt"
	"sort"
)

var time = 1
var count = 0

type stack struct {
	data []int
	top  int
}

func (s *stack) push(vertex int) {
	s.data[s.top] = vertex
	s.top++
}

func (s *stack) pop() int {
	s.top--
	return s.data[s.top]
}

type graphVertex struct {
	entry      bool
	t1         int
	low        int
	comp       int
	name       int
	graphEdges []int
}

func Tarjan(listIncidence *[]graphVertex) {
	var s stack
	s.top = 0
	s.data = make([]int, len(*listIncidence))
	for e, _ := range *listIncidence {
		if (*listIncidence)[e].t1 == 0 {
			visitVertexTarjan(listIncidence, e, &s)
		}
	}
}

func indexOf(slice []int, element int) int {
	for i, x := range slice {
		if x == element {
			return i
		}
	}
	return -1
}
func visitVertexTarjan(listIncidence *[]graphVertex, numV int, s *stack) {
	(*listIncidence)[numV].t1 = time
	(*listIncidence)[numV].low = time
	time++
	s.push(numV)
	for _, e := range (*listIncidence)[numV].graphEdges {
		if (*listIncidence)[e].t1 == 0 {
			visitVertexTarjan(listIncidence, e, s)
		}
		if ((*listIncidence)[e].comp == -1) && ((*listIncidence)[numV].low > (*listIncidence)[e].low) {
			(*listIncidence)[numV].low = (*listIncidence)[e].low
		}
	}
	if (*listIncidence)[numV].low == (*listIncidence)[numV].t1 {
		u := s.pop()
		(*listIncidence)[u].comp = count
		for u != numV {
			u = s.pop()
			(*listIncidence)[u].comp = count
		}
		count++
	}
}
func entryVertex(listIncidence *[]graphVertex, numV int) {
	if !(*listIncidence)[numV].entry {
		(*listIncidence)[numV].entry = true
		for _, e := range (*listIncidence)[numV].graphEdges {
			entryVertex(listIncidence, e)
		}
	}
}
func main() {
	var n, m, v1, v2 int
	fmt.Scan(&n)
	fmt.Scan(&m)
	listIncidence := make([]graphVertex, 0)
	base := make([]int, 0)
	for i := 0; i < n; i++ {
		var s graphVertex
		s.name = i
		s.comp = -1
		s.t1 = 0
		s.graphEdges = make([]int, 0)
		listIncidence = append(listIncidence, s)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&v1, &v2)
		listIncidence[v1].graphEdges = append(listIncidence[v1].graphEdges, v2)
	}
	Tarjan(&listIncidence)
	condencation := make([]graphVertex, 0)
	for i := 0; i < count; i++ {
		var s graphVertex
		s.name = i
		s.entry = false
		s.graphEdges = make([]int, 0)
		condencation = append(condencation, s)
	}
	for v := range listIncidence {
		for _, e := range listIncidence[v].graphEdges {
			if listIncidence[v].comp != listIncidence[e].comp {
				if indexOf(condencation[listIncidence[v].comp].graphEdges, listIncidence[e].comp) == -1 {
					condencation[listIncidence[v].comp].graphEdges = append(condencation[listIncidence[v].comp].graphEdges, listIncidence[e].comp)
				}
			}
		}
	}
	for v := range condencation {
		if !condencation[v].entry {
			for _, e := range condencation[v].graphEdges {
				entryVertex(&condencation, e)
			}
		}
	}
	for v := range condencation {
		if !condencation[v].entry {
			minV := n
			for e := range listIncidence {
				if listIncidence[e].comp == v && e < minV {
					minV = e
				}
			}
			if minV != n {
				base = append(base, minV)
			}

		}
	}
	sort.Ints(base)
	for _, x := range base {
		fmt.Printf("%d ", x)
	}
}
