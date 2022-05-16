package main

import "fmt"

type prioque struct {
	heap  []*graphVertex
	count int
}

func (p *prioque) heapify(i, n int) {
	var l, j, r int
	for {
		l = 2*i + 1
		r = l + 1
		j = i
		if l < n && p.heap[i].key > p.heap[l].key {
			i = l
		}
		if r < n && p.heap[i].key > p.heap[r].key {
			i = r
		}
		if i == j {
			break
		}
		p.heap[i], p.heap[j] = p.heap[j], p.heap[i]
		p.heap[i].index = i
		p.heap[j].index = j
	}
}

func (p *prioque) insert(v *graphVertex) {
	i := p.count
	p.count = i + 1
	p.heap[i] = v
	for i > 0 && p.heap[(i-1)/2].key > p.heap[i].key {
		p.heap[(i-1)/2], p.heap[i] = p.heap[i], p.heap[(i-1)/2]
		p.heap[i].index = i
		i = (i - 1) / 2
	}
	p.heap[i].index = i
}

func (p *prioque) init(n int) {
	p.heap = make([]*graphVertex, n)
	p.count = 0
}

func (p *prioque) decreaseKey(vi, k int) {
	i := vi
	p.heap[vi].key = k
	for i > 0 && p.heap[(i-1)/2].key > k {
		p.heap[(i-1)/2], p.heap[i] = p.heap[i], p.heap[(i-1)/2]
		p.heap[i].index = i
		i = (i - 1) / 2
	}
	p.heap[i].index = i
}

func (p *prioque) extractMin() *graphVertex {
	min := p.heap[0]
	p.count--
	if p.count > 0 {
		p.heap[0] = p.heap[p.count]
		p.heap[0].index = 0
		p.heapify(0, p.count)
	}
	return min
}

type graphEdge struct {
	length   int
	toVertex int
}

type graphVertex struct {
	name, index, key, value int
	graphEdges              []graphEdge
}

func MST_Prim(listIncidence []*graphVertex) int {
	length := 0
	v := listIncidence[0]
	var p prioque
	p.init(len(listIncidence) - 1)
	for {
		v.index = -2
		for _, e := range v.graphEdges {
			if listIncidence[e.toVertex].index == -1 {
				listIncidence[e.toVertex].key = e.length
				listIncidence[e.toVertex].value = v.name
				p.insert(listIncidence[e.toVertex])
			} else {
				if listIncidence[e.toVertex].index != -2 && e.length <= listIncidence[e.toVertex].key {
					listIncidence[e.toVertex].value = v.name
					p.decreaseKey(listIncidence[e.toVertex].index, e.length)
				}
			}
		}
		if p.count == 0 {
			break
		}
		v = p.extractMin()
		length += v.key
	}
	return length
}

func main() {
	var n, m, v1, v2, length int
	fmt.Scan(&n)
	fmt.Scan(&m)
	listIncidence := make([]*graphVertex, 0)
	for i := 0; i < n; i++ {
		var s graphVertex
		s.name = i
		s.index = -1
		s.graphEdges = make([]graphEdge, 0)
		listIncidence = append(listIncidence, &s)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&v1, &v2, &length)
		var a graphEdge
		a.length = length
		a.toVertex = v2
		listIncidence[v1].graphEdges = append(listIncidence[v1].graphEdges, a)
		a.toVertex = v1
		listIncidence[v2].graphEdges = append(listIncidence[v2].graphEdges, a)
	}
	fmt.Println(MST_Prim(listIncidence))
}
