package main

import (
	"fmt"
	"os"
	"sort"
)

type graphVertex struct {
	name       int
	graphEdges []int
	com        int
}

func DFS(listIncidence *[]graphVertex) int {
	island := 0
	for i, _ := range *listIncidence {
		if (*listIncidence)[i].com == -1 && len((*listIncidence)[i].graphEdges) != 0 {
			visitVertex(listIncidence, i, island)
			island += 2
		}
	}
	return island
}
func visitVertex(listIncidence *[]graphVertex, v int, island int) {
	(*listIncidence)[v].com = island
	for _, e := range (*listIncidence)[v].graphEdges {
		if (*listIncidence)[v].com == (*listIncidence)[e].com {
			fmt.Println("No solution")
			os.Exit(0)
		}
		if (*listIncidence)[e].com == -1 {
			if island%2 == 0 {
				visitVertex(listIncidence, e, island+1)
			} else {
				visitVertex(listIncidence, e, island-1)
			}
		}
	}
}

func compareGroups(g1, g2 []int) []int {
	for i := range g1 {
		if g1[i] > g2[i] {
			return g2
		}
		if g2[i] > g1[i] {
			return g1
		}
	}
	return g1
}

func rekur(resultGroup *[]int, startGroup []int, GoodGays []int, soch []int, teamsList [][][]int, index int, n int) {
	if index == len(teamsList) {
		tryGroup := startGroup
		for i := range soch {
			if soch[i] == 0 {
				for _, v := range teamsList[i][0] {
					tryGroup = append(tryGroup, v)
				}
			} else {
				for _, v := range teamsList[i][1] {
					tryGroup = append(tryGroup, v)
				}
			}
		}
		c := 0
		for len(tryGroup) < n/2 {
			tryGroup = append(tryGroup, GoodGays[c])
			c++
		}
		sort.Ints(tryGroup)
		if len(tryGroup) <= n/2 && len(tryGroup) >= len(*resultGroup) {
			if len(tryGroup) > len(*resultGroup) {
				*resultGroup = tryGroup
			} else {
				*resultGroup = compareGroups(tryGroup, *resultGroup)
			}
		}
	} else {
		soch[index] = 0
		rekur(resultGroup, startGroup, GoodGays, soch, teamsList, index+1, n)
		soch[index] = 1
		rekur(resultGroup, startGroup, GoodGays, soch, teamsList, index+1, n)
	}
}

func main() {
	listIncidence := make([]graphVertex, 0)
	var n int
	var sym string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s graphVertex
		s.com = -1
		s.name = i
		s.graphEdges = make([]int, 0)
		listIncidence = append(listIncidence, s)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > j {
				fmt.Scan(&sym)
				continue
			} else {
				fmt.Scan(&sym)
				if sym == "+" {
					listIncidence[i].graphEdges = append(listIncidence[i].graphEdges, j)
					listIncidence[j].graphEdges = append(listIncidence[j].graphEdges, i)
				}
			}
		}
	}
	island := DFS(&listIncidence)
	var startTeam = make([]int, 0)
	var teamsList = make([][][]int, 0)
	for i := 0; i <= island; i += 2 {
		t1 := make([]int, 0)
		t2 := make([]int, 0)
		for _, v := range listIncidence {
			if v.com == i {
				t1 = append(t1, v.name)
			}
			if v.com == i+1 {
				t2 = append(t2, v.name)
			}
		}
		if len(t1) == len(t2) {
			t := compareGroups(t1, t2)
			for _, e := range t {
				startTeam = append(startTeam, e)
			}
		} else {
			teamPair := make([][]int, 2)
			teamPair[0] = t1
			teamPair[1] = t2
			teamsList = append(teamsList, teamPair)
		}
	}
	sort.Ints(startTeam)
	GoodGays := make([]int, 0)
	for _, v := range listIncidence {
		if v.com == -1 {
			GoodGays = append(GoodGays, v.name)
		}
	}
	resultTeam := startTeam
	soch := make([]int, len(teamsList))
	rekur(&resultTeam, startTeam, GoodGays, soch, teamsList, 0, n)
	for _, v := range resultTeam {
		fmt.Printf("%d ", v+1)
	}
}
