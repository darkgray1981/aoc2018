package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")
	p1()
	p2()
}

func p1() {

	t := time.Now()

	var state string
	codex := make(map[string]bool)

	for _, s := range strings.Split(INPUT, "\n") {

		var temp string
		scanned, err := fmt.Sscanf(s, "initial state: %s", &temp)
		if scanned == 1 && err == nil {
			state = temp
			continue
		}

		scanned, err = fmt.Sscanf(s, "%s => #", &temp)
		if scanned == 1 && err == nil {
			codex[temp] = true
			continue
		}
	}

	state = trim(state)
	offset := -strings.Index(state, "#")
	fmt.Println(state)

	for i := 0; i < 20; i++ {

		var b bytes.Buffer

		updated := false
		updatedOffset := 0
		for p := 0; p < len(state)-4; p++ {
			if codex[state[p:p+5]] {
				if !updated {
					updatedOffset = offset + p - 3
					updated = true
				}
				b.WriteByte('#')
			} else {
				b.WriteByte('.')
			}
		}

		offset = updatedOffset

		state = trim(b.String())
		// fmt.Println(state)
	}

	sum := 0
	for i, c := range state {
		if c == '#' {
			sum += offset + i
		}
	}

	fmt.Println("Done:", sum, len(state), time.Since(t))
}

func trim(s string) string {
	var b bytes.Buffer

	pos := strings.Index(s, "#")

	for i := 0; i < 5-pos; i++ {
		b.WriteString(".")
	}

	if pos >= 5 {
		s = s[pos-5:]
	}

	pos = strings.LastIndex(s, "#")

	b.WriteString(s[:pos+1])
	b.WriteString(".....")

	return b.String()
}

func p2() {
	t := time.Now()

	// generations := 20
	generations := 50000000000

	var state string
	codex := make(map[string]bool)

	for _, s := range strings.Split(INPUT, "\n") {

		var temp string
		scanned, err := fmt.Sscanf(s, "initial state: %s", &temp)
		if scanned == 1 && err == nil {
			state = temp
			continue
		}

		scanned, err = fmt.Sscanf(s, "%s => #", &temp)
		if scanned == 1 && err == nil {
			codex[temp] = true
			continue
		}
	}

	state = trim(state)
	offset := -strings.Index(state, "#")
	fmt.Println(state)

	for i := 0; i < generations; i++ {

		var b bytes.Buffer

		updated := false
		updatedOffset := 0
		for p := 0; p < len(state)-4; p++ {
			if codex[state[p:p+5]] {
				if !updated {
					updatedOffset = offset + p - 3
					updated = true
				}
				b.WriteByte('#')
			} else {
				b.WriteByte('.')
			}
		}

		newstate := trim(b.String())
		if newstate == state {
			fmt.Println("Found loop; skipping ahead.")
			delta := updatedOffset - offset

			offset += delta * (generations - i)
			break
		}

		offset = updatedOffset
		state = newstate
		// fmt.Println(state)
	}

	sum := 0
	for i, c := range state {
		if c == '#' {
			sum += offset + i
		}
	}

	fmt.Println("Done 2:", sum, len(state), time.Since(t))
}

const INPUT1 = `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`

const INPUT = `initial state: ##.......#.######.##..#...#.#.#..#...#..####..#.##...#....#...##..#..#.##.##.###.##.#.......###....#

.#### => .
....# => .
###.. => .
..#.# => .
##### => #
####. => .
#.##. => #
#.#.# => .
##.#. => #
.###. => .
#..#. => #
###.# => .
#.### => .
##... => #
.#.## => .
..#.. => .
#...# => #
..... => .
.##.. => .
...#. => .
#.#.. => .
.#..# => #
.#.#. => .
.#... => #
..##. => .
#..## => .
##.## => #
...## => #
..### => #
#.... => .
.##.# => #
##..# => #`
