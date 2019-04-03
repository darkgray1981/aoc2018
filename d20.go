package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting...")
	// p1()
	// p2()
	// p12(INPUT, 1000)
	p12b("^NNNNN(EEEE|WWWW)SSSSS$", 10)
}

type Pos struct {
	x int
	y int
}

func p1() {

	t := time.Now()

	path := INPUT
	path = path[1 : len(path)-1]

	grid := make(map[Pos]byte)
	grid[Pos{0, 0}] = 'X'

	var buffer []byte

	count := 0
	farthest := 0

	var walk func(pi, bi, in int)
	walk = func(pi, bi, in int) {

		// fmt.Println("Start at", path[pi:pi+1], in)
		biOld := bi

		self := in
		pause := false
		done := false
		var i int

		for i = pi; i < len(path); i++ {

			d := path[i]

			if !done {

				if in == self {
					if d == '(' {
						in++
						if !pause {
							// fmt.Println("Ignite at", path[i+1:i+2], "from", path[pi:pi+1])
							walk(i+1, bi, in)
							done = true
						}
					} else if d == '|' {

						// Special loop case
						if path[i+1] == ')' {
							bi -= (i - pi) / 2
							break
						}

						pause = true
					} else if d == ')' {
						in--
						self--
						pause = false
					} else if !pause {
						// fmt.Println("[" + string(d) + "]")

						// if bi+1 > len(buffer) {
						// 	buffer = append(buffer, byte(d))
						// } else {
						// 	buffer[bi] = byte(d)
						// }
						bi++
					}
				} else if in > self {
					if d == '(' {
						in++
					} else if d == ')' {
						in--
					}
				}

				if in < self {
					panic("Went below surface")
				}
			} else if done {

				if in > self+1 {
					if d == '(' {
						in++
					} else if d == ')' {
						in--
					}
				} else if in == self+1 {
					if d == '(' {
						in++
					} else if d == '|' {
						// fmt.Println("Ignite at", path[i+1:i+2], "from", path[pi:pi+1])
						walk(i+1, bi, in)
					} else if d == ')' {
						break
					}
				} else if in <= self {
					if pause {
						panic("Paused under own level!")
					}
				}
			}
		}

		if !done {
			count++
			// fmt.Printf("\r%9d %9d %9d", count, pi, in)
			_ = buffer
			// fmt.Println(count, string(buffer[:bi]), bi, biOld)
			_ = biOld

			farthest = max(farthest, bi)
		}

	}

	walk(0, 0, 0)

	fmt.Println("Done:", farthest, time.Since(t))
}

func p2() {

	t := time.Now()

	doors := 1000
	path := INPUT
	path = path[1 : len(path)-1]

	// fmt.Println(path)

	grid := make(map[Pos]byte)
	grid[Pos{0, 0}] = 'X'

	var buffer []byte

	paths := make(map[string]bool)

	count := 0

	var walk func(pi, bi, in int) int
	walk = func(pi, bi, in int) int {

		// fmt.Println("Start at", path[pi:pi+1], in)
		biOld := bi

		self := in
		pause := false
		done := false

		beyond := 0

		var i int

		for i = pi; i < len(path); i++ {

			d := path[i]

			if !done {

				if in == self {
					if d == '(' {
						in++
						if !pause {
							// fmt.Println("Ignite at", path[i+1:i+2], "from", path[pi:pi+1])
							beyond += walk(i+1, bi, in)
							done = true
						}
					} else if d == '|' {

						// Special loop case
						if path[i+1] == ')' {
							bi -= (i - pi) / 2
							break
						}

						pause = true
					} else if d == ')' {
						in--
						self--
						pause = false
					} else if !pause {
						// fmt.Println("[" + string(d) + "]")

						if bi+1 > len(buffer) {
							buffer = append(buffer, byte(d))
						} else {
							buffer[bi] = byte(d)
						}
						bi++
					}
				} else if in > self {
					if d == '(' {
						in++
					} else if d == ')' {
						in--
					}
				}

				if in < self {
					panic("Went below surface")
				}
			} else if done {

				if in > self+1 {
					if d == '(' {
						in++
					} else if d == ')' {
						in--
					}
				} else if in == self+1 {
					if d == '(' {
						in++
					} else if d == '|' {
						// fmt.Println("Ignite at", path[i+1:i+2], "from", path[pi:pi+1])
						beyond += walk(i+1, bi, in)
					} else if d == ')' {
						break
					}
				} else if in <= self {
					if pause {
						panic("Paused under own level!")
					}
				}
			}
		}

		if !done {
			count++
			// fmt.Printf("\r%9d %9d %9d", count, pi, in)
			_ = buffer
			// fmt.Println(count, string(buffer[:bi]), bi, biOld)

			paths[string(buffer[:bi])] = true
		}

		if biOld >= doors {
			beyond += bi - biOld
		} else if bi >= doors {
			beyond += bi - doors + 1
		}

		return beyond
	}

	walk(0, 0, 0)

	origin := Pos{0, 0}

	for p := range paths {
		me := origin
		for _, d := range p {
			me = move(me, byte(d), grid)
		}
	}

	// draw(grid)

	result := 0

	for _, dist := range chart(doors, grid) {
		if dist >= doors {
			result++
		}
	}

	fmt.Println("Done 2:", result, time.Since(t))
}

func p12(input string, doors int) {

	t := time.Now()

	path := input

	path = path[1 : len(path)-1]

	// fmt.Println(path)

	grid := make(map[Pos]byte)
	grid[Pos{0, 0}] = 'X'

	var buffer []byte

	paths := make(map[string]bool)

	count := 0

	var walk func(pi, bi, in int) int
	walk = func(pi, bi, in int) int {

		// fmt.Println("Start at", path[pi:pi+1], in)
		biOld := bi

		self := in
		pause := false
		done := false

		beyond := 0

		var i int

		for i = pi; i < len(path); i++ {

			d := path[i]

			if !done {

				if in == self {
					if d == '(' {
						in++
						if !pause {
							// fmt.Println("Ignite at", path[i+1:i+2], "from", path[pi:pi+1])
							beyond += walk(i+1, bi, in)
							done = true
						}
					} else if d == '|' {

						// Special loop case
						if path[i+1] == ')' {
							bi -= (i - pi) / 2
							break
						}

						pause = true
					} else if d == ')' {
						in--
						self--
						pause = false
					} else if !pause {
						// fmt.Println("[" + string(d) + "]")

						if bi+1 > len(buffer) {
							buffer = append(buffer, byte(d))
						} else {
							buffer[bi] = byte(d)
						}
						bi++
					}
				} else if in > self {
					if d == '(' {
						in++
					} else if d == ')' {
						in--
					}
				}

				if in < self {
					panic("Went below surface")
				}
			} else if done {

				if in > self+1 {
					if d == '(' {
						in++
					} else if d == ')' {
						in--
					}
				} else if in == self+1 {
					if d == '(' {
						in++
					} else if d == '|' {
						// fmt.Println("Ignite at", path[i+1:i+2], "from", path[pi:pi+1])
						beyond += walk(i+1, bi, in)
					} else if d == ')' {
						break
					}
				} else if in <= self {
					if pause {
						panic("Paused under own level!")
					}
				}
			}
		}

		if !done {
			count++

			paths[string(buffer[:bi])] = true
		}

		if biOld >= doors {
			beyond += bi - biOld
		} else if bi >= doors {
			beyond += bi - doors + 1
		}

		return beyond
	}

	foo := walk(0, 0, 0)

	fmt.Println("Attempted result:", foo)

	origin := Pos{0, 0}

	for p := range paths {
		me := origin
		for _, d := range p {
			me = move(me, byte(d), grid)
		}
	}

	// draw(grid)
	charted := chart(doors, grid)

	result := 0

	for _, dist := range charted {
		result = max(result, dist)
	}

	fmt.Println("Done:", result, time.Since(t))

	t = time.Now()

	result = 0

	for _, dist := range charted {
		if dist >= doors {
			result++
		}
	}

	fmt.Println("Done 2:", result, time.Since(t))
}

func p12b(input string, doors int) {

	t := time.Now()

	path := input

	// Strip ^ and $ from input before we start
	path = path[1 : len(path)-1]

	// fmt.Println(path)

	grid := make(map[Pos]byte)

	var walk func(p Pos, pi, in int)
	walk = func(p Pos, pi, in int) {

		level := in
		pause, done := false, false

		var i int
		for i = pi; i < len(path); i++ {

			d := path[i]

			if !done {

				if in == level {
					if d == '(' {
						in++
						if !pause {
							walk(p, i+1, in)
							done = true
						}
					} else if d == '|' {

						// Special loop case
						if path[i+1] == ')' {
							break
						}

						pause = true
					} else if d == ')' {
						in--
						level--
						pause = false
					} else if !pause {
						p = move(p, byte(d), grid)
					}
				} else if in > level {
					if d == '(' {
						in++
					} else if d == ')' {
						in--
					}
				}

				if in < level {
					panic("Went below surface")
				}
			} else if done {

				if in > level+1 {
					if d == '(' {
						in++
					} else if d == ')' {
						in--
					}
				} else if in == level+1 {
					if d == '(' {
						in++
					} else if d == '|' {
						walk(p, i+1, in)
					} else if d == ')' {
						break
					}
				} else if in <= level {
					if pause {
						panic("Paused under own level!")
					}
				}
			}
		}
	}

	origin := Pos{0, 0}
	grid[origin] = 'X'
	walk(origin, 0, 0)

	// draw(grid)
	charted := chart(doors, grid)

	result := 0
	for _, dist := range charted {
		result = max(result, dist)
	}

	fmt.Println("Done:", result, time.Since(t))

	t = time.Now()

	result = 0
	for _, dist := range charted {
		if dist >= doors {
			result++
		}
	}

	fmt.Println("Done 2:", result, time.Since(t))
}

func chart(limit int, grid map[Pos]byte) map[Pos]int {

	origin := Pos{0, 0}

	frontier := []Pos{origin}
	seen := make(map[Pos]int)

	seen[origin] = -1

	distance := 0

	// Best first search for routes
	for len(frontier) > 0 {

		distance++

		for count := len(frontier); count > 0; count-- {

			p := frontier[0]
			frontier = frontier[1:]

			for _, cp := range open(p, grid) {
				if seen[cp] == 0 {
					frontier = append(frontier, cp)
					seen[cp] = distance
				}
			}
		}
	}

	return seen
}

func open(p Pos, grid map[Pos]byte) []Pos {
	result := make([]Pos, 0, 4)

	if grid[Pos{p.x, p.y - 1}] == '-' {
		result = append(result, Pos{p.x, p.y - 2})
	}
	if grid[Pos{p.x + 1, p.y}] == '|' {
		result = append(result, Pos{p.x + 2, p.y})
	}
	if grid[Pos{p.x - 1, p.y}] == '|' {
		result = append(result, Pos{p.x - 2, p.y})
	}
	if grid[Pos{p.x, p.y + 1}] == '-' {
		result = append(result, Pos{p.x, p.y + 2})
	}

	return result
}

func draw(grid map[Pos]byte) {

	minX, maxX, minY, maxY := 0, 0, 0, 0

	for p := range grid {
		minX, maxX = min(p.x, minX), max(p.x, maxX)
		minY, maxY = min(p.y, minY), max(p.y, maxY)
	}

	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			c := grid[Pos{x, y}]
			if c == 0 {
				c = '#'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}

}

func move(p Pos, d byte, grid map[Pos]byte) Pos {

	if d == 'N' {
		grid[Pos{p.x, p.y - 1}] = '-'
		p = Pos{p.x, p.y - 2}
	} else if d == 'E' {
		grid[Pos{p.x + 1, p.y}] = '|'
		p = Pos{p.x + 2, p.y}
	} else if d == 'S' {
		grid[Pos{p.x, p.y + 1}] = '-'
		p = Pos{p.x, p.y + 2}
	} else if d == 'W' {
		grid[Pos{p.x - 1, p.y}] = '|'
		p = Pos{p.x - 2, p.y}
	} else {
		panic(fmt.Sprintf("Pos=%v, dir=%c\n", p, d))
	}

	grid[p] = '.'

	return p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// const INPUT_TEST = `^abc(def|ghi)$`

const INPUT_TEST = `^abc(def|ghi)jkl(mn|op(qr|st))uv(wx|y|z)$`

const INPUT1 = `^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$`
const INPUT2 = `^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$`
const INPUT3 = `^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$`

const INPUT_FOO = `^N(N(N|N(N|N(N(N(N(N(N(NN(N|N(N(N(N(N|NN(N|N))|N)|NN)|N))|N)|N(N|N(N(N|NN(N|N(N(N(NN|N(N(NN(N(N(N|N(N|N(N(N|NN(NN(N(N(N(N(N(N|N(NN(N|N(N|N(N(N(N|N(N|N(N|N)))|NN)|N)))|N))|N)|N(NN(N(NN|N(N(N(N|N(N|N(N(NN|N(N(N(NN(N|N(N(N(N|N)|N)|N))|N(N(N|N(N(N|N)|N(N|NNN)))|N(N(N(N|N)|N)|N)))|NNN(N|N(N|NN(N|N(N|N(NN(N|N(N|N(N(N|N(N|N)|N)|N(N|N(NN(NN|N(N|N(N(N(N(N|N(N|N(N|N(N(N(N(N|NNN)|N)|N)|N))))|N(N|NN(NN|N(N(N(N(N|N)|N(N|N(N|N)))|N(N|NN(N|NN(N|N(N(N|N)|N(N|N))))))|N))))|N)|N)))|N(NN|N))))))|N(N|N(N(N(N|N(N(N|N(N(N(N|N(N|N|N(N|N)))|N(N(N(N(N(NN(N|NN(N|N(N(N(NNN(N|N(N|N))|N(N(N(N|N)|NNN(N|N))|NN))|N)|N(N|N(N(N|N(N|N(N|N(N|N(NN(N|N)|N(NN(N|N(N|N(N(N(N(N(N|N(N(N(N(N(N(N|N)|N(N|N))|N)|N(N|N(N(N|N(N(N(N(N(N|N)|N(N(N|N(N|N))|N(N|NN)))|N)|N(N(N(N|N(N(N|N)|N(NN|N)))|N(N(N(N|N)|N(N|N(N|N)))|N|N))|N(N(N(N(N|N(N|N(NNN(N|NN(N|N(N(N(N(NN|NN(N|N(N(N(N|NN(N(N(N|N)|N(N(N|NNNN(N(N(N(N(N(NN(N(N|N(N(N|N(NN(N(NN(N(N(N(N|NN)|N(NN(NN|N)|N))|N(NN|N))|N(N(NN(N(N|N(N|NN(N(N|NN)|N(N(N|N)|NN(NNN(N(N|N(N|N(N|N(N(N(NN|N(N|N))|N(N|N(NNN(N|N(N(N(NN(N|N)|N(N|N(N|N(N(N(NN|N)|N(N|NN))|N(N|N)))))|N)|N(N|NN(N|N(N|N(N(N|N(N|N))|N(N(NN|NN)|N)))))))|N(N|N))))|N))))|N(N|N))|N)))))|N)|N(N(N|N(N|N))|N))|N))|NN)|N(NN(N(N(NN|N(N|N(N|N(NN(N(N(N(N(N(N(N|N)|N(N(N(N(NN(N(N|N(N|N(N(N|N)|N)))|N)|NNN(N|N(N|N(NN(N|N(N|N)|N)|N))))|N)|N)|N(NN(N|N)|N)))|N(N(NN(N(N|NNN)|N(NN|N))|N(N|N))|N))|N)|N)|N)|NN(NN(N|N)|N))|N))))|NN)|N(N|N))|N))|N))|N))|N)|N)|N(N|N))|N(N(N|N)|N(N(N|N)|N)))|N(N(N|N(N|N(N(N|NN(NN(N|N)|N(NN(NN(N|N(N(N|N(N|NN))|N(N(N|N(N(N|N(N(N|NN(N(N(N|N(N|N(N|N(N(N|N(N|N(N(N|N(N(N|N)|NNN))|N(N|N))))|N(N(N|N)|NN(N(N|N(NN(N|N(N|N(N(N|N(NN(N(NN(N(N|N(N(N(N(N|N)|N(N|N(N(N(NN(N|NN)|N(N|N(N|NN(N(N(N(N|N)|N(N(N|N(NN|N(N(N(N(N(N|NN(N(N|N(NN(N|NNN(N|N(NN(NNNN(N(N(N(N|N(N|NN))|N)|N(N|N(N|N(NNNN(NN(N|N)|N)|N))))|NN)|N)|N)))|N))|N))|N)|N)|NN(NN(N(N|N)|N)|N))|N)))|NN))|N)|N))))|N(NNN(N|N(NN|NN))|N))|NN)))|N)|N(NN|N(N(N|N(N|N(N(N|N(N(N|N(NN(N(N(N(N(N(N(N(N(N(N|N)|N(N|N))|N)|N)|N(N(N|N(N(N(N(N(N|N)|N)|NN)|N(N(N(NN(N|N(N|NN(N(N(N(N|N(N|N(N(N|N(N(N(N|N(N(N|N)|N))|N(N|N))|N))|N(N|N(N|NN(N(N|N(NNN|N(N(N|N(N(N(NN(N|N)|NN)|N)|N))|N(N(NN(N(N(N(N(N|NN(N|N))|N(N(N(N|N(N|N(NN(N|N)|N)))|N(N(N|N(N(N|N(NN|N))|N(N(NNN|N(N|NN))|N)))|N(N|NN(N|N))))|N(N(N|N(N(N|N(N|N(N|N)))|N(N|N)))|N)))|N)|N)|N)|N(N|N(N|N(N(NN|N)|N(N|N(NN(N(N|N(N|N))|N)|N))))))|N(N(N|N)|N)))))|N(N|N(N|NN(N(N|N(N|N))|N)))))))))|N)|N(N(NN(NNN(N|NNNN(N(N(N|N(N|N))|NN)|N(N|N(N|N(N(N|N)|N))|N)))|N)|NN(N|N(N(N(N|NN(NN(N|N(N|NN(N(N(N(NN(N(N|N)|N|N(N|NN))|N)|N(N(N|N(N|N(N(N(N(N(N(N|N(N|N))|N(N(N(N|N)|N(N(N|N(N|N))|N))|N))|NN)|N)|N)|N)))|N))|N(N(N(N|N)|N)|N))|N)))|N))|N)|N)))|N))|N)))|N)|N)|N))|N))|N(N(N(N|N)|N)|NN)))|N)|N(N(N|N)|N))|N)|N)|N)|N))|N(NN|N)))|N(N(N|N(N(N|N(NN|N(N(N(N(N|NN(N|N(NN(N(N|NN(N(N(N(N(N(N(N|N(N(N(N(N|NN)|N)|N)|N(N|N(N(N(N|N)|NN(N(N|N(N|N(N(N(N|N)|N)|NN(N(N|N)|N(N(N|N(N|N))|N(N(NN(N(N|N(N|N))|N)|NN(N|N(N|N(N(N|NN(N|N(N(N(N|NN(N(N|NN(N|NN(N(N|N)|NN)))|N(N|N(N|N(N|N(N(N(N|N)|N(NN|N))|N))))))|N)|NN(N(N(N(N|N(NN(N(NN(N(N|N)|NN(N|NN(N|N)))|N(N|N|N))|N)|N))|N)|N)|N))))|N))))|N(N|N)|N))))))|N))|N(N|N(N(N|NN(N(N|N(N|N(N(NNNN(N|N(N(N(N(N|N)|N(N|N(N|N)))|N)|N(NN(N|NN)|N(N(NN(N|NNN(N|N(N(N(N|N(N|N(N(N|N(N(N(N(N(N|N(N|N))|N)|N)|N(N|NN(N(N|N)|N)))|N(N|N)))|N(N(N|N(N(N(N(N(NN(N(N(NN(N|NNN(N|N))|NNN)|NN(N|N))|N(N|N(N(N(N(NN(N(N(N|N)|N(NN|N))|N)|NN(N(N|N)|N(N(N|NN)|N)))|N(N|N))|N)|N)))|N)|N)|N)|N)|N))|N(N|N)))))|N)|N)))|N(N|N))|N))))|N)|N)))|N))|N(N|N)))))))|N)|N)|N)|N)|N)|N(N|N(N(NN|N)|N))))|N)|N(N(N(N(N|N)|N(N(N|N(N|N))|N))|N)|N(N(N|N)|NN)))))|N)|N(N(NNN|NN(N|N))|N))|N)))|N))|N(N|N)))))|N))))|N)|N)|NN(N(NN|N(N(N|N)|N(N|N)))|N))|N))|N)))|N))|N(N(N(N|N)|N(N|N))|NNN)))))))|N)|NN))|N))|N(N|N)))|N)))|N)|N)))|N)))|N))|N)|N(N(NNN|N)|N)))|N))|N))|N(N(NN|NN)|N))|N(N|N))))|N(N|N))|N)|N(NN(N|N)|N))))|N)))|N)|N)|N)))|N))|N(N(N|N|N)|N))))|N)|N))|N)|N)|N(N(N(N(N|N)|N(N(N(N(NNN|N(N|N))|N)|N(N(NN|NN(N|NN))|N(N|N|N(N(N|N)|N))))|N))|N)|NNN(N|N)))|N)))|N))))))|N)))))|N(NN|N(N(N|N)|N(N|N|NN))))|N)|N)|N(N(N(N|N(N(N|N(N(N(N|N)|N(NN(N|N(N(N|N)|N(N(N|NN(N(N|N)|N))|N)))|N(N(N(N|N(N(N|N)|N))|N(N|N(N|NN(N|N))))|N)))|N))|N(N|N(N|N(N|NN(N(N|N)|N))))))|N)|N))|N))|N))|N(N(NN|NNNN)|N)))|N)|N))))))))|N))|N)))|N(N|N(N(NN(N(N(N|N(N(N|N)|N))|N)|NN(N(N|N)|N))|N)|NN)))|N(N(N(N|N)|N(N|N))|N(N|N))))|N)|N))|N)|N)|N)|N))|N)))|NN)|N)|N)|N))|N)|N(N|N(NN|NN)))))|N)))|N)|N)|N)|N)))|N)$`
const INPUT_BAR = `^N(N(N|N(N|N(N(N(N(N(N(N(N|)N(N|N(N(N(N(N|N(N|)N(N|N))|N)|N(N|)N)|N))|N)|N(N|N(N(N|N(N|)N(N(N|)|N(N(N(N(N|)N|N(N(N(N|)N(N(N(N|N(N|N(N(N|N(N|)N(N(N|)N(N(N(N(N(N(N|N(N(N|)N(N|N(N|N(N(N(N|N(N|N(N|N)))|N(N|)N)|N)))|N))|N)|N(N(N|)N(N(N(N|)N|N(N(N(N|N(N|N(N(N(N|)N(N|)|N(N(N(N(N|)N(N|N(N(N(N|N)|N)|N))|N(N(N|N(N(N|N)|N(N|N(N|)N(N|)N)))|N(N(N(N|N)|N)|N)))|N(N|)N(N|)N(N|N(N|N(N|)N(N|N(N|N(N(N|)N(N|N(N|N(N(N|N(N|N)|N)|N(N|N(N(N|)N(N(N|)N(N|)|N(N|N(N(N(N(N|N(N|N(N|N(N(N(N(N|N(N|)N(N|)N)|N)|N)|N))))|N(N|N(N|)N(N(N|)N(N|)|N(N(N(N(N|N)|N(N|N(N|N)))|N(N|N(N|)N(N|N(N|)N(N|N(N(N|N)|N(N|N))))))|N))))|N)|N)))|N(N(N|)N(N|)|N))))))|N(N|N(N(N(N|N(N(N|N(N(N(N|N(N|N|N(N|N)))|N(N(N(N(N(N(N|)N(N|N(N|)N(N|N(N(N(N(N|)N(N|)N(N|N(N|N))|N(N(N(N|N)|N(N|)N(N|)N(N|N))|N(N|)N))|N)|N(N|N(N(N|N(N|N(N|N(N|N(N(N|)N(N|N)|N(N(N|)N(N|N(N|N(N(N(N(N(N|N(N(N(N(N(N(N|N)|N(N|N))|N)|N(N|N(N(N|N(N(N(N(N(N|N)|N(N(N|N(N|N))|N(N|N(N|)N)))|N)|N(N(N(N(N|)|N(N(N|N)|N(N(N|)N|N)))|N(N(N(N|N)|N(N|N(N|N)))|N|N))|N(N(N(N(N|N(N|N(N(N|)N(N|)N(N|N(N|)N(N|N(N(N(N(N(N|)N|N(N|)N(N(N|)|N(N(N(N|N(N|)N(N(N(N|N)|N(N(N|N(N|)N(N|)N(N|)N(N(N(N(N(N(N(N|)N(N(N|N(N(N|N(N(N|)N(N(N(N|)N(N(N(N(N|N(N|)N)|N(N(N|)N(N(N|)N|N)|N))|N(N(N|)N|N))|N(N(N(N|)N(N(N|N(N|N(N|)N(N(N|N(N|)N)|N(N(N|N)|N(N|)N(N(N|)N(N|)N(N(N|N(N(N|)|N(N|N(N(N(N(N|)N|N(N|N))|N(N|N(N(N|)N(N|)N(N|N(N(N(N(N|)N(N|N)|N(N|N(N|N(N(N(N(N|)N|N)|N(N|N(N|)N))|N(N|N)))))|N)|N(N|N(N|)N(N|N(N|N(N(N|N(N|N))|N(N(N(N|)N(N|)|N(N|)N)|N)))))))|N(N|N))))|N))))|N(N|N))|N)))))|N)|N(N(N|N(N|N))|N(N|)))|N))|N(N|)N)|N(N(N|)N(N(N(N(N|)N|N(N|N(N|N(N(N|)N(N(N(N(N(N(N(N|N)|N(N(N(N(N(N|)N(N(N|N(N|N(N(N|N)|N)))|N)|N(N|)N(N|)N(N|N(N|N(N(N|)N(N|N(N|N)|N)|N))))|N)|N)|N(N(N|)N(N|N)|N)))|N(N(N(N|)N(N(N|N(N|)N(N|)N)|N(N(N|)N|N))|N(N|N))|N))|N)|N)|N)|N(N|)N(N(N|)N(N|N)|N))|N))))|N(N|)N)|N(N|N))|N))|N))|N))|N)|N)|N(N|N))|N(N(N|N)|N(N(N|N)|N)))|N(N(N|N(N|N(N(N|N(N|)N(N(N|)N(N|N)|N(N(N|)N(N(N|)N(N(N|)|N(N(N|N(N|N(N|)N))|N(N(N|N(N(N|N(N(N|N(N|)N(N(N(N|N(N(N|)|N(N|N(N(N|N(N|N(N(N(N|)|N(N(N|N)|N(N|)N(N|)N))|N(N|N))))|N(N(N|N)|N(N|)N(N(N|N(N(N|)N(N|N(N|N(N(N|N(N(N|)N(N(N(N|)N(N(N|N(N(N(N(N|N)|N(N|N(N(N(N(N|)N(N|N(N|)N)|N(N|N(N|N(N|)N(N(N(N(N|N)|N(N(N|N(N(N|)N|N(N(N(N(N(N|)(N|N(N|)N(N(N|N(N(N|)N(N|N(N|)N(N|)N(N|N(N(N|)N(N(N|)N(N|)N(N|)N(N(N(N(N|N(N|N(N|)N))|N)|N(N|N(N|N(N(N|)N(N|)N(N|)N(N(N|)N(N|N)|N(N|))|N))))|N(N|)N)|N)|N)))|N))|N))|N)|N)|N(N|)N(N(N|)N(N(N|N)|N)|N))|N)))|N(N|)N))|N)|N))))|N(N(N|)N(N|)N(N|N(N(N|)N|N(N|)N))|N))|N(N|)N)))|N)|N(N(N|)N|N(N(N|N(N|N(N(N|N(N(N|N(N(N|)N(N(N(N(N(N(N(N(N(N(N|N)|N(N|N))|N)|N)|N(N(N|N(N(N(N(N(N|N)|N)|N(N|)N)|N(N(N(N(N|)N(N|N(N|N(N|)N(N(N(N(N|N(N|N(N(N|N(N(N(N|N(N(N|N)|N))|N(N|N))|N))|N(N|N(N|N(N|)N(N(N|N(N(N|)N(N|)N|N(N(N|N(N(N(N(N|)N(N|N)|N(N|)N)|N)|N))|N(N(N(N|)N(N(N(N(N(N|N(N|)N(N|N))|N(N(N(N|N(N|N(N(N|)N(N|N)|N)))|N(N(N|N(N(N|N(N(N|)N|N))|N(N(N(N|)N(N|)N|N(N|N(N|)N))|N)))|N(N|N(N|)N(N|N))))|N(N(N|N(N(N|N(N|N(N|N)))|N(N|N)))|N)))|N)|N)|N)|N(N|N(N|N(N(N(N|)N|N)|N(N|N(N(N|)N(N(N|N(N(N|)|N))|N)|N))))))|N(N(N|N)|N)))))|N(N|N(N|N(N|)N(N(N|N(N|N))|N)))))))))|N)|N(N(N(N|)N(N(N|)N(N|)N(N|N(N|)N(N|)N(N|)N(N(N(N|N(N|N))|N(N|)N)|N(N|N(N|N(N(N|N)|N))|N)))|N)|N(N|)N(N|N(N(N(N|N(N|)N(N(N|)N(N|N(N|N(N|)N(N(N(N(N(N|)N(N(N|N)|N|N(N|N(N|)N))|N)|N(N(N|N(N|N(N(N(N(N(N(N|N(N(N|)|N))|N(N(N(N|N)|N(N(N|N(N|N))|N))|N))|N(N|)N)|N)|N)|N)))|N(N|)))|N(N(N(N|N)|N)|N))|N)))|N))|N)|N)))|N))|N)))|N)|N)|N))|N))|N(N(N(N|N)|N)|N(N|)N)))|N)|N(N(N|N)|N))|N)|N)|N)|N))|N(N(N|)N|N)))|N(N(N|N(N(N|N(N(N|)N|N(N(N(N(N|N(N|)N(N|N(N(N|)N(N(N|N(N|)N(N(N(N(N(N(N(N|N(N(N(N(N|N(N|)N)|N)|N)|N(N|N(N(N(N|N)|N(N|)N(N(N|N(N|N(N(N(N|N)|N)|N(N|)N(N(N|N)|N(N(N|N(N|N))|N(N(N(N|)N(N(N|N(N|N))|N)|N(N|)N(N|N(N|N(N(N|N(N|)N(N|N(N(N(N|N(N|)N(N(N|N(N|)N(N|N(N|)N(N(N|N)|N(N|)N)))|N(N|N(N|N(N|N(N(N(N|N)|N(N(N|)N|N))|N))))))|N)|N(N|)N(N(N(N(N|N(N(N|)N(N(N(N|)N(N(N|N)|N(N|)N(N|N(N|)N(N|N)))|N(N|N|N))|N)|N))|N)|N)|N))))|N))))|N(N|N)|N))))))|N))|N(N|N(N(N|N(N|)N(N(N|N(N|N(N(N(N|)N(N|)N(N|)N(N|N(N(N(N(N|N)|N(N|N(N|N)))|N)|N(N(N|)N(N|N(N|)N)|N(N(N(N|)N(N|N(N|)N(N|)N(N(N|)|N(N(N(N(N|)|N(N|N(N(N|N(N(N(N(N(N|N(N|N))|N)|N)|N(N|N(N|)N(N(N|N)|N)))|N(N|N)))|N(N(N|N(N(N(N(N(N(N|)N(N(N(N(N|)N(N|N(N|)N(N|)N(N|N))|N(N|)N(N|)N)|N(N|)N(N|N))|N(N|N(N(N(N(N(N|)N(N(N(N|N)|N(N(N|)N|N))|N)|N(N|)N(N(N|N)|N(N(N|N(N|)N(N|))|N)))|N(N|N))|N)|N)))|N)|N)|N)|N)|N))|N(N|N(N|))))))|N)|N)))|N(N|N))|N))))|N)|N)))|N))|N(N|N)))))))|N)|N)|N)|N)|N)|N(N|N(N(N(N|)N|N)|N))))|N)|N(N(N(N(N|N)|N(N(N|N(N|N))|N))|N)|N(N(N|N)|N(N|)N)))))|N)|N(N(N(N|)N(N|)N|N(N|)N(N|N))|N))|N)))|N))|N(N|N)))))|N))))|N)|N)|N(N|)N(N(N(N|)N(N|)|N(N(N|N)|N(N|N)))|N))|N))|N)))|N))|N(N(N(N|N)|N(N|N))|N(N|)N(N|)N(N|))))))))|N)|N(N|)N))|N))|N(N|N)))|N)))|N)|N)))|N)))|N))|N)|N(N(N(N|)N(N|)N|N(N|))|N)))|N))|N(N|)))|N(N(N(N|)N|N(N|)N)|N))|N(N|N))))|N(N|N))|N)|N(N(N|)N(N|N)|N))))|N)))|N)|N)|N)))|N))|N(N(N|N|N)|N))))|N)|N))|N)|N)|N(N(N(N(N|N)|N(N(N(N(N(N|)N(N|)N|N(N|N))|N)|N(N(N(N|)N|N(N|)N(N|N(N|)N))|N(N(N|)|N|N(N(N|N)|N))))|N))|N(N|))|N(N|)N(N|)N(N|N)))|N)))|N))))))|N)))))|N(N(N|)N|N(N(N(N|)|N)|N(N|N|N(N|)N))))|N)|N)|N(N(N(N|N(N(N|N(N(N(N|N)|N(N(N|)N(N|N(N(N|N)|N(N(N|N(N|)N(N(N|N)|N))|N)))|N(N(N(N|N(N(N|)(N|N)|N))|N(N|N(N|N(N|)N(N|N))))|N)))|N))|N(N|N(N|N(N|N(N|)N(N(N|N)|N))))))|N)|N))|N))|N))|N(N(N(N|)N|N(N|)N(N|)N(N|)N)|N)))|N)|N))))))))|N))|N)))|N(N|N(N(N(N|)N(N(N(N|N(N(N|N)|N))|N)|N(N|)N(N(N|N)|N))|N)|N(N|)N)))|N(N(N(N|N)|N(N|N))|N(N|N))))|N)|N))|N)|N)|N)|N))|N)))|N(N|)N)|N)|N)|N))|N)|N(N|N(N(N|)N|N(N|)N)))))|N)))|N)|N)|N)|N)))|N)$`

const INPUT = `^WWWWSWNNWWNWWNWSSWSSSSWNWWWNWNWSWWNWSWWSSSESWWSWSSSSENNESSSSENNNESSSSWWSSSWNWWWNENESENNWWWNE(NWNENWWSWWSESE(N|SS(E|SWSWSWSEENE(SENESSWSWSEENENN(EEENNNEENNNESEENENESEEESENESENENESESSWSWSESSSWNWSWWWWWWSWSESSSWNNWSWSWNNENNE(NWWSWS(WNN(NE(NWNEENNESENESSWWSES(WWNSEE|)E(S|NENNNN(EESWSEEESWSWS(EEEN(E(E|NNN(ENEWSW|)W(WW|SS))|W)|W(NNEWSS|)W)|W))|S)|WWSW(N|SESSSWWWSEEESSESSENEENESSWSEESWSEEENEESWSWWWWWNWSSWSSWSWNNENWWWSSE(SSSWSESWSEEEENEESSW(N|WWSWSWWWWSSWSSSSENNNEENN(WSNE|)ESSSSEES(WWWWNNES(NWSSEEWWNNES|)|EENENENNENNNNNW(NEEEESSSENNNEESSW(SSWWSSESWSWWN(ENWNENNW(NENWESWS|)S|WSSWS(EENEES(ENESENENNNWW(SESWWEENWN|)N(ENNESSESSEEEESWWSWS(EEEN(W|ENNEESSW(N|S(EEENWNENNW(S|NNWNEENNENEESSESENNWNEEEEEENENWWNENNESENNENNNEENESSESSSWNNWNWSSSWSSSW(NWES|)SSSW(WSEESWSSENENNESESENNENENESESSSWWW(NENESNWSWS|)SSWWN(W(SW(SWSSWNWWSS(WNNNNWNNEES(SSEEN(E|WNNNWNNWWSW(S(EENSWW|)SWSSE(N|SESSWS(E|WNWWWS(WNNENNNNN(ESSE(N|SSW(N|SEE(N|E)))|W(NENSWS|)SWSESS)|EE)))|N))|W)|E(EEEENESENNEEESWS(WNSE|)EEENNNW(WNN(WSSWWWSW(WSNE|)N|ESENENNEESSENNNWNENENENWWWWSWNWNENNNENNENNNNWSWWNENEENEESSSSEENEEESESESWWNWNWSWWWSSENEESWSSWSS(ESESEEESSSESWSEENESSWWSWSEESWSWS(EEEN(W|NEESWSEEEENEES(W|EENNW(WWWNWNENWNENNWSWWSESWW(SS(WNWESE|)ENESSSEN(SWNNNWESSSEN|)|NNNEENEENWWNNESEEEESSEEESWSWSEENENNNNWSWWNENWNWNWWNWWWNWWNNEEENWWNNESENNEENWWWSWNNWWWSESE(SSSWSSW(SESSESE(N(NWES|)EES(W|E(ES(E(E|N)|W)|N))|SWWWW(NWSWNW(S|NENWN(WSS(S|W)|N(N|ESENESSS(WNSE|)E(SEWN|)N)))|SSES(SENENE(NWW(S|W)|S)|W)))|NNNN(ESNW|)NWW(SESNWN|)WWWWNNNNWS(SSSS|WNWSWWSWWSWSWNNENNWNNENESS(W|EENEEENNEENWNNWSSWNWWNENWNNNESENNWNWSWNNENWWWWSESWWNNWWNNWWWNNNNWSSSWNWWWWSSESSSESESWSESWWNWSSSSSENNEEESSWW(NEWS|)SSSWSWSSWSWNWWSSE(N|SESSENNN(W|ESSESSSWSEESWWSSSE(ENN(WSNE|)ENE(S|ENNWNEEENEEESSESSSSSESWWNWW(NN|WSSENESSE(SWWS(S|WWNNE(E|S)|E)|EN(W|ESE(NNNENNWNW(SSEWNN|)NNNWN(EESEES(ENSW|)WWSESEN(SWNWNEWSESEN|)|NNNN(E|WNNE(NNNWSWS(WNWSSWWSWNNNW(SSW(N|SW(N|S(W|EE(ESEEN(NESSEE(SWWWWSE(EEEENN|SWWNNWNWS(WNSE|)SESWSES(SEWN|)W)|NNNWSS)|W)|N))))|NNESENENWNENWNW(WWW|NNEEESS(WNWESE|)ENNEEESESWSWNNWSSSSS(SWW(WSSWNSENNE|)NENWNNES(NWSSESNWNNES|)|ENNESSENNEENW(NENE(SSSSSWS(W(NNE|WS)|ES(S|EEN(NWS|ES)))|NWN(NN|WW(SESNWN|)WWNWSWNWSWWNNW(SS|WNEEEES(WSNE|)E(EE|NNWWWNWWW(N(W|E)|S(S|EE))))))|W))))|E)|S)))|ES(EE(ENEWSW|)SSWNWSSS(NNNESEWNWSSS|)|W))))))|SSWS(EE|WWNWWS(WNNNENNWWNWSSE(SWSESSWNWSS(E|WWWSWWWWNNWSW(NNNWSSWWNWNENE(SS|NENNESESWS(ESENEESEEENWWNWWWNNESEEESEENE(SSSW(N|SWWWWW(N|WW|SSEEN(W|EEE)))|NWW(WWNEENNNWSWNNNENNWNWNWNWWWWNNESEENN(ESSESESEEESW(SESWSES(EEENNE(NWNENENEENE(SSWWEENN|)NE(S|N(ESNW|)WWWWNEEENWWWNWWNNNWNNWWWNENNWWS(E|WS(SSE(SSSESWSSEESSSWW(NN(ESNW|)WW(SESNWN|)NNE(S|NWNEN(N|W))|SEEEEE(ENEN(E(S|E)|WNWSSWWNENWNEN(ESENSWNW|)NNWW(SESWWEENWN|)NWNEE(NWW|SE))|SWSSSW(SEWN|)NNN))|NN)|WNWNEENWWNNEES(W|EES(ENNEESWSEESESWS(WWNNES|SES(W|ENN(W|NEESSW(N|SSEE(SWSS(ENSW|)WNNWW(SE|NE)|N(EN(ESNW|)WNENWN(E|WNNE(S|NWNNWWNNNWNWSSSE(SWWNN(NNWNNEEES(ENEEENNWNNESESENESSENNEEEEESSENENNESSSESENNNNW(NWNWWNNNNWSSSSSE(E|S(WWWNWNWWWNENENWNWNNWNEESSENEEENWNENNENESSWSEESWSW(S(WWWS(SENESSWS(W(N|W)|SE(S|N))|W)|EENESENEEEESSWSEENNENNNWNNWNNWNNNESSENENEESESESSWNWSSSEEN(W|ESSSWSWWW(SSSSEEENNESSSWSSWNNWWSSE(N|SWSEESWSWWN(NNNWSWWSEESSWNWSW(NNNNENWW(NEEN(EESWS(EENNNN|S)|WWWNEN(E(S|N(W|E))|WWSW(N|SSE(ESNW|)N)))|S)|SSWSESWWSSWWWNWWWS(WWWNWNNNN(EESSEES(WSWNWNN(SSESENSWNWNN|)|ENESEESENNWWNWN(WNNWSSS(WNNNWS|E)|EE(SE(ESNW|)N|N)))|WNWSSESWWNWSW(SSSENESEE(NNW(WW|S)|SE(N|S(EENW|WWNWS)))|W|N))|EESWSSES(ENN(EN(EESSW(N|SW(N|SES(EENN(WSNE|)EEEEENWNNEEENWWNNNNESENESESWWS(WNSE|)ESSSWSSSSENESENNESEEESWWWSWSEESWSESENEEEENWNENENEESWSESENNEEENWWWNNNNENNNWWWNNWWSESSWSSEEN(W|NES(ENSW|)SSSWWWS(EEE|SWW(NENN(NWNNNNE(NNNNWSSSWSWSSSSWWWS(EEEE(ESWWWEEENW|)NNNN|W(SWNSEN|)NNNNEENE(SSWWSEE(WWNEENSWWSEE|)|NNNNWSSW(WNENW(NNN(W|EENWNEESSSENNNNESENNNEENWWNWSSSWNWNW(SSSENSWNNN|)NW(NEEE(S(S|W)|NWWNENNW(NWNENWNENWW(S|NNWNWSS(ESNW|)WWNENWWWWNNNNWNENWWNEENNEEESESWWSW(NNEWSS|)SSSEEENNN(WSSWNSENNE|)EENESSSEENEESWSSSSSESSW(SSEEN(NENEENWNNEEEEEEENESESSESEEEEE(SSSSSWSWWWWWSSENEESWSSESSEENE(SSSWSESSWSWSWNNWSSWSEESENEE(SSW(SESSSSWNNNWWSSSSWSWNNNWWSESSSEEENEEN(WNNNSSSE|)ESSSSWSWSEE(SWSSSE(NN|SWSWSSENE(SSSWWN(E|WSWNNWNENNE(NNWSWNWSWSSE(ENWESW|)SWWNNWNWN(WSWW(WSSWWSSWNWNEN(NEESWENWWS|)WWWSWWSW(NWNENWNENNE(NWWWNWW(N(EENNNW|WWS(WNSE|)E)|SES(SE(NESNWS|)SSSSS(W(NWES|)S|E)|W))|SESS(W(SEWN|)N|E))|SES(SEESWSSSW(SSEESES(WWWNEWSEEE|)ENNNEEES(EEENEESEESEENWNNNWSSWNWWNWNNWNWWSSE(N|ESS(E|W(SEWN|)NWSWWWWNEEENNWWNW(SS(EE|WSSS(ESNW|)W)|NN(WW(NEEN|SE)|ESESENEENN(WSNE|)EENWNEESE(ESSWW(NEWS|)SS(WNNWSNESSE|)ESWSESENEEENWWNN(NESSEN(N|ESSSSESENEEE(NNWWWWSEEE(WWWNEEWWSEEE|)|SSSWNNWSSSESE(N|SSSWNWN(NWWW(NNN(ESE(SWEN|)N|WN(N|E))|SSWNWN(E|WW(SSSESEN(NWNSES|)EEEE(NWNEWSES|)ESESE(N|SWSWSSWSWW(SEEESSW(SEENNE(NNNW(NENSWS|)S(W|S)|SSSWWSEESSSWNWN(E|W(NN|SSSWSWNNN(WSS(WNNWWW(NNNE(NNWSNESS|)SSENNE|SSEENW)|SSEEENESS(WWWW|EENNNW(WNSE|)SS))|E(N|S)))))|N)|NENNE(S|NN(ESNW|)WN(E|WWSES(E|SWNWSS(E(E|S(W|S))|WNW(WNNE(NWW(NNESNWSS|)SSSWNWW(EESENNSSWNWW|)|E(SWEN|)E)|S)))))))|NWW(WW|N))))|E))))|W(SS|W))|N)))))|WW)|WW(NEN(W|E(S|N))|WWWSWWWSE(WNEEENSWWWSE|)))|W))|NN(ESNW|)N)|EE(NWNEEEES(WWSNEE|)ENNE(NNN(WSWNWS(SE(SWEN|)E|WWNW(S|NENE(SS|NWNNNWWWNEEENNWWNNNWNNEENENESSWSW(SESS(WNSE|)ESEN(NWNNE(NNE(NNW(NNWNWSSESWWSWNNENWNENWNENENWWSWSWWWWWW(NNENEESS(W(N|W)|EENENNNWWNENWWW(NEEENWNN(EES(EESWWSESWSEESWSSEEE(NNW(SWEN|)NNWNENNN(ESSESSW(N|SEENEEESE(NNWWWW|SWWS(WN(NE|WWSE)|E)))|WW)|SSSS(WNSE|)EE(NWES|)S(W|SE(N|S(ESE(NEWS|)SS(E|WWN(E|N)|S)|W))))|W)|W)|SW(SWWWSEE(SWSSNNEN|)EENE(N|SEESWS)|N)))|SEEEES(ENESSSSSWWSWW(SEEE(NEEWWS|)SSWSW(NWSWNN(EEE|W(NEWS|)SSSW(SESWSNENWN|)NWWWNE)|SEENESSS(WWN(WSNE|)E|EE))|NNE(ENNESS|S))|WW))|S)|SS)|S)|EESS(EENWESWW|)W(SWW(NEWS|)SEEE(SS|N)|N))|W))))|E(NESNWS|)S)|SSE(N|SSS))|S))|SSS))|N))|N)|N)|N(N|W))|NNWNN(E(N|S)|WSS(S(S|E)|W)))|NWWNENWWSSWNWNNWNEES(EEENWWNNESEE(SSSS|NNW(S|WWWWNNESEEENWWNEEE(NWNNNE(SS|NNWWNNNESE(SWEN|)NNNNWSSWWWWSWWNENENWN(EESEE(SWWEEN|)N(W|EE)|WSSWWSESWWWSSEEESSSSWNWWSSE(ESWSWSEENESENNENW(WSNE|)NNNNNN(EEEE(NWES|)SWSESWSESSWS(EENNNNE(WSSSSWENNNNE|)|W(NN(E|NW(S|NNE(NWES|)S))|S(WSSSSEESWWWWNENWWSSS(EE|WNNNWSWWNNNNE(N(E|WWSSWNNNEENWNENE(N(ESEENW|WWNNNEE(SWSEWNEN|)NWNWNN(WWSESS(SSWSESS(E|SWNNWNWWWWNWSWSSENEESSSEEE(NNWNWSSE(WNNESEWNWSSE|)|E(E|SWSWSSWWW(SS(S|EEN(W|EENE(SEESSSWWW(NEENWWW(EEESWWEENWWW|)|SES(ENESEENN(WS|ESSEE)|WSS(WNNNSSSE|)E(SWSNEN|)N))|N(N|W))))|NNN(ESE(SW|NE)|NWWN(EENSWW|)WNWWNENWNNWSSSWWNENNWWNNESENEENN(WWWWS(EEE|WWWSSSWNWNWSSWNWWSSE(ESWWWNNNNENEE(SWEN|)EEE(S|N(EEE|WWWWWWSWNWSSWSEE(SSWSESSWSSSSESWSWWSSSESSWNWSSEEESEEESSWWWN(EE|WSSWSWWWWWSESSSEENWNEESSSSSSWWWWNNWWSSE(SSSESWWNWNN(ESNW|)NWNNNWNNNNNWSSSWNNWNWSSSWNNWNWSSSSSSESESWSESESWSEEESENEE(SWSSWNWWNWSSSWNNNWWNENWN(WWNENN(ESSNNW|)WSWSSWNNNWSWSSSSSENN(EEEESSSESE(N|SWSWSEESWS(EES(ENEENNWSWNN(W(W|SS)|EEEEEN(WWWNWS|ENEENN(ESSSSENE(NNEEEE(SWS(WNWSNESE|)ES(W|EEN(ESEWNW|)W)|NENWNNWSSWWWS(EEE|W(SS|NNNN(WSNE|)ESSENNE(ENNNEES(ENENWWWWNENNN(W(SS|WWW)|NESEESSS(ESE(SWSSWN|NENNWSWNNNEENNNW(SSWWW(SEWN|)W|NNENWWSWWNENNENESES(ENENWNNWWWNNW(NNW(NENEN(EEEENN(WSNE|)(N|ESSEENN(WSNE|)E(E(E|SWSSE(SWWNWSSWWWNEN(ESNW|)WWWSESWW(N|SEEEEESENENEEENEE(NWWEES|)EESESSEE(NWES|)SSWNW(SS|WNWN(WW(NEEWWS|)S(WWW(NEEWWS|)SS(ENSW|)WSSWNNN(ENSW|)WSW(SESSEESSSWSWSWW(N(NE(S|N(W|NES(ENSW|)S))|W)|SESSW(N|S(WW|EENNNESENN(NENE(NWNNSSES|)SSSES(ENESSNNWSW|)WSSSWSESESSESS(ESNW|)WNWWWNEENWWWNNN(E(SSEWNN|)NWNEE(S|NNN)|WWWSEESSSESSEEE(WWWNNWESSEEE|))|W))))|NWN(WSWNSENE|)E)|E)|E)))|N))|NWNNES))|W)|S)|SSS(EEESNWWW|)W(WS(WNSE|)E(SWS(S|E)|E)|N))|WW)))|WNW(NEWS|)S))|W)|SS))))|SSSS(WWNWN(EESNWW|)WW(NENESNWSWS|)SES(E|WWSWNW(SSW(NWSWENES|)SEENES|NEN(ESNW|)W))|E))|WSW(WSNE|)NN)))|W)|WNWWNWNENNWWWN(EENESE(NNWWEESS|)SESWSS|WWNWNWSWSSWSWNWNWNNWNNENWWWNNWSWSSSSWWWNNNENNWWS(WNWNNWWWSSWWWNNNNWWSWSESWWSEEE(NNN|SEESWSWSESSSSSWSSWNNWNENNE(SS|NWNNWWNWSSW(SSESSW(N|SEE(SSE(N|SSSWNWN(W(NENWESWS|)SSESESWW(SESSW(SSSESWSSENENESSWSEESWSESEESSWSWNN(WSWWSESSWNW(SSEESENNN(ESSESESEESWSWSSWWSSWNWWW(NNNENEEN(E(SSS(ENN(N|E)|WW(WSEE|NE))|N)|WWNWSS)|SSENESEESENNNESSSESWSSESWWSSESWWSWWNNNWSW(SESWSSESESWSW(NN|SSSSENNEN(EEEEENNNNWNN(WSWSESSWW(N(NWN(W|E)|E)|S(EEENNSSWWW|)W)|ENNEESSW(S(EENENNNNWS(WNNW(SSWWEENN|)N(W|EEENWWWNNNW(S|NENN(WSNE|)EN(NNNNEENWNEESESSESEEESENNN(WWS(WNNW(S|NWNWNEENWNWNNWWNEEESES(W|EEE(SE(N|ESWSWWN(W(SS(W|E(S(S|W)|E))|N(E|W))|E))|NWWNENWNWS(S|WWWWNWNNWWSSE(N|SSSWWNN(ESNW|)NWS(WNNENNNENNWWS(E|SWNNN(WSSSS(WNSE|)ES(ENSW|)W|EEENESEESENNEEESSWSEESWWSESEES(WWWS(E|WNN(W(NNN(WWW(NEWS|)SS(WSS|ENES)|E(NNEWSS|)SS)|S)|E))|ENEESSE(EEENWNNNNWWNNNENNWSWSWNNNWNWWWNWWNWSWNWW(NNNESENNNWW(SEWN|)NEEENNESESWSSW(SSS(ENNESSES(ENNNN(W(SS|WNEENWNN(ESNW|)WNWNNWSSWSSW(W|NN))|ESEEEEEESSWSWSSS(ENEEESENENN(WSWWW(W|NEEN(E|NW(NNWWNNN(ESSNNW|)W(N|SSSWWW)|S)))|ESSEEESWSSS(E(S|ENEENNEN(ESSSW(N|SW(SEENE(SSWWEENN|)NENNWS|W))|NNWN(WSWWS(EES(ENSW|)SSW(WSNE|)NN|WWNEN(W|NNESE(SWEN|)E))|EE)))|WNNNWSW(N|WWWNWSSEEESENESSWWSWNNWSS(WWNENSWSEE|)SSES(ENNW|WSE))))|W(NNW(S|NWW(NNEES(W|E(S|EN(W|E)))|S(W|E)))|S)))|W)|WW)|N)|SW(N|SEE(N|SESWWNWSS(WW(S(WNSE|)E|NENNWS)|EESENENNN(W|ESSSEENWN(EESSSENE(NWNSES|)SSWWSEESWSS(ENEE(NNNWSS|S(ENESSWSE(WNENNWESSWSE|)|W))|W)|N))))))|SWW(N(N|W)|S)))))|SSSEEEE(NNN|SWWSSESWSW(N|WSESENE(SSWENN|)EN(NNESENN(E|NW(NE|SWW))|W)))))))))|E)|E(SSSSWSESSWSESSEESWWWWSESSSSSESWWSWWSSSEEESEESWSSS(ENNEEN(ESENEWSWNW|)WNNE(NWWW(SEWN|)WN(WSWNSENE|)NESEENE(S|N(ESNW|)NNNNWSSW(SESWSNENWN|)WNNNNEEES(WWSSNNEE|)ENNNEN(ESSESSWS(W(W|NN(N|E))|EENNNE(SEWN|)N)|NNWSWS(E|WNN(E|W(NENWNNESEEENW(N|W)|SS))|SS)))|S)|SSSSWNNNWWWSSEE(NWES|)SS(EE|WNWWS(WWWNWSWWNNWSWWS(WWWNWSWNNNNESES(W|ENNNN(WSSNNE|)NESESSW(SS(SWEN|)ENE(S|ENEES(W|ES(WSNE|)ENNESSSENNE(NNNENNN(WWWWNENNN(WSWWSW(SEE(NEWS|)SWWSEEESWWWSEES(WWWNNW(SSS|W)|S|EEN(W|NNES(ENSW|)SSS))|N)|ENE(EENWWWW(S|NNNESES(W|EENNE(NWNWNNE(EENWWNENN(WNENWNWWWNWSWW(SSSSEEESEE(N(E|WNENWWWS(WNNEEEE(WWWWSSNNEEEE|)|E))|SWWSSE(SWS(EES(W|S)|W(NNNNNWSWS(E|WNN(NN|E))|WW))|N))|NEN(ENSW|)W)|E)|S)|SSS)))|SSWSESW(ENWNENSWSESW|)))|ESSSS(EENNW(N(E|N)|S)|W))|S)))|N))|EE)|E)))|E))|W)))|SS)|W)|N))|W))|NNEENENWNN(ESEESSW(SSSW(NN|S)|N)|WSW(SSENSWNN|)N)))|N)|NNNNESEN(NNWW(SE|NE)|E))|E)|N)|N)|E))|NNNE(NW(NEWS|)W|S)))|NNNE(NNNW(SS|NENE(ENWWW(S|NNNEESEEENNNEEEEESWWWWSSSS(WWN(WWWNSEEE|)E|ENNESENN(ESSENNNNEESSESESSSSSS(EES(ENNWNW(S|NEESEESS(WNSE|)ESENNEESEENNEESSW(N|SEENEES(ENENNENNNNWSWNNWSSSES(ENSW|)SWS(WWNEN(E|NNNNNNWNEEESEESW(WWNSEE|)SEEESSE(NNENEEENWNWWS(WSWWNENENENENNNWSWNNWNNEES(SEENNW(NNNWWSWWWNENNES(EEENESENNENWNNWSSSWWNNE(NWWSWNWWSWSWNNNWWNWSSESSE(NN|SSEES(ENN(NE(ES(W|E(SEWN|)N)|N)|WW)|WWWNNWNNWSWNWNENNE(SS|NWN(EEEEEESS(WNW(WW|S)|ENNEESWSS(WWSNEE|)EEENW(NNEEESEEESENN(WWW|ESSSESWSWNWN(E|W(WN(WW(SE|NW)|E)|SSESESWSWSESWSWS(WNNEWSSE|)SESESSWW(N(N|E)|SESSWS(W(W|NN(E|NN))|SENESSWSES(EEESSENESE(SESWWNWSS(WNNSSE|)S(EEN(W|ESESEE(NW|SWW))|SS)|NNENNEESS(WNSE|)ENE(S|NE(S|N(WW(S|NWNNWSWWWSWNNNEEE(SWWEEN|)NNWWNENWWWNNNNNESSSSEENNNW(SS|NNESEENNNWWWNNW(NWNEN(ESES(W|ESS(WNSE|)EESESSEESE(N(E|NNWW(SEWN|)NE(EE|NNNWWN(EEESNWWW|)WSSS(WNNWN(E|W)|E(NESNWS|)S)))|SSSSWS(E|WSWSSE(N|SW(SS|WWNNNNNES(ENENWWNENN(ESES(W|S)|W(SWWSS(ENSW|)SSS|NN))|SSS))))))|W)|SSS(EEESNWWW|)W(S(W(SSW(N|SSE(SSSSEE(NWES|)S(WWW(SWSEEE(NWES|)SSSSWW(NENNW(W|S)|SEEEE(SWEN|)NE(S|N(WWSNEE|)NEES(W|E)))|NNNWN(N|E|W))|EE)|N))|N)|E)|N))))|E))))|WWW(S|W)|S))))))|W))|WWWWWWWSESSSWSESSSWNWN(E|NWSWNNWNENNN(WSWWN(E|WWSESWWWNN(ESNW|)WSWWN(WWSSE(N|SSWWSESENESSSENNNNWNNEES(W|ESSW(SSENENNESSSWSSEEN(NESSSSESSWWWWNENE(SEWN|)NWWWNW(NEWS|)WWSWNWWNN(EESWENWW|)WN(E|NNWSWSESWSSENESSWWW(NNNNNNNE(NENEN(WWWS(E|S)|ESSS(E|W(N|W)))|S)|SSENESESENESESWWSEEENNNESENN(W(NEWS|)WW(S|WW(NEWS|)W)|E(SSSWWSSSWSWSWNWWWWNENNN(WSW(NNEEWWSS|)S(E|SSSSSESS(WNSE|)ENNNEN(WWSNEE|)ESES(WSSSWNNN(SSSENNSSWNNN|)|EEN(EN(NEEEESSS(WWNENWWSSW(ENNEESNWWSSW|)|EENWNENWNEEENNEEENESSSENNNENWWNEENESSSEEENWNNNNNWSS(SSS|WNNNW(SSSWSWNWN(EE|NWSWWNN(NW(SSSSEEESSE(SWSWW(WWWWSS(ENESEN|SWNWW(SE|NEEN))|NENWNE)|N)|N(E|NWW(SSENSWNN|)N(EEN(E|W)|WW)))|E(S|E)))|NNNNEEE(SWWSES(W|E(S(E(SENESSSWNWSSSEEN(EN(EESWSSEESWWWWSW(NNEENSWWSS|)SEESSSS(ENNEEEENWNN(NEE(SS(WNSE|)SE(N|SS(ENEWSW|)SWWW(SEEWWN|)N(WSWN|EEN))|NN(EEENSWWW|)WSWNW(NENWESWS|)W)|WSWW(SEEWWN|)N(E|W))|WSWSWSEE(N|ESSSSWWS(WWS(WNNWWWNNESEEENN(NWSSWNNWSWNNNESEEENWWNNWSWWWNWNW(SSWNNWSSSESSESW(WNNWSNESSE|)SSSENEES(SE(EES(WS|EN)|NNNWNENNNWW(S(WNSE|)ESWSSW|N))|W)|NEE(NWWEES|)SESENENENEESSSW(N(W|N)|SEE(NE(S|NWN(EESSNNWW|)NNWWWWS(NEEEESNWWWWS|))|S)))|ESEE(SWWSW|NW))|S)|EE)))|NN)|W)|N)|W)|N))|NE(S|NWWSWNWS(NESENEWSWNWS|))))))|W)|W)))|ESSSEEEN(E|WW))|E))))|W)|N)))|E))|ESSE(ESWSWN|NN)))))))|S)|S)|S)|W)|E)|SWSW(NNN|SEE(SWS(W(NWSNES|)S|E)|N))))|E)|WSSWWSWNWSWNW(SS(EES(WS(ESW|WN)|ENEEEN(E(NN|SS(S|W))|W))|W)|NN(ESE(EEE|N)|W(SWWEEN|)N)))))|SSSSWN)|WNNW(NWSWNN(WWWSS(EENWESWW|)W(SSENSWNN|)N|EEN(ESSNNW|)WN(E|NN))|SS))|WW)))|SS))|EEES(W|E)))))|E))))|NN)|E)|NWWNEN(ESNW|)NNWWNE(NWWWS(ESSS(WNNSSE|)SSENNENW(ESWSSWENNENW|)|WNN(W(NN|S)|EE(EE|NN)))|E))|N))|N)))|N))|EESSEESS(WNWSW(N(NN|WW)|SE(SS|E))|ENNNN(WSWNSENE|)ESENEESWSWSSS(WNNNSSSE|)EENWNENESS(NNWSWSNENESS|))))))))|E)|ESENEEEESWWWS(ESWENW|)W))|SSS))|SS(EEEEE|S)))|E)))|WWW)|N)))|SS)))|S))|W)|NWNW(NWNWWN(NE(NWNEWSES|)SENEESS(WNSE|)S|WSWNWSWN(SENESEWNWSWN|))|S)))|S))|WWWSESW(ENWNEEWWSESW|)))|WSWWWWWSW(SSSWSEE(SSSWWWNN(ESENSWNW|)NN|NE(NWNENSWSES|)S)|N))|S(E|S))))|S(E|SS))|EEE)|S(SWW(SEEWWN|)N(E|W)|E))))|W)))|W)|W)|W)))|E))|NEEN(WWNWNENWWNEEN(E|N|W)|E))))|N)|S))|SS)|WW)|WSSWWWWWS(WWWSESSSSWWNWWWSW(NWSWWWWS(E(E|S)|WNNNESEENNW(NNEEEN(WWNWSWSSWS(W(SW(SEWN|)W(NEWS|)W|NNNWNEE(E|SS))|E)|ESEEENE(SSWSS(EE(NWES|)S|WWNN(ESNW|)WSSWW(S|NN(ESNW|)W))|NESENE(EEENESEN(SWNWSWENESEN|)|S|NWWN(NWSWS(E|S)|EE))))|S))|SEENESEE(WWNWSWENESEE|))|EEE(SWEN|)EEEESEEE(NWWEES|)SSSWNNW(W|S)))|N)))|W))))))|W)))))|SSSESSEEE(N(ENSW|)WWNNWN|SSWW(WWN(NNWNWWWSSENESESWWW(EEENWNSESWWW|)|EEE)|SEESSS(S|E|WW(NENWESWS|)W))))|WW)|W)|W(WWWWWSESE(SWW(N|SSSWSWW(SES(W|S(EENWNENENNNESSENNESSSWWSEEESWWSW(NW(S|N)|SSS(WSS(ENSW|)SWWS(E|WWSSS(ENE(SE|NW)|WNWWS(WWNEN(WWSSSE|NEEE(SWWEEN|)NWWNEEEN(ES(ENES|SW)|WW))|E)))|ENEN(NEEENE(SSS(EE|S(WNWSW(NNEEWWSS|)(W|S)|S))|NWW(S|NNW(S|NE(NWES|)ESSE(E|N))))|W)))|S))|NEN(E|W(N|WSWWN(E|WW(SESEWNWN|)NEENWN(W(NN|S)|E))))))|N)|S))|S))|W))|SES(EESESSS(WNNW(NWSNES|)SS|ENESEN(ESNW|)NNN(NESENSWNWS|)WW(SESWENWN|)W)|W)))|E)|E))))))))|N))|S)))|WNWWNWSS(EE|WNNNNENESEE(N(NNWWWWS(EEESNWWW|)SW(SSWWN(W(NN|SSSES(EENNW(W|S)|WWNNNN))|E)|NNNEEEENWN(EESSNNWW|)W(WS(W|E)|N))|E)|SS(WNWWEESE|)E)))|WNWNW(NE(N(NNNNNN|W)|E(S|E))|WSES(E|W))))|SS)|N))|N)|N)|E)|N))|W)))|W(NNEWSS|)W)|W)|W)|W))|N)|SSSWWN(E|WSWSSSEE(ENN(WWSEWNEE|)E|SW(WWNNNSSSEE|)S)))))|N)))|E)|S)|N)|N)))|EE)$`

const INPUT_TAEL = `^EESWSWWSSSENNEESWSEEENNW(S|N(ENNNW(NENNESSENENNESESENESENNNESENNNNNWSSWNWSWNNNNNWSSSWSWWNENENNNEEENNWWNENWNWNWSSWSSSWSSSS(ENNNE(NNE(SEEWWN|)N(W|N)|SS)|WSSWWNENWWNNENNWWNNNEEENENES(ENNNWNEENEENEEESSENNENNWSWWWWNWNWSWSWSWWNWWWWNNNENEENEEENWNENEEENWNEEENWNENESSENNNWNNNWWWNNENEESWS(EENNNWWWWNNESEEENESSSEESENNWNNESEESESSWNW(SSSSESWWSWSWNW(SSSEEEN(WW|EN(W|NESSSSSSWWSEEENENESEEENNWSWNWWW(S|NENNNNENWW(NEN(WWSNEE|)NEESEEESWWW(WNSE|)SEESENESSSEEEESSWWSEEESWSEENNNWNENNENWNNEES(EEEENNWWNNWNNWSSWNNWWSESWWNWWNNWSWWWWWWNNNEEEEES(EEEEEN(WWWW|EESWS(WWSWNSENEE|)EEEESENEENWN(EEEEESWWWSEEEEESSWNWWWSSENESSWWSSENESENNNEEESEENENE(SSWSESSSWNNWNWWWW(NEWS|)SSWSSSESEENNW(N(WSNE|)ENN(WSNE|)ESSSSEN(NN|EESWSWWSWWSWSSENESSSSSEENNEE(SSSW(SESWSESWWNWSSEEESSSSWSESWSWNWNENWNWNEESENNWWWWNNWSWWWSSWWWSSSSSWNWNWNNENNWSWSSSWSWWWNWSWWWSWNNWNENNESSENENENWNNWWNENESEEEESSWSSW(NNN(WSNE|)E|SWS(WSWWEENE|)EEE(SWEN|)N(W|E(NN(WSNE|)NE(ENENNNESSSENENNNNNEENENWNWSSWNNNEENWNNESESE(SS(WNSE|)SSEE(NWES|)SSSWNNWSW(WWSESEE(NWES|)EEEEE(NWWEES|)SWWWWS(EESSNNWW|)WSWNWSW(NN(N|EEE)|SSWN(N|WSW(NWES|)SSS(S|W)))|N)|N(NWNNNNWWSWW(SWWSESWSWNWWW(NENEN(E(SSWENN|)EE|W)|SEESESWWNWSSSEESENESEE(SSWWWN(EE|WSSSE(NEEESS|SSWW(SEWN|)N(NNWNNWSSSS(ENSW|)WNNNNWSSS(WWWSWNWWSWNWWNEEEENWWNENEES(SSEENE(NWW(NNWW(WWSWNWSW(N|WWWSWSSSEESWSESSEENNW(NE(NNW(S|NW(WSEWNE|)NEEEESWW(EENWWWEEESWW|))|ESSEN(EESSW(SWW(NEWS|)WWWWSWSEENESESSSENNESSENE(SSSSSENN(ESSSSSWWSSSEENEEEESWWSESWWSWSSEEN(W|NESESWSWSSESWWWSSSWWSSWSEENEN(W|ENNESENN(EEEENENNNWNENNEESSW(SEEENWNENESEEESSWWSESWWWWSWSSWN(WSWSEEESSWSWSWSWWWSSEEEESWSESWSWWNENNWSWWN(E|WWSSSESWSEESWSWWN(WSSSESENESEESWSEEENNESENENEESWSEESSESWWNWWWW(NEE(EE|N)|WWSSESWWNWWSSE(SENESEESWWWSSWWWNEN(ESNW|)WN(NWSWNNEENNEEN(ESNW|)WWWSWNWNNN(WNNE(NWNNNEENE(SSWWSNEENN|)NWNNWNNESE(NNN(WSWNWWSWNWSSSSWNNWNWWNWWWWNENNNENNENNWNEEESSEEENNENWNNEES(SEESSSE(SSWNWSSWNWWWN(WSSESWWNWN(ENWNNSSESW|)WSSESEESS(E(NEN(W|EE(SWEN|)N(WWNEWSEE|)E)|S)|WNWWWN(N|E))|EENE(S|ENNWSW))|NNN(EESSWN|NWNNW(WNWWNWSWSE(EE|SSSW(SESWENWN|)NNWNWSSS(ENSW|)WNN(WWSS(ENSW|)SWSWSSWNWWWNN(EEN(W(N|W)|EE(SWSWWEENEN|)E)|WSW(SSSSSENNEN(WNSE|)EESSSSENNNEE(SSSW(NN|SWWSWWNN(WWSWSESSWNWNWNWN(EE(S|E)|WSSSWNWSSEEEN(N|ESESSWNWWWWSESSWNWWNNE(NWNENWWNNWNNNNEENWWWNEEEENWWNWWWWWNNNNNENWNNEENNNENWWSSSWWWNWNEESENNNNWWNNESEEESWSEENNNENWNNWWWSWS(WNWNNWSSSWNWSSWSWWSSWWWWSSESWWWNWWNEENE(SS|NWWSWWSWSWSSWWSSESSSSSWWSSSSSSEEENNESSSWSESEENWNEEEESWSESSENENWNNNENNNWSWS(SWWWNNESENNWNENWWWSWS(E(ENSW|)SSS|WWSW(SEWN|)NNEN(E(NNNEENEEESENNWNENESEESENNESESWSSWNWSWW(NENWESWS|)SWWWS(EEEENEESWSS(WNWSW(S|N)|ENESSENEENNNEEEENNNNN(ESSSSSSSSWWWNW(SSWSEENESSSEEESWWWWSWNWWSWNWNWW(SS(WNSE|)SESESEESWWWNW(N|SSWSWWWWNNNE(SSEEN(ENSW|)W|NW(WWSESSWNWNNWSSSSSSEESEESSWNWSWNNWWWNWWSWWNWSWWWWSSSEENEEN(WWWSNEEE|)EESSSESWSSEEN(NENNW(S|WNEEESE(NNWWW(N|W)|ESWSSSWN(NN|WSSWWSSWNWWNWNNESE(SEWN|)NNWWWN(ENEE(SWEN|)N|WWWNWSWWSWSESWWW(SESESESSSWWW(SSSSSENEESSW(N|WSW(SSESSW(SESWSSSESWSSESSS(WNNSSE|)EEEENNNNWSWW(SEESWW|N(W|NE(NNESENNWNEESESSW(S(WW|ESSSSENNNNNNNNWNWWNNENEENNNESSSENNNEEESSSSSWNW(NNNESS|WSESSENEEEENWNW(S|NNEEENWNWNW(SSEWNN|)WNEENWWNNNEENEEENNNESSSSSWSESENESSENNENNNNWNWNNW(NNENNN(WSSWSSSWWSWWSS(EENWESWW|)WSWWSWNWSWNNEN(WN(WSSSSWNWNNWW(NWN(W|ENESES(E(NNEN(ESSWENNW|)WW(WWNEEEE(WWWWSEWNEEEE|)|S)|SS)|W))|SSSSEN(ESEESWSSENEENN(WSNE|)ESSE(SWS(WNWSWWSWNNWNW(WN(EEESNWWW|)WSWNNEE(WWSSENSWNNEE|)|SSESWSEEESWWSWWW(N(E(NWES|)E|W)|SEESWWSSE(NEWS|)SWS(WNSE|)S))|E)|NN)|NN))|E)|EES(E|W))|EEENENWNNW(WWN(EEEESSEENWNENENEEEEESWSSWNNWWSESSEESWSSWSSSENNESENEEESENENENNNEENEE(NNNWNWWNE(E|NWWSWSESWWSWWW(SWSSSSWSEEENWNEESSENNE(S|NWWNN(WSWSEWNENE|)EENEES(EN(NW|ES)|SW(N|WW)))|NNENNN(WWW(N|WW(WNWNSESE|)SEESE(ENW|SW))|ESE(SWS(SWEN|)E|NN(WW|EES(WSNE|)ENE(S|NNWSWNNWWNW(NEESEE(N(W|EE)|S)|S)))))))|SWSWWSSESWSSWW(NENSWS|)WWWNWSWSW(SWSWWWWNENNWSWSSSSWW(SSSENNEENEN(EESSSEENESENNWWW(S|NENEEENEN(WWSWENEE|)EESWSEEESWWWSEEESSSSSESWSWNNNNNWWWW(SSSESENE(SSSEEEEESEENWNNWNW(SS(E|W)|NNESEES(W|SSEESWSSWSSWWSSESSWNWWNENWNNNN(WWSSSSSE(NNNN|SWSEESE(ESWWSWNWN(E|WNNNWWWSWNNEENWNNWNNEEE(NWWWWWWSESE(SWSWNWN(E|WNNNNNN(NW(NENSWS|)SSWNNWSSWWSSSW(SEESSSWSWSESWSWSESSSWNWSSSSEEENEENENEEEESSENEENWNENN(WSWS(WNWWNWSS(WWSWNNENNW(NEENNNESEN(ESEESSE(NNNESNWSSS|)SWWW(SEEWWN|)NENWWSW(S(WSSNNE|)E|N)|NWNNWSW(NNWNEEN(ESSWENNW|)WW|S(SSWSEWNENN|)E))|SW(SSSS(WSWNSENE|)E|N))|EE)|S)|EEE(NWES|)SWSSW(SSEEENW(W|NE(NWES|)ESESENESESEESS(EENWNNNNNEESWSEENEENWWNWNWWW(S(EE|WSESWS(SEN|WNW))|NENWN(ENNEEEEEESWWWWWSESS(WNSE|)ESENESEENWNWWN(WSNE|)EEESEEESSWW(NEWS|)SWW(SESWSESEESSENNNEENENESSENNNNESSSESWSWWSWWSW(NNEENSWWSS|)S(WWWWWWWWWNNESENEN(W(NE|WWWS)|ESES(ENSW|)WW)|EENESEENEN(WWSNEE|)NENESENNWWNEEESSESSEEENWNNNNEEEENENWWWWNENEENNESESESW(SSENENENNESSSENENWNNWWN(NEN(EENESSSEEESEEESESSSSSENEESWSSSEESWWSWNWWWWWSEEEESES(ENENESS(EEENNWSWNNEEESSE(SWEN|)NNNWWNNWSSWNWN(WSNE|)ENENNWNWS(S(S|E)|WWNENWW(NENEENNNENEENWNNNNESESWSESSSSSSE(SWSSW(NW(S|WN(W(W|S)|NENE(SSWENN|)NW(NE|WS)))|SESE(NNN|SSSWNNWN(SESSENSWNNWN|)))|NNNNNNNNNWNENNWWWNENEE(SWEN|)NWNENNNWSWWWNWSWWSSESSSWWNENWWSWWWW(SESSESSSEENWNNNWNEESSEEN(W|ESSSSWNNW(W|SSSSEEEE(SWSWNWSSWSS(W(SEWN|)WNENNENWNNWWWNWSSSESWSEE(SSWNWWWNNNE(NWNENWWN(E|WSSESSWNW(NNNWESSS|)W)|SS)|ENWNN(ESNW|)W)|EEN(W|N))|NNNNN(ESENEWSWNW|)W(NNE(S|NWNNN(W|EEE(E|SSW(NWSNES|)S)))|SSSSWW))))|NNWSWNWNEEEEENWNNNWSWWNENENWNNWNN(ESENEESENESENNWWNENENWNWNN(WWWSWSSEESSE(NENWNNWSW(ENESSEWNNWSW|)|SWWNW(NE|SW))|ESENEESWS(W|SENESEESENNNNESE(NNNWWS(E|WWNENWWWNNWSW(NNWNENENESENNWWNEENWWWNEEENEEENESE(NNN(WSWWWNWWWWNWNEESENENN(ESSS(ESEWNW|)W|NWN(E|WW(NEEWWS|)SWSS(ENE(ESWENW|)N|WSW(NNENWESWSS|)SESWSWWN(E|WSSSSWNWWNNNWSSWNWSSEESSWWWS(EESWSSEEN(W|EESS(W(WWSW(NW(N|S)|SE(ENSW|)SWSWS(W|EE(S(EN|SWNWS)|N)))|N)|ENNNNWW(WSEEWWNE|)NE(NWES|)EEEESSENNESE(SWSWWSWW(NNE(NWES|)S|S(S|EEEE(SWEN|)N(ENENNSSWSW|)W))|NNNWN(WSWW(NE|SEEE)|EEEN(EES(WSWWEENE|)EE|WW(W|N))))))|WNWSW(WWSNEE|)NNNESENE(SEEWWN|)NWNENN(EE(SWEN|)EEEE(NWES|)(E|SWSSEN)|WWS(WSESWENWNE|)E))))))|N)|SWSESSWSSWWW(SEESEE(SWSESWWNWWNW(N(EESEWNWW|)WN(WSNE|)E|S)|N(W|NN))|NEENWNE(E|NNWWS(SS|E))))|SESES(WWWNEWSEEE|)E(SEESW(SEWN|)W|N)))|SWSSE(SSSSWNNWSWSESSWWSSSWNWSWSE(EEESENNWNEEE(SWSEWNEN|)NNNWSSWW(EENNESNWSSWW|)|SWWSWSS(ENSW|)W(SWNWSNESEN|)NNNWNNNEEE(SSW(WNEWSE|)S|NWNEEENE(SSWWSNEENN|)E(NNWNWW(N(NESENESSENEE(WWSWNNSSENEE|)|W)|SS(SWWWSSWW(S|NN(ES|WS|N))|E(N|E)))|E)))|N))))|WSWS(ESSS(ENNSSW|)SSWSWWWW(NW(S|N(WNSE|)E)|SSENEEE(ENEEESWW(EENWWWEEESWW|)|SWSES(SWNWSW(NN(ENSW|)WSW(SEWN|)(WWW|N)|SS)|E)))|WWW))))|SS))|W)|WWNWWWS(WNWWS(E|WWWWWWWNENWWSWWNENNEENWWWN(EENEES(SESS(W(N|WW)|SENNNNWNEEEN(EESWSSENENENWNW(W|NENWNWN(NNWSWENESS|)EESE(EESSEE(NWNSES|)SWWW(NNWSSNNESS|)SESEE(NWES|)SWWW(N|WSES(EEEN(WW|ESENNW)|WW(WWSWW(SS(WW|ENE(S|EEN(ESNW|)W))|NENWNE)|N)))|N))|WW(N|W)))|W)|WWSES(ESWSSS(WWWN(EE|WWS(WWW(W|NEEN(ENSW|)WW)|E))|EEE)|W)))|EE))|W)|WSS(WNNW(WWWSSWNWNENNESENNWWWSSWWWWWNNWNEESENEE(NNWSWNNWSWWNWSS(EEE|WNWNWWWWWWSSENESSSESENNEE(ESE(N|SESSWWWWSWSWWSSENEEEEESESWSW(NNWSWW(NEWS|)WSWW(N(NNNNEENWN(NWWSWNNN(ESENSWNW|)NWNENWNEEEENWNENEEEESEENWNWWWWWNEEENWNEESEES(WW|SESSENESSS(ENESEE(NWNN(ESNW|)WNENWNNWWSESS(SSEWNN|)WWN(E|NNW(SS|NWSWNNENWWSWS(WWWSESWWNNWNEEENNNWWNEENNEEEEN(EESWSESWWSWNN(E|WSSW(NN|SE(SWSNEN|)EEEEN(W|EN(NN|ESEEES(ENN(ESSNNW|)NNWSSW(W|NNNEEN)|WSWWW(NEEWWS|)WSESENESEE(SS|N(W|N)))))))|NWSWNWWNWSS(WNWSSWNNW(SSWSSW(NN|WWSWW(NNESNWSS|)SWWWNWWSW(SEE(N|EEESSSENESEENNNN(NEEEENN(WSNE|)E(ENESNWSW|)SSSWS(EENESNWSWW|)WWSEESSESWSEESSWSWNWWSWNNNWSSSWWSWNWNENNESE(SWEN|)NN(WWWSSWSWSSWNWSWW(NNNNN(WWSEWNEE|)ESEES(ENNWNW(S|WN(WSNE|)NESEEE(NWWEES|)S(SS|W))|SWW(NEWS|)S)|S(EEEES(W|EEEESSENNENN(EE(E(EE|N)|SWSESWWSSEE(NWES|)SSE(NN|SSW(NWNWN(E|WNWW(SESSS(WWNENSWSEE|)EN(ESEWNW|)N|WNW(NEESNWWS|)S))|SEEE(N(ENSW|)W|SWW(SEEWWN|)W))))|WSWWWWN(E|N)))|SS))|EENN(NNNEEWWSSS|)EESSSW(SEEEWWWN|)NN)|WSS(WWNENSWSEE|)S))|W))|NN)|EE))|E)))|S(E(SSSWNN|EE|N)|WW))|W(WWW(N(EEE|WWWW(SEEE|NEE))|S)|S)))|EE(S|ENNESES(E|W)))|E)|SWWSES(WW(SEWN|)N|ENENEEN(SWWSWSNENEEN|)))|S(EEEN(NNNNESSEEE(NN(ESNW|)NWSSWNNN(EE|WSW(WS(S|WWN(WSWENE|)E)|N))|SWWWSSENE(EEEE|SSWSESS(NNWNENSWSESS|)))|W)|S)))|NWNWSW(N|S)))|SSW(WW|N))|S)|S))|W(N|WW)))|WW)|W))|WWN(WSWWNEN(WW(SSWWWNN(EESWENWW|)WSSWWWWNN(WNWW(NEEWWS|)SSWN(N|WWSES(WWWWN(EE|WWWNENNWSWWNWNENNNN(WWWSSSSSSE(ESSS(WNWSWWWWW(NNNNESSEEE(SWWWEEEN|)(NWNNE(S|NNWN(ENSW|)WWW(NENWNW(S|W(NEEWWS|)W)|SEES(S(WWNEWSEE|)SS|E)))|EE)|WWWNNNN(SSSSEEWWNNNN|))|ENNESSEEE(WWWNNWESSEEE|))|NNNENWNE(WSESWSNENWNE|))|EEENWWNW(NN(ENNNEESSW(SEESE(NNWNNN(ESSES(ENSW|)S|WWN(E|WSW(S|NW(NNWN(WSNE|)ENE(SSSEWNNN|)E|S))))|SW(WWN(W|E)|S(E|SSS(E|WWW(NEEWWS|)SS(ENESNWSW|)S))))|N)|W(S|W))|S)))|EEEENWN(SESWWWEEENWN|)))|ESEENWNE(WSESWWEENWNE|))|N)|E)|E)))|NN))|NN)|EESSS(ENENWNEESSS(E|W)|W(SS|NN))))|N)|SSS(WNNWESSE|)SE(SWEN|)NNN))|N))|ESES(WS|EN))))|NWWNEE)|NNN(WW|E(SSEEEWWWNN|)N)))|W)|NNW(WNENW(NNEES(W|E(SS(WNSE|)SS|ENENNNESE(NN(ESNW|)W(NENWNSESWS|)WWSSWS(S|WWW|E)|SS(WNSE|)SSSENE)))|W)|S))|N))|NWW(SEWN|)N)|SSS))|SSSSSSENN(ESSNNW|)N))))|N)|S)))|N)|N))|NEN(ESNW|)WN(E|N))|NNNNNNNNNNNNNENNNW(SS|NENWNENWNENNESSSSSSSSESSW(N|SEESWSESWWNW(SSEEEENNNNNEEEEENNWWNWNEEENEEEE(NNNNWSWNNNWSSSSS(EENWESWW|)WNWSWWWNWSSE(EEE|SSSWW(S(S|EEE(EE|N))|NN(ESNW|)NNNW(SSSS|NNNESSEEEE(SWEN|)EENWNNNE(SS|NNESSENNEEESS(WSSWNNNE(WSSSENSWNNNE|)|ENNNWWNNNNE(SSS|NNWWWWNENEES(W|ENNWWWNWNNNWSWWSSWNWWNNNNENWWNEENNWNWWSES(E|WSWSWSSW(NNNNE(ENWNNW(SS|NENNW(NNEES(ENENESSWSWSW(SS(S|ENESEESENNESESWSWSSSSSWSE(SWEN|)EENNNW(NEN(W|ESSEEEESSSWW(NN(WSWN|ES)|SSESEEESSENESE(ENE(S|EEEEEEE(SWWWEEEN|)NEENWWWWNWWSWNWWNWSWNNWNWSWNNNW(NNWWSWNW(SSEEEN|W(W|NEEEEENENWWNNWSWNWWNENWWNWNNWWS(ESWSSENESESS(WNWWS(E|W(SEWN|)(N|W))|ES(W|SEN(ESE(E|N)|N)))|WWNNWNEEES(WSNE|)ENESEENESSWWSESENESEENNW(WNNEES(ENENWNEENWNNEES(ENEEESWSSENEEENEESSW(N|SEESESSSEENNNESENNEEEEESWSESWWSSSSWNWSWWNN(NNEESWS(EENNNWNEES(NWWSESNWNEES|)|S)|WSSWNWN(E|WWSSSWWSSWWWSSSWWWWSESWS(ESEENNW(NEESESWSS(ES(W|ENENW(W|NENNENNEENWNW(SWSSWNW(NE|SSES)|NEENE(NWNSES|)EESSSSESWWW(WWSSSW(SSEEEENNENEEN(EEEENESSSWNWSWS(WWN(ENEWSW|)WSWSS(EENWESWW|)S(WNWWW(SEEWWN|)WW|S)|ESE(SWEN|)NESENNNESSES(ESENENNNWSW(SEWN|)NNNESEEE(NWNENNNNWNWSSWSSS(WNNNWNWNEENNEEES(WWSSWENNEE|)ENNENESSWSEENNEEENWWWNNWWWS(EE|WNWNNWWNNEENEEEEENWNWSWNNWSWNNEENEESENENWWWWN(EEEEEEEEES(SWWWW(NEEEWWWS|)SESEEN(W|ESS(E|SSWSSSWWWWWWNN(EENNESSS(EENNNENWW(SSS|WNW(N|SWNWWN))|WW)|WSWNWS(SSENESEEESENESENEEENEN(WW(N|S)|ESSWSWSWNWSS(SSW(NNNWWW|WSWW(NEWS|)SWWSW(NWES|)S(EENEEN|S))|EE))|WW))))|ENEEES(SWNWESEN|)EEE(SSS|EN(ESESENN(W|ESSSE(NNN|SSEEEN(EEEESWS(WNWSWENESE|)SSEENN(WSNE|)NESESE(SSSWNNW(SSS(EESNWW|)WWNNES|N)|NEES(EE(NW|SEE)|W))|WW)))|WWW)))|WWWSWWSEE(ENESNWSW|)SSWSES(EE|SWWWWWWWWNWWWNNEENWWWNENN(EEEEEESEE(N(EE|W)|SSE(E|SWWSWS(EENESE|WWNW(NNEEN(NWSWWN(WSNE|)E|ES(ENSW|)SW(S|W))|WW))))|WWS(E|WWN(WWSWNWWWWWWSESESESSENESSS(W(SEWN|)NWWWWWWWWWNENENNESEE(NWNWN(WWS(E|WNWWWSWWWN(WWWSSSWSWSSSENNEEENNEN(WW(NE|SSW)|ESEEESWWSSESWSSSENNENE(NN(W(W|S)|NENWN(EESNWW|)W(NEWS|)WW)|EESWS(E|SWN(N|WSS(WWW(WWWWWNWWNEN(ESEEENNWN(WSSEWNNE|)EESE(NNWNEWSESS|)SWSS(ENSW|)WWW|WWSSW(NNNNNNENNNEN(WWSSSNNNEE|)ESS(ENNSSW|)SW(N|SSWSSEN(SWNNENSWSSEN|))|SSESWSSEEE(SWWWEEEN|)NW(W|NNN(W|E))))|S)|EE)))))|EE))|E)|ESWWS(EEENSWWW|)W(N|W))|ENESENNWWNENWWN(W(WN(EE|W)|S)|EEE(ESS(WNSE|)ESWSSSSES(W|SE(EENNESEEENEEN(ESS(ENE(NWES|)SSSSSWNNW(NEWS|)SSSESS(WNWSW(NNEWSS|)SSSEE(NWNEWSES|)ESESWSW(NNWSWS(WWNEN(NNWWSESWWW(NENWESWS|)SESSE(NN|EE)|E)|E)|SES(E(NN|SS)|W))|ENNEEEN(NWW(SE|NENW)|E(S|E)))|WW)|WWWW(SEWN|)NWWWNW(WSSE(N|E(SWSEWNEN|)EE)|NN(ES|NWNE)))|SSS))|N)))|E))))))|E(NEN(W|N)|S))|SSW(SSSWSSSSW(NNNWESSS|)WSSESESWW(SESWSSEEENNW(SWEN|)NENESSESE(NNNN(W(WNENNWNENWW(SSS(E|W(SESNWN|)W)|NEENENWN(WSSNNE|)NNEESWSEEEE(SS(WNWSS(SWN(WSNE|)NN|E)|E)|N(WWNSEE|)EENEENW(ESWWSWENEENW|)))|SS)|E)|ESSEEESSEE(EEESENE(S|N(WWWWWWN(EEEENSWWWW|)NN|EE))|S(WWWNWSSS(ENSW|)SSWNWSWNW(SSES(WSSSSSS(ENSW|)WWNNWNEN(E(NN|SSS)|W)|E)|NEENNE(SS|NN(EE|WSWWWNN(EE(SW|ENWWN)|WWWSESSWWSS(WWNNE(S|N(NNE(SSEWNN|)N(NNN|WW)|W))|ENESEE(SWEN|)NN(EESWENWW|)W(S|NN))))))|S)))|N)|N))|W))|WWWSWW(SESWW|NE))|N)|NENNW(NEWS|)(S|W)))))|WW)|S)|WNNNW(NENESENNNNESE(SSW(SEWN|)N|EENNE(SS|NNNEE(SSW(SEWN|)N|NWWWWNNWWSSWNWNN(E(NE(SEEENESSS(WNSE|)EEE(NWWNN|SSENN)|N(WWSNEE|)N)|S)|WSSWSWSEEE(NWES|)SWWWSS(WWN(WSNE|)E|SENNESSS(SENNNEN(ENEEN(W(WWSNEE|)NN|ESSWSSWN(W|N))|W)|WW))))))|S)))))|W)|W)|S))))|SSSSSS(WNWESE|)EESEE(EN(WWNNWSWN(SENESSNNWSWN|)|ESE(N|EEE(EE|N)))|S)))|SWSWW(SSEN|NE))))|SS))|N)|W)|S))|S)|SSSSESWSEEENNW(NEE(NWNNW(NE(ESSNNW|)N|SWS(E|S))|SSENESESWWSSESESWWNWNWSWNN(EENSWW|)WWSSE(N|SESESSWSSSSEENEESSSS(ENNNE(SSS|NENWNNWWWNW(N(NWES|)EEES(EENNW(NNN(W(S|W)|NEE(ENWNNWSSW(ENNESSNNWSSW|)|SWSSE(SESSSEE(NWNNN(E(SS|E)|W)|SWWWWNENN(SSWSEEWWNENN|))|N)))|S)|W)|SSWS(EEEN(ESSNNW|)W|S)))|WNN(N|WWWWW(SS|NENNW(S|NNE(S|N(E|NW(N|S)))))))))|S))))))))))|SWSWW(NEWS|)(SEEENENESE(N|S(SWNWSSWWSEESWSWS(WWWNEEN(E|WN(WWWWWSESESWS(W(SSSWWSWWW(NENWNEES(EENWESWW|)S|S(E|SS))|NN)|EE(NNN(W|E)|E))|NNE(SS|EE)))|EEEE(NW(NENNWS|W)|S)|S)|E))|WW))|N(E|N)))))))))|W)|NN)))|NEN(WNEWSE|)ESEE(ES(ENESNWSW|)WW|N(W|N)))|NEEESW)|WSWWSEESWS(E|WNWSWW(NENNNNENEE(S(SW(WSSNNE|)N|E)|NNNE(ES(ENSW|)W|NWWW(NEEEN(N|WW)|SWSW(NNEWSS|)SWWN(W(W|SSS(ENESE(NENEEN(W|N)|S)|W(SEWN|)WWWSWSSWWNWN(W(SSESWENWNN|)W|NESE(NN(ESNW|)W(W|N)|S))))|E))))|SS(WNSE|)E(SWEN|)N))))|WN(NEWS|)WSW(S|N))|S)|W))|E))|EES(ENNWESSW|)W)|S)))|E(S|NN(WSNE|)N)))|ENWNE(E|N))|N))|NNENNNW(NENWWWWW(SEWN|)NEN(ENENNNENENESSSSESWSS(EEN(W|ESSEEE(SSSS(EES(ENENSWSW|)(WW|SS)|WWNWWNWSW(S(E|S)|NNNEES(ESEE(NWES|)S|W)))|NN(ESENSWNW|)NNWNWSS(ESWSEWNENW|)WNW(NWNE(E(EEESNWWW|)S|NNNW(WWWSWNWNWSSES(E|SW(WN(NWNNNN(WSS(WN(WSNE|)N|S)|E(NWES|)S(SS|EEESENNENENWNNW(SSSWSNENNN|)NNWSWNWNEENWNENE(NWNEEN(E(N|S)|WWWSWNWSWSESSE(NENWESWS|)SWWWNENW(WSWSSS(EESEE(NNW(WWNSEE|)S|S(ENSW|)W)|WN(NNNEWSSS|)WWWSSS(WWWNENES(NWSWSEWNENES|)|E(S|EN(WNEWSE|)E)))|NN))|S(E|SS))))|E)|SS(ENSW|)S))|SS))|S)))|WNW(NENN(WSNE|)N|WS(E|W)))|W)|SS)))|SS)))|W)|NEEN(W|NESSS(WWSEESWS(WNSE|)E|EENNEE(SWSNEN|)EN(WWWNENWWSSSS(NNNNEEWWSSSS|)|E))))|SS)|S)|ESEES(ENSW|)WW)|E)|N))|E))|NNNE(S|EEENNEE(WWSSWWEENNEE|)))|N)|WW)))|NN)|NNENW(NENE(NWN(E|N)|ES(E(S|N)|W))|WWW(SEESNWWN|)W))|N)|N))|S))|NENE(NWW(NN(EE(SWEN|)EE|NNWWW(NEWS|)SSW(SEEENNWS(NESSWWEENNWS|)|NN))|S)|SEE(NWES|)SWSEENEES(EESSNNWW|)W))|S)|S)|W)|S)|E)))|NWNNW(SWWWEEEN|)NENENNE(NEE(SWEN|)N|SSSWSS)))|NENNW(NWWNEN(WW|NEESS(WNSE|)S)|S))|E))|S)|S)))|NN)|NNWSWW(SS|NENENE(S|N(N|WWSW(S|W))))))|S)|NWWSW(WNENWW(NWN(WSNE|)EEEES(SENE(NWES|)S|WW)|SS)|S))|WWWS(WNSE|)EE))|WWWWSEEEE(WWWWNEWSEEEE|))|W)|SSSS))))|NEENWNEE(S|NWWW(NEWS|)SWS(S|E)))|N)|W)|SSSWN(N|WWS(WNSE|)ESESWSESWW(EENWNEWSESWW|))))|SS)|W))$`
