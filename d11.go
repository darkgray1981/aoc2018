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

func power(x, y, serial int) int {
	id := x + 10
	level := id * (id*y + serial) / 100
	level = (level % 10) - 5

	return level
}

func p1() {

	t := time.Now()

	serial := 9798

	best := [3]int{-1, -1, -999999999}
	for y := 1; y <= 298; y++ {
		for x := 1; x <= 298; x++ {
			sum := 0

			sum += power(x, y, serial) + power(x+1, y, serial) + power(x+2, y, serial)
			sum += power(x, y+1, serial) + power(x+1, y+1, serial) + power(x+2, y+1, serial)
			sum += power(x, y+2, serial) + power(x+1, y+2, serial) + power(x+2, y+2, serial)

			if sum > best[2] {
				best = [3]int{x, y, sum}
			}
		}
	}

	fmt.Println("Done:", best, time.Since(t))
}

func p2() {

	t := time.Now()

	serial := 9798

	best := [4]int{-1, -1, -1, -999999999}

	var grid [301][301]int

	for y := 1; y <= 300; y += 1 {
		for x := 1; x <= 300; x += 1 {
			grid[y][x] = power(x, y, serial)
		}
	}

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {

			max := x
			if max < y {
				max = y
			}

			max = 301 - max

			var sum int

			for step := 1; step < max; step++ {

				for i := 0; i < step; i++ {
					sum += grid[y+i][x+step-1]
					// sum += power(x+step-1, y+i, serial)
				}

				for i := 0; i < step-1; i++ {
					sum += grid[y+step-1][x+i]
					// sum += power(x+i, y+step-1, serial)
				}

				if sum > best[3] {
					best = [4]int{x, y, step, sum}
				}

			}

		}
	}

	fmt.Println("Done 2:", best, time.Since(t))

}

const INPUT1 = ``

const INPUT = ``
