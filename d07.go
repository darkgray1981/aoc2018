package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")
	p1()
	p2()
}

const WORKER_LIMIT = 5
const OFFSET_TIME = 60

func p1() {

	t := time.Now()

	afters := make(map[rune][]rune)
	befores := make(map[rune][]rune)
	used := make(map[rune]bool)

	for _, s := range strings.Split(INPUT, "\n") {
		var before, after rune

		scanned, _ := fmt.Sscanf(s, "Step %c must be finished before step %c can begin.", &before, &after)
		if scanned != 2 {
			panic("Couldn't scan! " + s)
		}

		afters[after] = append(afters[after], before)
		befores[before] = append(befores[before], after)
		used[before] = true
		used[after] = true
	}

	var order string

	for len(used) > 0 {
		for c := 'A'; c <= 'Z'; c++ {
			if _, ok := used[c]; !ok {
				continue
			}

			if _, ok := afters[c]; ok {
				continue
			}

			for _, step := range befores[c] {
				afters[step] = remove(c, afters[step])
				if len(afters[step]) == 0 {
					delete(afters, step)
				}
			}

			delete(used, c)
			order += string(c)
			break
		}
	}

	fmt.Println("Done:", order, time.Since(t))
}

func p2() {

	t := time.Now()

	afters := make(map[rune][]rune)
	befores := make(map[rune][]rune)
	used := make(map[rune]bool)

	for _, s := range strings.Split(INPUT, "\n") {
		var before, after rune

		scanned, _ := fmt.Sscanf(s, "Step %c must be finished before step %c can begin.", &before, &after)
		if scanned != 2 {
			panic("Couldn't scan! " + s)
		}

		afters[after] = append(afters[after], before)
		befores[before] = append(befores[before], after)
		used[before] = true
		used[after] = true
	}

	var order string

	type Job struct {
		Item     rune
		TimeDone int
	}

	var workers []Job
	var queue []rune

	for k := range used {
		if _, ok := afters[k]; !ok {
			queue = append(queue, k)
			delete(used, k)
		}
	}

	timestamp := 0

	for len(workers) > 0 || len(queue) > 0 {

		if len(workers) > 0 {
			sort.Slice(workers, func(i, j int) bool {
				if workers[i].TimeDone == workers[j].TimeDone {
					return workers[i].Item < workers[j].Item
				}
				return workers[i].TimeDone < workers[j].TimeDone
			})

			timestamp = workers[0].TimeDone
		}

		for len(workers) > 0 && timestamp >= workers[0].TimeDone {
			c := workers[0].Item

			for _, step := range befores[c] {
				afters[step] = remove(c, afters[step])
				if len(afters[step]) == 0 {
					delete(afters, step)
				}
			}

			workers = workers[1:]
			delete(used, c)
			order += string(c)
		}

		for k := range used {
			if _, ok := afters[k]; !ok {
				queue = append(queue, k)
				delete(used, k)
			}
		}
		sort.Slice(queue, func(i, j int) bool { return queue[i] < queue[j] })

		for len(queue) > 0 && len(workers) < WORKER_LIMIT {
			item := queue[0]
			queue = queue[1:]
			workers = append(workers, Job{item, timestamp + OFFSET_TIME + int(item-'A'+1)})
		}

	}

	fmt.Println("Done 2:", order, timestamp, time.Since(t))
}

func remove(item rune, s []rune) []rune {
	if len(s) == 0 {
		return []rune{}
	}

	for i, r := range s {
		if r == item {
			s[i], s[len(s)-1] = s[len(s)-1], s[i]
			return s[:len(s)-1]
		}
	}

	return s
}

const INPUT1 = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

const INPUT = `Step Y must be finished before step L can begin.
Step N must be finished before step D can begin.
Step Z must be finished before step A can begin.
Step F must be finished before step L can begin.
Step H must be finished before step G can begin.
Step I must be finished before step S can begin.
Step M must be finished before step U can begin.
Step R must be finished before step J can begin.
Step T must be finished before step D can begin.
Step U must be finished before step D can begin.
Step O must be finished before step X can begin.
Step B must be finished before step D can begin.
Step X must be finished before step V can begin.
Step J must be finished before step V can begin.
Step D must be finished before step A can begin.
Step K must be finished before step P can begin.
Step Q must be finished before step C can begin.
Step S must be finished before step E can begin.
Step A must be finished before step V can begin.
Step G must be finished before step L can begin.
Step C must be finished before step W can begin.
Step P must be finished before step W can begin.
Step V must be finished before step W can begin.
Step E must be finished before step W can begin.
Step W must be finished before step L can begin.
Step P must be finished before step E can begin.
Step T must be finished before step K can begin.
Step A must be finished before step G can begin.
Step G must be finished before step P can begin.
Step N must be finished before step S can begin.
Step R must be finished before step D can begin.
Step M must be finished before step G can begin.
Step Z must be finished before step L can begin.
Step M must be finished before step T can begin.
Step S must be finished before step L can begin.
Step S must be finished before step W can begin.
Step O must be finished before step J can begin.
Step Z must be finished before step D can begin.
Step A must be finished before step C can begin.
Step P must be finished before step V can begin.
Step A must be finished before step P can begin.
Step B must be finished before step C can begin.
Step R must be finished before step S can begin.
Step X must be finished before step S can begin.
Step T must be finished before step P can begin.
Step Y must be finished before step E can begin.
Step G must be finished before step E can begin.
Step Y must be finished before step K can begin.
Step J must be finished before step P can begin.
Step I must be finished before step Q can begin.
Step E must be finished before step L can begin.
Step X must be finished before step J can begin.
Step T must be finished before step X can begin.
Step M must be finished before step O can begin.
Step K must be finished before step A can begin.
Step D must be finished before step W can begin.
Step H must be finished before step C can begin.
Step F must be finished before step R can begin.
Step B must be finished before step Q can begin.
Step M must be finished before step Q can begin.
Step D must be finished before step S can begin.
Step Y must be finished before step I can begin.
Step M must be finished before step K can begin.
Step S must be finished before step G can begin.
Step X must be finished before step L can begin.
Step D must be finished before step V can begin.
Step B must be finished before step X can begin.
Step C must be finished before step L can begin.
Step V must be finished before step L can begin.
Step Z must be finished before step Q can begin.
Step Z must be finished before step H can begin.
Step M must be finished before step S can begin.
Step O must be finished before step C can begin.
Step B must be finished before step A can begin.
Step U must be finished before step V can begin.
Step U must be finished before step A can begin.
Step X must be finished before step G can begin.
Step K must be finished before step C can begin.
Step T must be finished before step S can begin.
Step K must be finished before step G can begin.
Step U must be finished before step B can begin.
Step A must be finished before step E can begin.
Step F must be finished before step V can begin.
Step Q must be finished before step A can begin.
Step F must be finished before step Q can begin.
Step J must be finished before step L can begin.
Step O must be finished before step E can begin.
Step O must be finished before step Q can begin.
Step I must be finished before step K can begin.
Step I must be finished before step P can begin.
Step J must be finished before step D can begin.
Step Q must be finished before step P can begin.
Step S must be finished before step C can begin.
Step U must be finished before step P can begin.
Step S must be finished before step P can begin.
Step O must be finished before step B can begin.
Step Z must be finished before step F can begin.
Step R must be finished before step V can begin.
Step D must be finished before step L can begin.
Step Y must be finished before step T can begin.
Step G must be finished before step C can begin.`
