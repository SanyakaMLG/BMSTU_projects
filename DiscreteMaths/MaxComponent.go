package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Comp struct {
	number    int
	vertex    int
	edge      int
	minVertex int
}

type graphVertex struct {
	mark       int // 0  - white, 1 - black
	comp       int
	name       int
	graphEdges []int
}

func DFS(listIncidence *[]graphVertex, c *Comp) {
	var s Comp
	s.number = 0
	for i, v := range *listIncidence {
		if v.mark == 0 {
			VisitVertex1(listIncidence, &s, &(*listIncidence)[i])
			if s.vertex > c.vertex {
				*c = s
			}
			if s.vertex == c.vertex {
				if s.edge > c.edge {
					*c = s
				}
				if s.edge == c.edge {
					if s.minVertex > c.minVertex {
						*c = s
					}
				}
			}
			s.vertex = 0
			s.edge = 0
			s.minVertex = -1
			s.number++
		}
	}
}
func VisitVertex1(listIncidence *[]graphVertex, component *Comp, vertex *graphVertex) {
	(*vertex).mark = 1
	(*vertex).comp = (*component).number
	(*component).vertex++
	if (*vertex).name < (*component).minVertex {
		(*component).minVertex = (*vertex).name
	}
	for _, e := range (*vertex).graphEdges {
		(*component).edge++
		if (*listIncidence)[e].mark == 0 {
			VisitVertex1(listIncidence, component, &((*listIncidence)[e]))
		}
	}
}

func main() {
	bufstdin := bufio.NewReader(os.Stdin)
	color := "[color = red]"
	edges := make([][2]int, 0)
	var c Comp
	c.edge = 0
	c.vertex = 0
	c.number = -1
	c.minVertex = -1
	var ve [2]int
	var n, m, v1, v2 int
	fmt.Fscan(bufstdin, &n, &m)
	listIncidence := make([]graphVertex, 0)
	for i := 0; i < n; i++ {
		var s graphVertex
		s.mark = 0
		s.name = i
		s.comp = -1
		s.graphEdges = make([]int, 0)
		listIncidence = append(listIncidence, s)
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(bufstdin, &v1, &v2)
		ve[0] = v1
		ve[1] = v2
		edges = append(edges, ve)
		listIncidence[v1].graphEdges = append(listIncidence[v1].graphEdges, v2)
		listIncidence[v2].graphEdges = append(listIncidence[v2].graphEdges, v1)
	}
	DFS(&listIncidence, &c)
	fmt.Println("graph {")
	for _, v := range listIncidence {
		if v.comp == c.number {
			fmt.Println(strconv.Itoa(v.name) + " " + color)
		} else {
			fmt.Println(v.name)
		}
	}
	for _, ve := range edges {
		if listIncidence[ve[0]].comp == c.number {
			fmt.Println(strconv.Itoa(ve[0]) + " -- " + strconv.Itoa(ve[1]) + " " + color)
		} else {
			fmt.Println(strconv.Itoa(ve[0]) + " -- " + strconv.Itoa(ve[1]))
		}
	}
	fmt.Println("}")
}
