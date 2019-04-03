package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting...")
	p1()
	p2()
}

func p1() {

	t := time.Now()

	var players, last int

	scanned, _ := fmt.Sscanf(INPUT, "%d players; last marble is worth %d points", &players, &last)
	if scanned != 2 {
		panic("Couldn't scan input!")
	}

	board := []int{0}
	score := make([]int, players)
	current := 0

	for i := 1; i <= last; i++ {

		// fmt.Println(i-1, board)

		if i%23 == 0 {
			var point int
			current = (10*len(board) + current - 7) % len(board)
			board, point = remove(current, board)
			score[i%players] += i + point
		} else {
			current = (current + 2) % len(board)
			board = insert(i, current, board)
		}
	}

	// fmt.Println(board)
	result := highscore(score)

	fmt.Println("Done:", result, time.Since(t))
}

func p2() {

	t := time.Now()

	var players, last int

	scanned, _ := fmt.Sscanf(INPUT, "%d players; last marble is worth %d points", &players, &last)
	if scanned != 2 {
		panic("Couldn't scan input!")
	}

	last *= 100

	current := &Node{0, nil, nil}
	current.next, current.prev = current, current

	score := make([]int, players)

	for i := 1; i <= last; i++ {

		if i%23 == 0 {
			var point int

			// for i := 0; i < 7; i++ {
			// 	current = current.prev
			// }

			current = current.prev.prev.prev.prev.prev.prev.prev

			point = current.value
			current = removal(current)
			score[i%players] += i + point
		} else {
			current = insertion(i, current.next.next)
		}
	}

	result := highscore(score)

	fmt.Println("Done:", result, time.Since(t))
}

func highscore(s []int) int {
	best := -1
	for _, v := range s {
		if v > best {
			best = v
		}
	}
	return best
}

func insert(n, i int, s []int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = n
	return s
}

func remove(i int, s []int) ([]int, int) {
	n := s[i]
	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]
	return s, n
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func insertion(v int, n *Node) *Node {
	result := &Node{v, n.prev, n}
	result.prev.next = result
	result.next.prev = result
	return result
}

func removal(n *Node) *Node {
	result := n.next
	result.prev = n.prev
	result.prev.next = result
	return result
}

const INPUT1 = `9 players; last marble is worth 25 points`

const INPUT = `458 players; last marble is worth 72019 points`
