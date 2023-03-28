package main

import (
	"fmt"
	"math"
	"sort"
)

type Attraction struct {
	x, y       int
	ancestor   *Attraction
	dependence int
}

type Path struct {
	length           int
	parent1, parent2 *Attraction
}

func find(v *Attraction) *Attraction {
	a := v
	for a.ancestor != nil {
		a = a.ancestor
	}
	return a
}

func union(v, u *Attraction) {
	ancestor1 := find(v)
	ancestor2 := find(u)
	if ancestor1.dependence < ancestor2.dependence {
		ancestor1.ancestor = ancestor2
	} else {
		if ancestor1.dependence == ancestor2.dependence {
			ancestor1.dependence++
		}
		ancestor2.ancestor = ancestor1
	}

}

func spanningTree(roads []Path, m int) float64 {
	var short float64
	short = 0
	for i := 0; i < m; i++ {
		parent1 := roads[i].parent1
		parent2 := roads[i].parent2
		if find(parent1) != find(parent2) {
			short += math.Sqrt(float64(roads[i].length))
			union(parent1, parent2)
		}
	}
	return short
}

func MST_Kruskal(roads []Path, m int) float64 {
	sort.Slice(roads, func(i, j int) bool {
		return roads[i].length < roads[j].length
	})
	short := spanningTree(roads, m)
	return short
}

func main() {
	var n, x, y int
	r := 0
	var shortPath float64
	var a Attraction
	var s Path
	fmt.Scan(&n)
	lenRoad := n * (n - 1) / 2
	park := make([]Attraction, n)
	roads := make([]Path, lenRoad)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		a.x = x
		a.y = y
		a.ancestor = nil
		a.dependence = 1
		park[i] = a
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			s.parent1 = &park[i]
			s.parent2 = &park[j]
			s.length = (park[i].x-park[j].x)*(park[i].x-park[j].x) + (park[i].y-park[j].y)*(park[i].y-park[j].y)
			roads[r] = s
			r++
		}
	}
	shortPath = MST_Kruskal(roads, lenRoad)
	fmt.Printf("%.2f\n", shortPath)
}
