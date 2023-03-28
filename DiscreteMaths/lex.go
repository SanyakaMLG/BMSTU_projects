package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// avl tree and skiplist from ADS 3rd module

type SkipList struct {
	key   string
	value int
	next  []*SkipList
}

func InitSkipList(m int) *SkipList {
	l := &SkipList{}
	l.next = make([]*SkipList, m)
	for i := 0; i < m; i++ {
		l.next[i] = nil
	}
	return l
}

func (x *SkipList) Succ() *SkipList {
	y := x.next[0]
	return y
}

func Skip(l *SkipList, m int, key string, p []*SkipList) []*SkipList {
	x := l
	for i := m - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < key {
			x = x.next[i]
		}
		p[i] = x
	}
	return p
}

func (l *SkipList) Search(m int, key string) (int, bool) {
	p := make([]*SkipList, m)
	p = Skip(l, m, key, p)
	x := p[0].Succ()
	if x == nil || x.key != key {
		return 0, false
	}
	return x.value, true
}

func (l *SkipList) Insert(m int, key string, value int) {
	p := make([]*SkipList, m)
	p = Skip(l, m, key, p)
	x := InitSkipList(m)
	x.key = key
	x.value = value
	x.next = make([]*SkipList, m)
	r := rand.Int() * 2
	i := 0
	for i < m && r%2 == 0 {
		x.next[i] = p[i].next[i]
		p[i].next[i] = x
		i++
		r = r / 2
	}
	for i < m {
		x.next[i] = nil
		i++
	}
}

type Node struct {
	key     string
	value   int
	balance int
	parent  *Node
	left    *Node
	right   *Node
}

type AVLTree struct {
	root *Node
}

func (n *Node) MapEmpty() bool {
	if n == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Minimum() *Node {
	if n == nil {
		return n
	} else {
		x := n
		for x.left != nil {
			x = x.left
		}
		return x
	}
}

func (n *Node) Succ() *Node {
	if n.right != nil {
		return n.right.Minimum()
	} else {
		x := n.parent
		for x != nil && n == x.right {
			n = x
			x = x.parent
		}
		return x
	}
}

func (n *Node) search(key string) (int, bool) {
	x := n
	for x != nil && x.key != key {
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	if x == nil {
		return 0, false
	} else {
		return x.value, true
	}
}

func (n *Node) Insert(key string, value int) *Node {
	y := Node{key, value, 0, nil, nil, nil}
	if n == nil {
		*n = y
	} else {
		x := n
		for {
			if key < x.key {
				if x.left == nil {
					x.left = &y
					y.parent = x
					break
				}
				x = x.left
			} else {
				if x.right == nil {
					x.right = &y
					y.parent = x
					break
				}
				x = x.right
			}
		}
	}
	return n
}

func (n *Node) ReplaceNode(x, y *Node) {
	if x == n {
		*n = *y
		if y != nil {
			y.parent = nil
		}
	} else {
		p := x.parent
		if y != nil {
			y.parent = p
		}
		if p.left == x {
			p.left = y
		} else {
			p.right = y
		}
	}
}

func (n *Node) RotateLeft(x *Node) {
	y := x.right
	n.ReplaceNode(x, y)
	b := y.left
	if b != nil {
		b.parent = x
	}
	x.right = b
	x.parent = y
	y.left = x
	x.balance--
	if y.balance > 0 {
		x.balance -= y.balance
	}
	y.balance--
	if x.balance < 0 {
		y.balance += x.balance
	}
}

func (n *Node) RotateRight(x *Node) {
	y := x.left
	n.ReplaceNode(x, y)
	b := y.right
	if b != nil {
		b.parent = x
	}
	x.left = b
	x.parent = y
	y.right = x
	x.balance++
	if y.balance < 0 {
		x.balance -= y.balance
	}
	y.balance++
	if x.balance > 0 {
		y.balance += x.balance
	}
}

func (n *Node) InsertAvl(key string, value int) *Node {
	a := n.Insert(key, value)
	for {
		x := a.parent
		if x == nil {
			break
		}
		if a == x.left {
			x.balance--
			if x.balance == 0 {
				break
			}
			if x.balance == -2 {
				if a.balance == 1 {
					n.RotateLeft(a)
				}
				n.RotateRight(x)
				break
			}
		} else {
			x.balance++
			if x.balance == 0 {
				break
			}
			if x.balance == 2 {
				if a.balance == -1 {
					n.RotateRight(a)
				}
				n.RotateLeft(x)
				break
			}
		}
		a = x
	}
	return a
}

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

func (n *Node) Assign(s string, x int) {
	*n = *n.InsertAvl(s, x)
}

func (n *Node) Lookup(s string) (int, bool) {
	return n.search(s)
}

func (l *SkipList) Assign(s string, x int) {
	l.Insert(len(l.next), s, x)
}

func (l *SkipList) Lookup(s string) (int, bool) {
	return l.Search(len(l.next), s)
}

func lex(sentence string, array AssocArray) []int {
	result := make([]int, 0)
	words := make([]string, 0)
	var word = ""
	for i := 0; i < len(sentence); i++ {
		if sentence[i] != ' ' {
			for i < len(sentence) && sentence[i] != ' ' {
				word += string(sentence[i])
				i++
			}
			words = append(words, word)
			word = ""
		}
	}

	k := 1

	for i := range words {
		idx, check := array.Lookup(words[i])
		if check {
			result = append(result, idx)
		} else {
			array.Assign(words[i], k)
			result = append(result, k)
			k++
		}
	}

	return result
}

func main() {
	var str string
	root := &Node{}
	tree := AssocArray(root)
	str, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	list := InitSkipList(len(str))
	skiplist := AssocArray(list)
	ans1 := lex(str, tree)
	ans2 := lex(str, skiplist)
	fmt.Println(ans1)
	fmt.Println(ans2)
}
