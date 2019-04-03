package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")
	p1(INPUT_REDDIT5) // INPUT -> 6208/1039, TAEL -> 11843/1078
	p2(INPUT_REDDIT5) // REDDIT -> ?/1087, REDDIT2 -> 10603/952, REDDIT3 -> 6256/973

	// findBreaker()
}

const (
	TORCH   = 1
	GEAR    = 2
	NEITHER = 4

	ROCKY  = 3 // TORCH & GEAR
	WET    = 6 // GEAR & NEITHER
	NARROW = 5 // TORCH & NEITHER
)

var mapping = []byte{'.', '=', '|'}
var terrain = []byte{ROCKY, WET, NARROW}

func findBreaker() {

	t := time.Now()

	var depth int
	for depth = 6000; depth <= 9999; depth += 3 {

		var w, h int

		w = 7
		h = 770

		cave := newCave(w+2, h+2)
		for y, row := range cave {
			for x := range row {
				if x == w && y == h {
					cave[y][x] = depth
				} else {
					cave[y][x] = gerosion(x, y, depth, cave)
				}

			}
		}

		if cave[h-1][w]%3 == 1 && cave[h+1][w]%3 == 1 && cave[h][w-1]%3 == 1 && cave[h][w+1]%3 == 1 {
			break
		}
	}

	fmt.Println("Depth:", depth, time.Since(t))
}

func p1(input string) {

	t := time.Now()

	lines := strings.SplitN(input, "\n", 2)
	depth, _ := strconv.Atoi(strings.SplitN(lines[0], " ", 2)[1])

	var w, h int
	scanned, err := fmt.Sscanf(lines[1], "target: %d,%d", &w, &h)
	if err != nil || scanned != 2 {
		panic("Couldn't scan target! " + lines[1])
	}

	risk := 0

	cave := newCave(w+1, h+1)
	for y, row := range cave {
		for x := range row {
			if x == w && y == h {
				cave[y][x] = depth
			} else {
				cave[y][x] = gerosion(x, y, depth, cave)
			}

			// ero := mapping[gerosion(cave[y][x], depth)%3]
			// fmt.Print(ero)

			risk += cave[y][x] % 3
		}
		// fmt.Println()
	}

	fmt.Println("Done:", risk, time.Since(t))
}

func p2(input string) {

	t := time.Now()

	lines := strings.SplitN(input, "\n", 2)
	depth, _ := strconv.Atoi(strings.SplitN(lines[0], " ", 2)[1])

	var w, h int
	scanned, err := fmt.Sscanf(lines[1], "target: %d,%d", &w, &h)
	if err != nil || scanned != 2 {
		panic("Couldn't scan target! " + lines[1])
	}

	cave := newCave(3*max(w, h)/2, 3*max(h, w)/2)
	for y, row := range cave {
		for x := range row {
			if x == w && y == h {
				cave[y][x] = depth
			} else {
				cave[y][x] = gerosion(x, y, depth, cave)
			}

			// ero := mapping[cave[y][x]%3]
			// fmt.Printf("%c", ero)
		}
		// fmt.Println()
	}

	target := Pos{w, h}
	origin := &Item{
		pos:      Pos{0, 0},
		tool:     TORCH,
		time:     0,
		from:     nil,
		priority: heuristic(target, Pos{0, 0}),
	}

	pq := make(PriorityQueue, 0, w+h)
	heap.Push(&pq, origin)

	type State struct {
		pos  Pos
		tool byte
	}

	seen := make(map[State]*Item)
	seen[State{origin.pos, origin.tool}] = origin

	var result *Item

	for pq.Len() > 0 {

		here := heap.Pop(&pq).(*Item)

		if here.pos == target {
			result = here
			break
		}

		current := terrain[cave[here.pos.y][here.pos.x]%3]

		for _, cp := range adj(here.pos) {

			ahead := terrain[cave[cp.y][cp.x]%3]
			// if cp == target {
			// 	ahead = 'X'
			// }

			// rocky 0 .	wet 1 =		narrow 2 |
			// item = pos, tool, time, prio

			var newTool byte
			var newTime int

			if cp == target && here.tool == GEAR {
				newTool, newTime = TORCH, here.time+1+7
			} else if cp == target && current == WET {
				newTool, newTime = TORCH, here.time+1+7+7
			} else if here.tool&ahead > 0 {
				newTool, newTime = here.tool, here.time+1
			} else {
				newTool, newTime = current&ahead, here.time+1+7
			}

			// if here.tool == ' ' && (ahead == '=' || ahead == '|') {
			// 	newTool, newTime = ' ', here.time+1
			// } else if here.tool == ' ' && ahead == '.' && current == '=' {
			// 	newTool, newTime = 'C', here.time+1+7
			// } else if here.tool == ' ' && ahead == '.' && current == '|' {
			// 	newTool, newTime = 'T', here.time+1+7

			// } else if here.tool == 'T' && (ahead == '.' || ahead == '|') {
			// 	newTool, newTime = 'T', here.time+1
			// } else if here.tool == 'T' && ahead == '=' && current == '.' {
			// 	newTool, newTime = 'C', here.time+1+7
			// } else if here.tool == 'T' && ahead == '=' && current == '|' {
			// 	newTool, newTime = ' ', here.time+1+7

			// } else if here.tool == 'C' && (ahead == '.' || ahead == '=') {
			// 	newTool, newTime = 'C', here.time+1
			// } else if here.tool == 'C' && ahead == '|' && current == '.' {
			// 	newTool, newTime = 'T', here.time+1+7
			// } else if here.tool == 'C' && ahead == '|' && current == '=' {
			// 	newTool, newTime = ' ', here.time+1+7

			// } else if ahead == 'X' {
			// 	if here.tool == 'T' {
			// 		newTool, newTime = 'T', here.time+1
			// 	} else if here.tool == 'C' || current == '|' {
			// 		newTool, newTime = 'T', here.time+1+7
			// 	} else { // wet + nothing
			// 		newTool, newTime = 'T', here.time+1+7+7
			// 	}
			// }

			it := seen[State{cp, newTool}]
			if it == nil {
				it = &Item{pos: cp, tool: newTool, time: newTime, from: here, priority: newTime + heuristic(target, cp)}
				heap.Push(&pq, it)
				seen[State{cp, newTool}] = it
			} else if it.time > newTime {
				pq.update(it, here, newTime, newTime+heuristic(target, cp))
			}
		}
	}

	// draw(result, cave)

	fmt.Println("Done 2:", result.time, time.Since(t))
}

func draw(path *Item, cave [][]int) {

	for y, row := range cave {
		for x := range row {
			cave[y][x] = int(mapping[cave[y][x]%3])
		}
	}

	from := path
	for from != nil {
		cave[from.pos.y][from.pos.x] = '#'
		from = from.from
	}

	cave[0][0] = 'M'
	cave[path.pos.y][path.pos.x] = 'T'

	var b bytes.Buffer
	for y, row := range cave {
		for x := range row {
			b.WriteByte(byte(cave[y][x]))
		}
		fmt.Println(b.String())
		b.Reset()
	}

}

// Use Manhattan distance to shrink search space
func heuristic(a, b Pos) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func gerosion(x, y, depth int, cave [][]int) int {
	result := 0

	if y == 0 && x == 0 {
		result = 0
	} else if y == 0 {
		result = x * 16807
	} else if x == 0 {
		result = y * 48271
	} else {
		result = cave[y][x-1] * cave[y-1][x]
	}

	result = (result + depth) % 20183

	return result
}

func newCave(x, y int) [][]int {
	cave := make([][]int, y)

	for i := range cave {
		cave[i] = make([]int, x)
	}

	return cave
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Pos struct {
	x int
	y int
}

func adj(p Pos) []Pos {
	result := make([]Pos, 0, 4)

	// Up, right, down, left
	if p.y-1 >= 0 {
		result = append(result, Pos{p.x, p.y - 1})
	}
	result = append(result, Pos{p.x + 1, p.y})
	result = append(result, Pos{p.x, p.y + 1})
	if p.x-1 >= 0 {
		result = append(result, Pos{p.x - 1, p.y})
	}

	return result
}

// Container for PQ
type Item struct {
	pos      Pos
	tool     byte
	time     int
	priority int
	from     *Item
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item, from *Item, time, priority int) {
	item.time = time
	item.from = from
	item.priority = priority
	heap.Fix(pq, item.index)
}

const WETWETWET = `depth: 6036
target: 7,770`

const INPUT2 = `depth: 510
target: 4,4`

const INPUT1 = `depth: 510
target: 10,10`

const INPUT = `depth: 10647
target: 7,770`

const INPUT_TAEL = `depth: 4080
target: 14,785`

const INPUT_REDDIT = `depth: 6969
target: 9,796`

const INPUT_REDDIT2 = `depth: 6084
target: 14,709`

const INPUT_REDDIT3 = `depth: 5913
target: 8,701`

const INPUT_REDDIT4 = `depth: 11394
target: 7,701`

const INPUT_REDDIT5 = `depth: 11739
target: 11,718`
