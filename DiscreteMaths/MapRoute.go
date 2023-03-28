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
		if l < n && p.heap[i].dist > p.heap[l].dist {
			i = l
		}
		if r < n && p.heap[i].dist > p.heap[r].dist {
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
	for i > 0 && p.heap[(i-1)/2].dist > p.heap[i].dist {
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

func (p *prioque) decreaseKey(vi, d int) {
	i := vi
	p.heap[vi].dist = d
	for i > 0 && p.heap[(i-1)/2].dist > d {
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

type Name struct {
	x, y int
}

type graphVertex struct {
	name       Name
	graphEdges []Name
	dist       int
	index      int
	w          int
}

func relax(u, v *graphVertex, w int) bool {
	changed := u.dist+w < v.dist
	if changed {
		v.dist = u.dist + w
	}
	return changed
}

func Dijkstra(listIncidence *[][]graphVertex, n int) {
	var q prioque
	q.init(n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				(*listIncidence)[i][j].dist = (*listIncidence)[i][j].w
			} else {
				(*listIncidence)[i][j].dist = 1000000000 // аналог бесконечности :)
			}
			q.insert(&(*listIncidence)[i][j])
		}
	}
	for q.count > 0 {
		v := q.extractMin()
		v.index = -1
		for _, e := range v.graphEdges {
			if (*listIncidence)[e.x][e.y].index != -1 && relax(v, &(*listIncidence)[e.x][e.y], (*listIncidence)[e.x][e.y].w) {
				q.decreaseKey((*listIncidence)[e.x][e.y].index, (*listIncidence)[e.x][e.y].dist)
			}
		}
	}
}

func main() {
	var n, d int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Scan(&d)
		fmt.Println(d)
	} else {
		listIncidence := make([][]graphVertex, n)
		for i := 0; i < n; i++ {
			listIncidence[i] = make([]graphVertex, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				var s graphVertex
				s.graphEdges = make([]Name, 0)
				listIncidence[i] = append(listIncidence[i], s)
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				var name Name
				name.x = i
				name.y = j
				fmt.Scan(&d)
				listIncidence[i][j].w = d
				listIncidence[i][j].name = name
				if i == 0 && j == 0 {
					name.x = 0
					name.y = 1
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
					name.x = 1
					name.y = 0
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
				}
				if i == 0 && j == (n-1) {
					name.x = 1
					name.y = n - 1
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
				}
				if i == n-1 && j == 0 {
					name.x = n - 1
					name.y = 1
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
				}
				if (i == 0 || i == n-1) && j != 0 && j != n-1 {
					name.x = i
					name.y = j + 1
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
					name.y = j
					if i == 0 {
						name.x = 1
					} else {
						name.x = n - 2
					}
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
				}
				if (j == 0 || j == n-1) && i != 0 && i != n-1 {
					name.x = i + 1
					name.y = j
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
					name.x = i
					if j == 0 {
						name.y = 1
					} else {
						name.y = n - 2
					}
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
				}
				if i != 0 && i != n-1 && j != 0 && j != n-1 {
					name.x = i
					name.y = j + 1
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
					name.x = i
					name.y = j - 1
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
					name.x = i + 1
					name.y = j
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
					name.x = i - 1
					name.y = j
					listIncidence[i][j].graphEdges = append(listIncidence[i][j].graphEdges, name)
				}
			}
		}
		Dijkstra(&listIncidence, n)
		fmt.Println(listIncidence[n-1][n-1].dist)
	}
}
