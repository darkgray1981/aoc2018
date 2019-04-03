package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")

	t := time.Now()

	var coords [][2]int

	for _, s := range strings.Split(INPUT, "\n") {
		var x, y int

		scanned, _ := fmt.Sscanf(s, "%d, %d", &x, &y)
		if scanned != 2 {
			panic("Couldn't scan! " + s)
		}

		coords = append(coords, [2]int{x, y})
	}

	inf := make(map[[2]int]bool)

	for x := -10000; x < 10000; x++ {

		y := -10000
		var shortest [][3]int

		for _, c := range coords {
			mandis := abs(x-c[0]) + abs(y-c[1])
			if len(shortest) == 0 || mandis == shortest[0][2] {
				shortest = append(shortest, [3]int{c[0], c[1], mandis})
			} else if mandis < shortest[0][2] {
				shortest = [][3]int{[3]int{c[0], c[1], mandis}}
			}
		}

		if len(shortest) == 1 {
			inf[[2]int{shortest[0][0], shortest[0][1]}] = true
		}

		shortest = [][3]int{}

		for _, c := range coords {
			mandis := abs(y-c[0]) + abs(x-c[1])
			if len(shortest) == 0 || mandis == shortest[0][2] {
				shortest = append(shortest, [3]int{c[0], c[1], mandis})
			} else if mandis < shortest[0][2] {
				shortest = [][3]int{[3]int{c[0], c[1], mandis}}
			}
		}

		if len(shortest) == 1 {
			inf[[2]int{shortest[0][0], shortest[0][1]}] = true
		}

		y = 10000
		shortest = [][3]int{}

		for _, c := range coords {
			mandis := abs(x-c[0]) + abs(y-c[1])
			if len(shortest) == 0 || mandis == shortest[0][2] {
				shortest = append(shortest, [3]int{c[0], c[1], mandis})
			} else if mandis < shortest[0][2] {
				shortest = [][3]int{[3]int{c[0], c[1], mandis}}
			}
		}

		if len(shortest) == 1 {
			inf[[2]int{shortest[0][0], shortest[0][1]}] = true
		}

		shortest = [][3]int{}

		for _, c := range coords {
			mandis := abs(y-c[0]) + abs(x-c[1])
			if len(shortest) == 0 || mandis == shortest[0][2] {
				shortest = append(shortest, [3]int{c[0], c[1], mandis})
			} else if mandis < shortest[0][2] {
				shortest = [][3]int{[3]int{c[0], c[1], mandis}}
			}
		}

		if len(shortest) == 1 {
			inf[[2]int{shortest[0][0], shortest[0][1]}] = true
		}

	}

	areas := make(map[[2]int]int)

	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			var shortest [][3]int

			for _, c := range coords {
				mandis := abs(x-c[0]) + abs(y-c[1])
				if len(shortest) == 0 || mandis == shortest[0][2] {
					shortest = append(shortest, [3]int{c[0], c[1], mandis})
				} else if mandis < shortest[0][2] {
					shortest = [][3]int{[3]int{c[0], c[1], mandis}}
				}
			}

			if len(shortest) == 1 && !inf[[2]int{shortest[0][0], shortest[0][1]}] {
				areas[[2]int{shortest[0][0], shortest[0][1]}] += 1
			}
		}
	}

	biggest := 0

	for _, v := range areas {
		if v > biggest {
			biggest = v
		}
	}

	fmt.Println("Done:", biggest, time.Since(t))

	t = time.Now()

	size := 0

	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			total := 0

			for _, c := range coords {
				total += abs(x-c[0]) + abs(y-c[1])
			}

			if total < 10000 {
				size++
			}
		}
	}

	fmt.Println("Done 2:", size, time.Since(t))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

const INPUT1 = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

const INPUT = `177, 51
350, 132
276, 139
249, 189
225, 137
337, 354
270, 147
182, 329
118, 254
174, 280
42, 349
96, 341
236, 46
84, 253
292, 143
253, 92
224, 137
209, 325
243, 195
208, 337
197, 42
208, 87
45, 96
64, 295
266, 248
248, 298
194, 261
157, 74
52, 248
243, 201
242, 178
140, 319
69, 270
314, 302
209, 212
237, 217
86, 294
295, 144
248, 206
157, 118
155, 146
331, 40
247, 302
250, 95
193, 214
345, 89
183, 206
121, 169
79, 230
88, 155`
