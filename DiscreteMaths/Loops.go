package main

import (
	"fmt"
	"sort"
)

var count = 1

type stack struct {
	data []*graphVertex
	top  int
}

func (s *stack) push(vertex *graphVertex) {
	s.data[s.top] = vertex
	s.top++
}
func (s *stack) pop() *graphVertex {
	s.top--
	return s.data[s.top]
}

type graphVertex struct {
	name           int
	dead           bool
	time           int
	command        string
	operand        int
	parent         *graphVertex
	ancestor       *graphVertex
	idom           *graphVertex
	sdom           *graphVertex
	label          *graphVertex
	graphEdgesTo   []*graphVertex
	graphEdgesFrom []*graphVertex
	bucket         []*graphVertex
}

func FindMin(v *graphVertex, s *stack) *graphVertex {
	var min *graphVertex
	if v.ancestor == nil {
		min = v
	} else {
		s.top = 0
		var u = v
		for u.ancestor.ancestor != nil {
			s.push(u)
			u = u.ancestor
		}
		for s.top != 0 {
			v = s.pop()
			if v.ancestor.label.sdom.time < v.label.sdom.time {
				v.label = v.ancestor.label
			}
			v.ancestor = u.ancestor
		}
		min = v.label
	}
	return min
}

func Dominators(listIncidence []*graphVertex, n int) []*graphVertex {
	var s stack
	s.data = make([]*graphVertex, n)
	for _, w := range listIncidence {
		if w.time == 1 {
			continue
		}
		for _, v := range w.graphEdgesFrom {
			u := FindMin(v, &s)
			if u.sdom.time < w.sdom.time {
				w.sdom = u.sdom
			}
		}
		w.ancestor = w.parent
		w.sdom.bucket = append(w.sdom.bucket, w)
		for _, v := range w.parent.bucket {
			u := FindMin(v, &s)
			if u.sdom == v.sdom {
				v.idom = v.sdom
			} else {
				v.idom = u
			}
		}
		w.parent.bucket = nil
	}
	for _, w := range listIncidence {
		if w.time == 1 {
			continue
		}
		if w.idom != w.sdom {
			w.idom = w.idom.idom
		}
	}
	listIncidence[len(listIncidence)-1].idom = nil
	return listIncidence
}

func DFS(r *graphVertex) {
	r.dead = false
	r.time = count
	count++
	for e := range r.graphEdgesTo {
		if r.graphEdgesTo[e].dead {
			r.graphEdgesTo[e].parent = r
			DFS(r.graphEdgesTo[e])
		}
	}
}

func main() {
	var n, v1, v2 int
	var s string
	fmt.Scan(&n)
	listIncidence := make([]*graphVertex, 0)
	assoc := make(map[int]int) // Метка -> порядковый номер в графе
	for i := 0; i < n; i++ {
		var s graphVertex
		s.name = i
		s.dead = true
		s.graphEdgesTo = make([]*graphVertex, 0)
		s.graphEdgesFrom = make([]*graphVertex, 0)
		s.bucket = make([]*graphVertex, 0)
		s.ancestor = nil
		s.sdom = &s
		s.label = &s
		listIncidence = append(listIncidence, &s)
	}
	for i := range listIncidence {
		fmt.Scan(&v1, &s)
		listIncidence[i].command = s
		if s != "ACTION" {
			fmt.Scan(&v2)
			listIncidence[i].operand = v2
		}
		assoc[v1] = i
	}
	for i := range listIncidence {
		switch listIncidence[i].command {
		case "ACTION":
			if i != n-1 {
				listIncidence[i].graphEdgesTo = append(listIncidence[i].graphEdgesTo, listIncidence[i+1])
				listIncidence[i+1].graphEdgesFrom = append(listIncidence[i+1].graphEdgesFrom, listIncidence[i])
			}
			break
		case "JUMP":
			v2 = assoc[listIncidence[i].operand]
			listIncidence[i].graphEdgesTo = append(listIncidence[i].graphEdgesTo, listIncidence[v2])
			listIncidence[v2].graphEdgesFrom = append(listIncidence[v2].graphEdgesFrom, listIncidence[i])
			break
		case "BRANCH":
			v2 = assoc[listIncidence[i].operand]
			listIncidence[i].graphEdgesTo = append(listIncidence[i].graphEdgesTo, listIncidence[v2])
			listIncidence[v2].graphEdgesFrom = append(listIncidence[v2].graphEdgesFrom, listIncidence[i])
			if i != n-1 {
				listIncidence[i].graphEdgesTo = append(listIncidence[i].graphEdgesTo, listIncidence[i+1])
				listIncidence[i+1].graphEdgesFrom = append(listIncidence[i+1].graphEdgesFrom, listIncidence[i])
			}
			break
		}
	}
	DFS(listIncidence[0])
	for i := 0; i < len(listIncidence); i++ {
		if listIncidence[i].dead {
			listIncidence[i] = listIncidence[len(listIncidence)-1]
			listIncidence[len(listIncidence)-1] = nil
			listIncidence = listIncidence[:len(listIncidence)-1]
			i--
		} else {
			for j := 0; j < len(listIncidence[i].graphEdgesFrom); j++ {
				if listIncidence[i].graphEdgesFrom[j].dead {
					listIncidence[i].graphEdgesFrom[j] = listIncidence[i].graphEdgesFrom[len(listIncidence[i].graphEdgesFrom)-1]
					listIncidence[i].graphEdgesFrom = listIncidence[i].graphEdgesFrom[:len(listIncidence[i].graphEdgesFrom)-1]
					j--
				}
			}
		}
	}
	sort.Slice(listIncidence, func(i, j int) bool {
		return listIncidence[i].time > listIncidence[j].time
	})
	n = len(listIncidence)
	listIncidence = Dominators(listIncidence, n)
	loops := 0
	for _, v := range listIncidence {
		for _, e := range v.graphEdgesFrom {
			for e != v && e != nil {
				e = e.idom
			}
			if e == v {
				loops++
				break
			}
		}
	}
	fmt.Println(loops)
}
