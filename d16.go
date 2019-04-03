package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")
	p1()
}

func addr(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]] + a[op[2]]

	return b
}

func addi(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]] + op[2]

	return b
}

func mulr(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]] * a[op[2]]

	return b
}

func muli(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]] * op[2]

	return b
}

func banr(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]] & a[op[2]]

	return b
}

func bani(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]] & op[2]

	return b
}

func borr(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]] | a[op[2]]

	return b
}

func bori(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]] | op[2]

	return b
}

func setr(op, a [4]int) [4]int {
	b := a

	b[op[3]] = a[op[1]]

	return b
}

func seti(op, a [4]int) [4]int {
	b := a

	b[op[3]] = op[1]

	return b
}

func gtir(op, a [4]int) [4]int {
	b := a

	if op[1] > a[op[2]] {
		b[op[3]] = 1
	} else {
		b[op[3]] = 0
	}

	return b
}

func gtri(op, a [4]int) [4]int {
	b := a

	if a[op[1]] > op[2] {
		b[op[3]] = 1
	} else {
		b[op[3]] = 0
	}

	return b
}

func gtrr(op, a [4]int) [4]int {
	b := a

	if a[op[1]] > a[op[2]] {
		b[op[3]] = 1
	} else {
		b[op[3]] = 0
	}

	return b
}

func eqir(op, a [4]int) [4]int {
	b := a

	if op[1] == a[op[2]] {
		b[op[3]] = 1
	} else {
		b[op[3]] = 0
	}

	return b
}

func eqri(op, a [4]int) [4]int {
	b := a

	if a[op[1]] == op[2] {
		b[op[3]] = 1
	} else {
		b[op[3]] = 0
	}

	return b
}

func eqrr(op, a [4]int) [4]int {
	b := a

	if a[op[1]] == a[op[2]] {
		b[op[3]] = 1
	} else {
		b[op[3]] = 0
	}

	return b
}

var ops map[string](func([4]int, [4]int) [4]int)

func init() {

	ops = map[string](func([4]int, [4]int) [4]int){
		"addr": addr,
		"addi": addi,
		"mulr": mulr,
		"muli": muli,
		"banr": banr,
		"bani": bani,
		"borr": borr,
		"bori": bori,
		"setr": setr,
		"seti": seti,
		"gtir": gtir,
		"gtri": gtri,
		"gtrr": gtrr,
		"eqir": eqir,
		"eqri": eqri,
		"eqrr": eqrr,
	}
}

func p1() {

	t := time.Now()

	var events [][][4]int
	var program string

	for _, s := range strings.Split(INPUT, "\n\n") {

		if len(s) == 0 || s[0] != 'B' {
			program = s
			continue
		}

		lines := strings.SplitN(s, "\n", 3)

		var before [4]int
		scanned, err := fmt.Sscanf(lines[0], "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		if err != nil || scanned != 4 {
			panic("Couldn't scan! " + lines[0])
		}

		var instruction [4]int
		scanned, err = fmt.Sscanf(lines[1], "%d %d %d %d", &instruction[0], &instruction[1], &instruction[2], &instruction[3])
		if err != nil || scanned != 4 {
			panic("Couldn't scan! " + lines[1])
		}

		var after [4]int
		scanned, err = fmt.Sscanf(lines[2], "After:  [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])
		if err != nil || scanned != 4 {
			panic("Couldn't scan! " + lines[2])
		}

		events = append(events, [][4]int{before, instruction, after})
	}

	count := 0

	for _, e := range events {
		match := 0
		for _, f := range ops {
			if e[2] == f(e[1], e[0]) {
				match++
			}
		}

		if match >= 3 {
			count++
		}
	}

	fmt.Println("Done:", count, time.Since(t))

	opname := make(map[string]int)
	var opcode [16]string
	changed := true

	for len(opname) < len(ops) && changed {
		changed = false

	outer:
		for _, e := range events {
			if len(opcode[e[1][0]]) > 0 {
				continue
			}

			matchname := ""
			matchcode := -1
			for k, f := range ops {
				if _, ok := opname[k]; ok {
					continue
				}

				if e[2] == f(e[1], e[0]) {
					if matchcode != -1 {
						continue outer
					}

					matchname = k
					matchcode = e[1][0]
				}
			}

			if matchcode != -1 {
				opname[matchname] = matchcode
				opcode[matchcode] = matchname
				changed = true
			}
		}
	}

	if len(opname) != 16 {
		panic("Could not identify all instructions!")
	}

	var registers [4]int

	for _, s := range strings.Split(program, "\n") {

		if len(s) == 0 {
			continue
		}

		var instruction [4]int
		scanned, err := fmt.Sscanf(s, "%d %d %d %d", &instruction[0], &instruction[1], &instruction[2], &instruction[3])
		if err != nil || scanned != 4 {
			panic("Couldn't scan! " + s)
		}

		registers = ops[opcode[instruction[0]]](instruction, registers)
	}

	fmt.Println("Done 2:", registers[0], time.Since(t))
}

const INPUT = `Before: [0, 0, 2, 2]
9 2 3 0
After:  [4, 0, 2, 2]

Before: [2, 1, 2, 3]
0 1 2 3
After:  [2, 1, 2, 2]

Before: [3, 1, 3, 1]
9 0 2 2
After:  [3, 1, 6, 1]

Before: [2, 0, 0, 3]
15 0 3 3
After:  [2, 0, 0, 0]

Before: [0, 3, 3, 2]
3 1 3 1
After:  [0, 1, 3, 2]

Before: [2, 3, 3, 2]
8 2 2 3
After:  [2, 3, 3, 9]

Before: [3, 2, 3, 2]
2 1 0 3
After:  [3, 2, 3, 1]

Before: [3, 3, 2, 0]
4 1 2 3
After:  [3, 3, 2, 1]

Before: [1, 2, 3, 2]
9 1 1 3
After:  [1, 2, 3, 4]

Before: [0, 2, 3, 3]
11 0 0 3
After:  [0, 2, 3, 0]

Before: [1, 1, 3, 3]
14 1 0 0
After:  [1, 1, 3, 3]

Before: [0, 0, 0, 0]
13 0 1 3
After:  [0, 0, 0, 1]

Before: [2, 3, 1, 2]
3 1 3 1
After:  [2, 1, 1, 2]

Before: [1, 0, 3, 0]
5 1 0 1
After:  [1, 1, 3, 0]

Before: [2, 3, 3, 0]
8 2 0 3
After:  [2, 3, 3, 6]

Before: [3, 2, 2, 1]
12 3 2 0
After:  [3, 2, 2, 1]

Before: [0, 1, 0, 0]
7 1 0 2
After:  [0, 1, 1, 0]

Before: [0, 0, 1, 2]
13 0 1 3
After:  [0, 0, 1, 1]

Before: [0, 3, 0, 3]
11 0 0 0
After:  [0, 3, 0, 3]

Before: [2, 3, 2, 2]
4 1 2 2
After:  [2, 3, 1, 2]

Before: [0, 0, 2, 0]
11 0 0 0
After:  [0, 0, 2, 0]

Before: [0, 0, 1, 3]
10 2 3 3
After:  [0, 0, 1, 3]

Before: [1, 0, 2, 1]
5 1 0 0
After:  [1, 0, 2, 1]

Before: [1, 0, 3, 3]
5 1 0 2
After:  [1, 0, 1, 3]

Before: [1, 1, 3, 0]
14 1 0 1
After:  [1, 1, 3, 0]

Before: [1, 3, 2, 2]
4 1 2 3
After:  [1, 3, 2, 1]

Before: [1, 2, 1, 3]
15 0 3 1
After:  [1, 0, 1, 3]

Before: [3, 3, 1, 2]
3 1 3 1
After:  [3, 1, 1, 2]

Before: [3, 1, 2, 3]
10 1 3 0
After:  [3, 1, 2, 3]

Before: [0, 0, 1, 0]
13 0 1 1
After:  [0, 1, 1, 0]

Before: [0, 0, 3, 3]
13 0 1 0
After:  [1, 0, 3, 3]

Before: [3, 3, 3, 0]
8 2 1 3
After:  [3, 3, 3, 9]

Before: [2, 3, 0, 2]
3 1 3 0
After:  [1, 3, 0, 2]

Before: [1, 0, 3, 1]
6 0 2 3
After:  [1, 0, 3, 3]

Before: [2, 3, 0, 2]
4 1 0 0
After:  [1, 3, 0, 2]

Before: [2, 1, 0, 2]
12 3 1 0
After:  [3, 1, 0, 2]

Before: [3, 3, 2, 2]
4 1 2 2
After:  [3, 3, 1, 2]

Before: [1, 0, 2, 1]
5 1 0 1
After:  [1, 1, 2, 1]

Before: [2, 3, 2, 0]
6 0 2 1
After:  [2, 4, 2, 0]

Before: [0, 2, 3, 0]
1 0 1 0
After:  [2, 2, 3, 0]

Before: [0, 3, 2, 1]
4 1 2 0
After:  [1, 3, 2, 1]

Before: [0, 2, 2, 1]
12 3 2 3
After:  [0, 2, 2, 3]

Before: [0, 2, 2, 1]
11 0 0 1
After:  [0, 0, 2, 1]

Before: [0, 0, 2, 2]
13 0 1 3
After:  [0, 0, 2, 1]

Before: [1, 0, 0, 0]
5 1 0 1
After:  [1, 1, 0, 0]

Before: [1, 2, 3, 3]
0 2 3 2
After:  [1, 2, 9, 3]

Before: [0, 1, 1, 0]
7 1 0 1
After:  [0, 1, 1, 0]

Before: [1, 0, 0, 1]
5 1 0 0
After:  [1, 0, 0, 1]

Before: [3, 2, 2, 0]
2 1 0 1
After:  [3, 1, 2, 0]

Before: [3, 2, 2, 3]
2 1 0 1
After:  [3, 1, 2, 3]

Before: [3, 0, 2, 1]
0 0 2 0
After:  [6, 0, 2, 1]

Before: [0, 1, 1, 3]
7 1 0 2
After:  [0, 1, 1, 3]

Before: [2, 3, 0, 2]
4 1 0 3
After:  [2, 3, 0, 1]

Before: [2, 0, 1, 1]
1 1 0 2
After:  [2, 0, 2, 1]

Before: [1, 0, 3, 3]
15 0 3 0
After:  [0, 0, 3, 3]

Before: [0, 1, 0, 3]
11 0 0 0
After:  [0, 1, 0, 3]

Before: [2, 2, 3, 0]
9 2 2 1
After:  [2, 6, 3, 0]

Before: [1, 3, 3, 2]
1 0 1 2
After:  [1, 3, 3, 2]

Before: [3, 1, 2, 2]
6 2 1 0
After:  [3, 1, 2, 2]

Before: [3, 2, 1, 1]
2 1 0 0
After:  [1, 2, 1, 1]

Before: [2, 3, 3, 2]
3 1 3 2
After:  [2, 3, 1, 2]

Before: [0, 1, 2, 2]
7 1 0 0
After:  [1, 1, 2, 2]

Before: [3, 2, 3, 1]
2 1 0 2
After:  [3, 2, 1, 1]

Before: [3, 2, 0, 3]
0 3 3 3
After:  [3, 2, 0, 9]

Before: [1, 0, 2, 0]
5 1 0 3
After:  [1, 0, 2, 1]

Before: [2, 3, 1, 2]
3 1 3 3
After:  [2, 3, 1, 1]

Before: [2, 3, 0, 2]
8 1 0 0
After:  [6, 3, 0, 2]

Before: [3, 3, 1, 2]
1 2 1 3
After:  [3, 3, 1, 3]

Before: [1, 0, 2, 2]
5 1 0 2
After:  [1, 0, 1, 2]

Before: [1, 0, 0, 2]
5 1 0 3
After:  [1, 0, 0, 1]

Before: [0, 2, 2, 2]
11 0 0 0
After:  [0, 2, 2, 2]

Before: [1, 2, 2, 1]
0 0 2 2
After:  [1, 2, 2, 1]

Before: [3, 2, 1, 3]
2 1 0 2
After:  [3, 2, 1, 3]

Before: [1, 1, 1, 0]
14 1 0 0
After:  [1, 1, 1, 0]

Before: [3, 3, 1, 3]
0 3 3 3
After:  [3, 3, 1, 9]

Before: [3, 3, 2, 2]
4 1 2 1
After:  [3, 1, 2, 2]

Before: [1, 0, 0, 3]
5 1 0 1
After:  [1, 1, 0, 3]

Before: [1, 0, 0, 3]
0 3 3 0
After:  [9, 0, 0, 3]

Before: [1, 1, 1, 2]
14 1 0 1
After:  [1, 1, 1, 2]

Before: [2, 3, 0, 1]
4 1 0 1
After:  [2, 1, 0, 1]

Before: [2, 3, 3, 1]
4 1 0 2
After:  [2, 3, 1, 1]

Before: [1, 3, 2, 2]
0 0 2 0
After:  [2, 3, 2, 2]

Before: [0, 1, 2, 2]
6 0 2 1
After:  [0, 2, 2, 2]

Before: [1, 1, 3, 1]
14 1 0 3
After:  [1, 1, 3, 1]

Before: [1, 1, 2, 3]
14 1 0 2
After:  [1, 1, 1, 3]

Before: [0, 3, 0, 1]
11 0 0 3
After:  [0, 3, 0, 0]

Before: [1, 3, 0, 2]
3 1 3 3
After:  [1, 3, 0, 1]

Before: [3, 1, 0, 3]
10 1 3 1
After:  [3, 3, 0, 3]

Before: [0, 1, 2, 0]
7 1 0 2
After:  [0, 1, 1, 0]

Before: [2, 3, 3, 3]
9 1 2 1
After:  [2, 6, 3, 3]

Before: [1, 1, 1, 3]
14 1 0 0
After:  [1, 1, 1, 3]

Before: [3, 2, 2, 0]
2 1 0 3
After:  [3, 2, 2, 1]

Before: [1, 1, 2, 2]
10 0 3 3
After:  [1, 1, 2, 3]

Before: [2, 3, 2, 1]
4 1 0 1
After:  [2, 1, 2, 1]

Before: [0, 3, 3, 2]
3 1 3 0
After:  [1, 3, 3, 2]

Before: [0, 0, 0, 2]
13 0 1 3
After:  [0, 0, 0, 1]

Before: [1, 1, 0, 3]
10 1 3 2
After:  [1, 1, 3, 3]

Before: [1, 1, 2, 1]
14 1 0 3
After:  [1, 1, 2, 1]

Before: [3, 3, 3, 2]
3 1 3 1
After:  [3, 1, 3, 2]

Before: [2, 1, 3, 1]
12 3 2 2
After:  [2, 1, 3, 1]

Before: [1, 0, 0, 2]
10 0 3 3
After:  [1, 0, 0, 3]

Before: [2, 2, 2, 3]
15 0 3 3
After:  [2, 2, 2, 0]

Before: [1, 3, 2, 3]
0 1 2 1
After:  [1, 6, 2, 3]

Before: [0, 3, 3, 3]
11 0 0 0
After:  [0, 3, 3, 3]

Before: [2, 1, 1, 0]
9 0 0 1
After:  [2, 4, 1, 0]

Before: [0, 0, 2, 2]
11 0 0 2
After:  [0, 0, 0, 2]

Before: [0, 1, 3, 0]
7 1 0 1
After:  [0, 1, 3, 0]

Before: [1, 1, 3, 0]
14 1 0 3
After:  [1, 1, 3, 1]

Before: [0, 1, 3, 3]
7 1 0 0
After:  [1, 1, 3, 3]

Before: [0, 0, 3, 3]
11 0 0 2
After:  [0, 0, 0, 3]

Before: [2, 3, 2, 3]
4 1 0 2
After:  [2, 3, 1, 3]

Before: [1, 2, 1, 3]
15 0 3 0
After:  [0, 2, 1, 3]

Before: [2, 2, 1, 3]
15 0 3 0
After:  [0, 2, 1, 3]

Before: [2, 0, 2, 3]
6 2 2 1
After:  [2, 4, 2, 3]

Before: [2, 2, 1, 3]
15 0 3 1
After:  [2, 0, 1, 3]

Before: [1, 3, 2, 0]
4 1 2 0
After:  [1, 3, 2, 0]

Before: [0, 1, 3, 3]
9 3 2 0
After:  [6, 1, 3, 3]

Before: [0, 3, 2, 3]
11 0 0 3
After:  [0, 3, 2, 0]

Before: [0, 3, 2, 3]
6 0 2 3
After:  [0, 3, 2, 2]

Before: [0, 0, 2, 2]
13 0 1 2
After:  [0, 0, 1, 2]

Before: [0, 3, 2, 2]
3 1 3 1
After:  [0, 1, 2, 2]

Before: [1, 3, 1, 0]
1 0 1 2
After:  [1, 3, 3, 0]

Before: [2, 3, 1, 0]
4 1 0 0
After:  [1, 3, 1, 0]

Before: [0, 0, 1, 3]
13 0 1 3
After:  [0, 0, 1, 1]

Before: [2, 3, 3, 1]
4 1 0 3
After:  [2, 3, 3, 1]

Before: [3, 2, 1, 2]
2 1 0 1
After:  [3, 1, 1, 2]

Before: [1, 3, 2, 2]
3 1 3 2
After:  [1, 3, 1, 2]

Before: [1, 1, 1, 2]
8 0 3 0
After:  [2, 1, 1, 2]

Before: [1, 3, 3, 1]
12 3 2 2
After:  [1, 3, 3, 1]

Before: [0, 0, 2, 0]
13 0 1 2
After:  [0, 0, 1, 0]

Before: [0, 1, 0, 0]
11 0 0 2
After:  [0, 1, 0, 0]

Before: [3, 3, 2, 2]
3 1 3 2
After:  [3, 3, 1, 2]

Before: [0, 1, 1, 3]
7 1 0 0
After:  [1, 1, 1, 3]

Before: [1, 1, 1, 2]
14 1 0 3
After:  [1, 1, 1, 1]

Before: [0, 0, 2, 3]
13 0 1 1
After:  [0, 1, 2, 3]

Before: [3, 1, 0, 2]
1 2 0 1
After:  [3, 3, 0, 2]

Before: [0, 0, 1, 1]
13 0 1 0
After:  [1, 0, 1, 1]

Before: [0, 2, 2, 0]
6 0 2 0
After:  [2, 2, 2, 0]

Before: [1, 2, 0, 3]
15 0 3 1
After:  [1, 0, 0, 3]

Before: [1, 0, 0, 3]
0 3 3 1
After:  [1, 9, 0, 3]

Before: [0, 0, 2, 1]
13 0 1 3
After:  [0, 0, 2, 1]

Before: [3, 2, 2, 1]
6 1 2 2
After:  [3, 2, 4, 1]

Before: [1, 3, 0, 3]
8 3 1 0
After:  [9, 3, 0, 3]

Before: [1, 1, 3, 3]
15 0 3 0
After:  [0, 1, 3, 3]

Before: [3, 2, 2, 1]
2 1 0 1
After:  [3, 1, 2, 1]

Before: [0, 1, 1, 0]
7 1 0 0
After:  [1, 1, 1, 0]

Before: [3, 0, 2, 3]
8 3 0 0
After:  [9, 0, 2, 3]

Before: [0, 3, 0, 2]
3 1 3 2
After:  [0, 3, 1, 2]

Before: [3, 2, 0, 1]
2 1 0 0
After:  [1, 2, 0, 1]

Before: [0, 0, 2, 3]
13 0 1 3
After:  [0, 0, 2, 1]

Before: [2, 2, 2, 2]
9 1 0 2
After:  [2, 2, 4, 2]

Before: [0, 3, 2, 2]
3 1 3 0
After:  [1, 3, 2, 2]

Before: [0, 0, 1, 2]
10 2 3 2
After:  [0, 0, 3, 2]

Before: [0, 2, 3, 3]
11 0 0 2
After:  [0, 2, 0, 3]

Before: [1, 1, 3, 0]
14 1 0 0
After:  [1, 1, 3, 0]

Before: [3, 3, 0, 2]
3 1 3 0
After:  [1, 3, 0, 2]

Before: [2, 2, 0, 2]
9 0 3 0
After:  [4, 2, 0, 2]

Before: [3, 3, 2, 1]
6 2 2 2
After:  [3, 3, 4, 1]

Before: [2, 0, 3, 3]
15 0 3 3
After:  [2, 0, 3, 0]

Before: [1, 1, 3, 3]
10 0 3 3
After:  [1, 1, 3, 3]

Before: [3, 0, 3, 3]
6 1 3 1
After:  [3, 3, 3, 3]

Before: [2, 3, 0, 3]
4 1 0 0
After:  [1, 3, 0, 3]

Before: [0, 3, 2, 3]
4 1 2 3
After:  [0, 3, 2, 1]

Before: [3, 1, 2, 0]
0 1 2 3
After:  [3, 1, 2, 2]

Before: [2, 0, 3, 3]
0 2 3 3
After:  [2, 0, 3, 9]

Before: [2, 3, 2, 3]
4 1 2 3
After:  [2, 3, 2, 1]

Before: [2, 0, 1, 3]
1 2 0 0
After:  [3, 0, 1, 3]

Before: [0, 0, 3, 2]
13 0 1 1
After:  [0, 1, 3, 2]

Before: [2, 3, 2, 0]
4 1 0 2
After:  [2, 3, 1, 0]

Before: [1, 1, 2, 2]
6 0 2 1
After:  [1, 3, 2, 2]

Before: [1, 3, 1, 2]
10 2 3 1
After:  [1, 3, 1, 2]

Before: [2, 2, 0, 3]
15 0 3 2
After:  [2, 2, 0, 3]

Before: [0, 3, 0, 0]
11 0 0 3
After:  [0, 3, 0, 0]

Before: [2, 1, 0, 3]
6 0 1 2
After:  [2, 1, 3, 3]

Before: [3, 2, 1, 2]
2 1 0 2
After:  [3, 2, 1, 2]

Before: [0, 1, 0, 0]
7 1 0 3
After:  [0, 1, 0, 1]

Before: [1, 1, 1, 1]
14 1 0 0
After:  [1, 1, 1, 1]

Before: [0, 2, 2, 3]
11 0 0 1
After:  [0, 0, 2, 3]

Before: [1, 1, 0, 2]
14 1 0 0
After:  [1, 1, 0, 2]

Before: [3, 1, 2, 2]
12 3 1 0
After:  [3, 1, 2, 2]

Before: [3, 3, 1, 2]
8 0 0 3
After:  [3, 3, 1, 9]

Before: [0, 3, 1, 2]
10 2 3 3
After:  [0, 3, 1, 3]

Before: [0, 2, 2, 3]
1 0 1 2
After:  [0, 2, 2, 3]

Before: [2, 3, 3, 1]
8 0 2 3
After:  [2, 3, 3, 6]

Before: [3, 0, 2, 2]
9 2 3 0
After:  [4, 0, 2, 2]

Before: [1, 0, 1, 2]
5 1 0 3
After:  [1, 0, 1, 1]

Before: [0, 3, 2, 2]
3 1 3 3
After:  [0, 3, 2, 1]

Before: [1, 1, 1, 3]
15 0 3 2
After:  [1, 1, 0, 3]

Before: [3, 3, 2, 2]
3 1 3 1
After:  [3, 1, 2, 2]

Before: [1, 0, 2, 0]
5 1 0 1
After:  [1, 1, 2, 0]

Before: [2, 1, 2, 1]
0 3 2 3
After:  [2, 1, 2, 2]

Before: [1, 0, 2, 3]
10 0 3 1
After:  [1, 3, 2, 3]

Before: [1, 1, 2, 3]
14 1 0 1
After:  [1, 1, 2, 3]

Before: [2, 3, 3, 2]
8 3 2 1
After:  [2, 6, 3, 2]

Before: [1, 1, 2, 3]
6 0 2 2
After:  [1, 1, 3, 3]

Before: [0, 0, 2, 1]
12 3 2 0
After:  [3, 0, 2, 1]

Before: [3, 0, 0, 3]
0 3 3 1
After:  [3, 9, 0, 3]

Before: [2, 3, 0, 1]
8 3 0 1
After:  [2, 2, 0, 1]

Before: [2, 1, 3, 3]
15 0 3 1
After:  [2, 0, 3, 3]

Before: [2, 3, 1, 2]
1 0 2 3
After:  [2, 3, 1, 3]

Before: [3, 0, 3, 3]
9 2 2 0
After:  [6, 0, 3, 3]

Before: [1, 1, 1, 1]
14 1 0 2
After:  [1, 1, 1, 1]

Before: [3, 1, 0, 1]
1 1 0 1
After:  [3, 3, 0, 1]

Before: [3, 3, 2, 0]
4 1 2 2
After:  [3, 3, 1, 0]

Before: [0, 0, 1, 0]
13 0 1 3
After:  [0, 0, 1, 1]

Before: [1, 0, 2, 3]
5 1 0 0
After:  [1, 0, 2, 3]

Before: [2, 0, 2, 0]
6 2 2 1
After:  [2, 4, 2, 0]

Before: [2, 3, 3, 3]
0 2 3 0
After:  [9, 3, 3, 3]

Before: [3, 3, 0, 3]
0 1 3 2
After:  [3, 3, 9, 3]

Before: [0, 0, 3, 1]
13 0 1 0
After:  [1, 0, 3, 1]

Before: [2, 2, 2, 1]
12 3 2 0
After:  [3, 2, 2, 1]

Before: [0, 1, 0, 2]
7 1 0 3
After:  [0, 1, 0, 1]

Before: [0, 1, 1, 2]
7 1 0 2
After:  [0, 1, 1, 2]

Before: [3, 3, 3, 0]
8 2 2 0
After:  [9, 3, 3, 0]

Before: [0, 3, 0, 3]
11 0 0 2
After:  [0, 3, 0, 3]

Before: [0, 0, 3, 3]
9 2 2 2
After:  [0, 0, 6, 3]

Before: [3, 3, 3, 2]
8 2 0 0
After:  [9, 3, 3, 2]

Before: [2, 2, 2, 3]
15 0 3 1
After:  [2, 0, 2, 3]

Before: [2, 3, 1, 1]
4 1 0 0
After:  [1, 3, 1, 1]

Before: [3, 2, 1, 3]
0 0 3 3
After:  [3, 2, 1, 9]

Before: [1, 3, 3, 2]
9 3 3 3
After:  [1, 3, 3, 4]

Before: [0, 2, 3, 2]
9 2 2 3
After:  [0, 2, 3, 6]

Before: [1, 2, 2, 2]
9 3 1 0
After:  [4, 2, 2, 2]

Before: [0, 0, 2, 1]
12 3 2 1
After:  [0, 3, 2, 1]

Before: [1, 2, 0, 3]
10 0 3 2
After:  [1, 2, 3, 3]

Before: [0, 3, 3, 3]
9 3 2 1
After:  [0, 6, 3, 3]

Before: [1, 0, 3, 1]
5 1 0 0
After:  [1, 0, 3, 1]

Before: [1, 1, 1, 0]
14 1 0 3
After:  [1, 1, 1, 1]

Before: [0, 0, 2, 1]
13 0 1 2
After:  [0, 0, 1, 1]

Before: [1, 1, 0, 2]
14 1 0 3
After:  [1, 1, 0, 1]

Before: [2, 2, 2, 2]
9 2 1 2
After:  [2, 2, 4, 2]

Before: [1, 1, 2, 1]
14 1 0 0
After:  [1, 1, 2, 1]

Before: [3, 2, 2, 2]
2 1 0 1
After:  [3, 1, 2, 2]

Before: [0, 1, 2, 1]
7 1 0 1
After:  [0, 1, 2, 1]

Before: [0, 0, 3, 2]
13 0 1 2
After:  [0, 0, 1, 2]

Before: [3, 3, 1, 2]
3 1 3 3
After:  [3, 3, 1, 1]

Before: [3, 1, 2, 1]
12 3 2 3
After:  [3, 1, 2, 3]

Before: [1, 0, 0, 0]
5 1 0 0
After:  [1, 0, 0, 0]

Before: [0, 1, 1, 1]
7 1 0 0
After:  [1, 1, 1, 1]

Before: [1, 0, 0, 3]
5 1 0 0
After:  [1, 0, 0, 3]

Before: [2, 3, 2, 0]
4 1 0 0
After:  [1, 3, 2, 0]

Before: [1, 3, 1, 2]
3 1 3 3
After:  [1, 3, 1, 1]

Before: [2, 3, 2, 0]
4 1 2 0
After:  [1, 3, 2, 0]

Before: [1, 1, 2, 2]
10 1 3 0
After:  [3, 1, 2, 2]

Before: [1, 1, 0, 1]
14 1 0 2
After:  [1, 1, 1, 1]

Before: [3, 1, 2, 3]
0 0 3 3
After:  [3, 1, 2, 9]

Before: [1, 1, 2, 1]
12 3 2 0
After:  [3, 1, 2, 1]

Before: [0, 1, 2, 0]
6 3 2 3
After:  [0, 1, 2, 2]

Before: [2, 3, 0, 2]
9 3 3 1
After:  [2, 4, 0, 2]

Before: [0, 1, 1, 2]
7 1 0 0
After:  [1, 1, 1, 2]

Before: [3, 2, 0, 2]
8 0 0 2
After:  [3, 2, 9, 2]

Before: [0, 1, 3, 2]
7 1 0 3
After:  [0, 1, 3, 1]

Before: [1, 1, 0, 2]
10 0 3 0
After:  [3, 1, 0, 2]

Before: [2, 3, 2, 2]
4 1 0 3
After:  [2, 3, 2, 1]

Before: [0, 0, 3, 3]
13 0 1 3
After:  [0, 0, 3, 1]

Before: [0, 0, 0, 3]
13 0 1 3
After:  [0, 0, 0, 1]

Before: [2, 1, 2, 1]
6 1 2 3
After:  [2, 1, 2, 3]

Before: [3, 1, 1, 2]
9 3 3 1
After:  [3, 4, 1, 2]

Before: [1, 1, 1, 2]
10 0 3 1
After:  [1, 3, 1, 2]

Before: [3, 0, 1, 0]
1 2 0 2
After:  [3, 0, 3, 0]

Before: [0, 0, 2, 0]
6 3 2 0
After:  [2, 0, 2, 0]

Before: [2, 2, 0, 0]
1 3 0 1
After:  [2, 2, 0, 0]

Before: [0, 0, 2, 2]
13 0 1 1
After:  [0, 1, 2, 2]

Before: [0, 3, 1, 2]
10 2 3 1
After:  [0, 3, 1, 2]

Before: [3, 0, 2, 1]
12 3 2 2
After:  [3, 0, 3, 1]

Before: [0, 3, 3, 2]
9 2 2 2
After:  [0, 3, 6, 2]

Before: [1, 1, 0, 3]
14 1 0 1
After:  [1, 1, 0, 3]

Before: [1, 3, 0, 2]
3 1 3 1
After:  [1, 1, 0, 2]

Before: [0, 3, 3, 2]
1 0 1 0
After:  [3, 3, 3, 2]

Before: [2, 1, 3, 1]
9 0 0 3
After:  [2, 1, 3, 4]

Before: [1, 1, 2, 0]
14 1 0 0
After:  [1, 1, 2, 0]

Before: [1, 1, 2, 3]
10 0 3 1
After:  [1, 3, 2, 3]

Before: [3, 0, 1, 3]
10 2 3 0
After:  [3, 0, 1, 3]

Before: [1, 2, 2, 3]
0 0 2 0
After:  [2, 2, 2, 3]

Before: [3, 2, 2, 0]
2 1 0 0
After:  [1, 2, 2, 0]

Before: [1, 2, 0, 3]
15 0 3 2
After:  [1, 2, 0, 3]

Before: [0, 2, 1, 2]
1 0 2 3
After:  [0, 2, 1, 1]

Before: [2, 3, 2, 3]
15 0 3 1
After:  [2, 0, 2, 3]

Before: [1, 1, 3, 3]
14 1 0 1
After:  [1, 1, 3, 3]

Before: [2, 2, 2, 1]
6 1 2 3
After:  [2, 2, 2, 4]

Before: [1, 2, 2, 3]
10 0 3 2
After:  [1, 2, 3, 3]

Before: [0, 2, 1, 3]
11 0 0 2
After:  [0, 2, 0, 3]

Before: [1, 1, 2, 0]
14 1 0 2
After:  [1, 1, 1, 0]

Before: [3, 0, 2, 0]
6 3 2 1
After:  [3, 2, 2, 0]

Before: [0, 3, 2, 1]
12 3 2 0
After:  [3, 3, 2, 1]

Before: [1, 0, 3, 0]
5 1 0 3
After:  [1, 0, 3, 1]

Before: [3, 3, 2, 0]
8 0 1 1
After:  [3, 9, 2, 0]

Before: [1, 2, 3, 3]
15 0 3 0
After:  [0, 2, 3, 3]

Before: [3, 3, 3, 3]
9 2 2 1
After:  [3, 6, 3, 3]

Before: [0, 1, 2, 0]
7 1 0 3
After:  [0, 1, 2, 1]

Before: [3, 1, 3, 3]
0 3 3 2
After:  [3, 1, 9, 3]

Before: [2, 3, 1, 2]
3 1 3 2
After:  [2, 3, 1, 2]

Before: [3, 2, 0, 3]
1 2 0 0
After:  [3, 2, 0, 3]

Before: [1, 2, 2, 0]
6 1 2 2
After:  [1, 2, 4, 0]

Before: [2, 3, 3, 2]
3 1 3 3
After:  [2, 3, 3, 1]

Before: [2, 1, 3, 2]
12 3 1 1
After:  [2, 3, 3, 2]

Before: [0, 1, 3, 0]
8 2 2 3
After:  [0, 1, 3, 9]

Before: [2, 0, 1, 2]
1 2 0 0
After:  [3, 0, 1, 2]

Before: [3, 2, 0, 2]
2 1 0 3
After:  [3, 2, 0, 1]

Before: [0, 1, 2, 0]
7 1 0 1
After:  [0, 1, 2, 0]

Before: [1, 0, 3, 1]
12 3 2 0
After:  [3, 0, 3, 1]

Before: [1, 3, 2, 2]
3 1 3 1
After:  [1, 1, 2, 2]

Before: [1, 0, 1, 1]
5 1 0 0
After:  [1, 0, 1, 1]

Before: [0, 1, 0, 2]
7 1 0 0
After:  [1, 1, 0, 2]

Before: [3, 3, 2, 1]
12 3 2 0
After:  [3, 3, 2, 1]

Before: [1, 0, 1, 3]
5 1 0 2
After:  [1, 0, 1, 3]

Before: [0, 1, 3, 3]
7 1 0 3
After:  [0, 1, 3, 1]

Before: [0, 1, 0, 1]
7 1 0 0
After:  [1, 1, 0, 1]

Before: [3, 3, 0, 2]
3 1 3 1
After:  [3, 1, 0, 2]

Before: [1, 2, 2, 1]
12 3 2 1
After:  [1, 3, 2, 1]

Before: [0, 0, 0, 0]
13 0 1 2
After:  [0, 0, 1, 0]

Before: [2, 0, 0, 3]
9 0 0 3
After:  [2, 0, 0, 4]

Before: [3, 2, 2, 1]
0 3 2 2
After:  [3, 2, 2, 1]

Before: [1, 2, 0, 3]
15 0 3 3
After:  [1, 2, 0, 0]

Before: [0, 0, 2, 3]
13 0 1 0
After:  [1, 0, 2, 3]

Before: [2, 3, 3, 3]
0 3 3 0
After:  [9, 3, 3, 3]

Before: [2, 3, 3, 0]
4 1 0 0
After:  [1, 3, 3, 0]

Before: [2, 3, 2, 3]
4 1 2 0
After:  [1, 3, 2, 3]

Before: [0, 0, 2, 1]
13 0 1 1
After:  [0, 1, 2, 1]

Before: [3, 2, 2, 0]
6 1 2 3
After:  [3, 2, 2, 4]

Before: [0, 1, 0, 0]
7 1 0 0
After:  [1, 1, 0, 0]

Before: [0, 3, 1, 2]
11 0 0 0
After:  [0, 3, 1, 2]

Before: [2, 1, 1, 3]
15 0 3 1
After:  [2, 0, 1, 3]

Before: [2, 3, 1, 0]
1 3 0 3
After:  [2, 3, 1, 2]

Before: [1, 1, 2, 1]
6 2 2 2
After:  [1, 1, 4, 1]

Before: [1, 1, 1, 2]
14 1 0 0
After:  [1, 1, 1, 2]

Before: [0, 0, 2, 3]
0 2 3 0
After:  [6, 0, 2, 3]

Before: [0, 3, 3, 3]
0 1 3 2
After:  [0, 3, 9, 3]

Before: [0, 0, 0, 2]
13 0 1 2
After:  [0, 0, 1, 2]

Before: [1, 0, 0, 1]
5 1 0 2
After:  [1, 0, 1, 1]

Before: [2, 3, 2, 3]
15 0 3 2
After:  [2, 3, 0, 3]

Before: [3, 3, 3, 2]
3 1 3 3
After:  [3, 3, 3, 1]

Before: [1, 3, 3, 3]
15 0 3 0
After:  [0, 3, 3, 3]

Before: [0, 0, 3, 0]
13 0 1 0
After:  [1, 0, 3, 0]

Before: [1, 2, 2, 0]
6 0 2 3
After:  [1, 2, 2, 3]

Before: [2, 2, 2, 3]
15 0 3 0
After:  [0, 2, 2, 3]

Before: [1, 0, 1, 3]
5 1 0 3
After:  [1, 0, 1, 1]

Before: [2, 0, 3, 1]
8 3 0 0
After:  [2, 0, 3, 1]

Before: [0, 1, 0, 1]
7 1 0 3
After:  [0, 1, 0, 1]

Before: [0, 1, 1, 3]
7 1 0 3
After:  [0, 1, 1, 1]

Before: [1, 1, 1, 0]
14 1 0 2
After:  [1, 1, 1, 0]

Before: [0, 0, 3, 0]
13 0 1 3
After:  [0, 0, 3, 1]

Before: [0, 2, 3, 1]
12 3 2 1
After:  [0, 3, 3, 1]

Before: [1, 1, 2, 2]
10 1 3 1
After:  [1, 3, 2, 2]

Before: [2, 3, 2, 2]
3 1 3 0
After:  [1, 3, 2, 2]

Before: [3, 2, 0, 3]
2 1 0 3
After:  [3, 2, 0, 1]

Before: [3, 3, 1, 3]
8 0 0 2
After:  [3, 3, 9, 3]

Before: [1, 1, 3, 2]
10 0 3 1
After:  [1, 3, 3, 2]

Before: [0, 1, 2, 3]
7 1 0 0
After:  [1, 1, 2, 3]

Before: [3, 2, 3, 2]
9 2 2 3
After:  [3, 2, 3, 6]

Before: [2, 0, 1, 2]
1 1 3 1
After:  [2, 2, 1, 2]

Before: [1, 0, 2, 2]
10 0 3 0
After:  [3, 0, 2, 2]

Before: [2, 0, 3, 1]
12 3 2 1
After:  [2, 3, 3, 1]

Before: [0, 3, 2, 0]
6 0 2 3
After:  [0, 3, 2, 2]

Before: [1, 2, 2, 3]
15 0 3 1
After:  [1, 0, 2, 3]

Before: [3, 2, 0, 3]
2 1 0 0
After:  [1, 2, 0, 3]

Before: [3, 0, 1, 3]
0 3 3 3
After:  [3, 0, 1, 9]

Before: [1, 0, 3, 3]
5 1 0 3
After:  [1, 0, 3, 1]

Before: [0, 1, 0, 3]
7 1 0 0
After:  [1, 1, 0, 3]

Before: [2, 2, 3, 3]
0 1 3 0
After:  [6, 2, 3, 3]

Before: [2, 2, 1, 0]
1 0 2 1
After:  [2, 3, 1, 0]

Before: [2, 1, 2, 3]
0 3 3 0
After:  [9, 1, 2, 3]

Before: [3, 0, 3, 0]
8 2 2 2
After:  [3, 0, 9, 0]

Before: [1, 3, 2, 3]
0 0 2 1
After:  [1, 2, 2, 3]

Before: [2, 2, 0, 3]
9 0 1 1
After:  [2, 4, 0, 3]

Before: [2, 2, 1, 3]
1 2 1 0
After:  [3, 2, 1, 3]

Before: [3, 2, 2, 2]
9 2 3 1
After:  [3, 4, 2, 2]

Before: [0, 1, 0, 2]
12 3 1 0
After:  [3, 1, 0, 2]

Before: [0, 3, 2, 1]
11 0 0 3
After:  [0, 3, 2, 0]

Before: [3, 2, 1, 1]
2 1 0 2
After:  [3, 2, 1, 1]

Before: [1, 1, 1, 1]
14 1 0 3
After:  [1, 1, 1, 1]

Before: [1, 1, 0, 0]
14 1 0 1
After:  [1, 1, 0, 0]

Before: [1, 0, 1, 3]
15 0 3 1
After:  [1, 0, 1, 3]

Before: [0, 1, 2, 3]
6 0 1 2
After:  [0, 1, 1, 3]

Before: [0, 1, 1, 2]
7 1 0 3
After:  [0, 1, 1, 1]

Before: [3, 0, 2, 2]
6 3 2 1
After:  [3, 4, 2, 2]

Before: [0, 1, 1, 2]
7 1 0 1
After:  [0, 1, 1, 2]

Before: [1, 1, 2, 3]
0 2 3 0
After:  [6, 1, 2, 3]

Before: [1, 1, 3, 1]
14 1 0 0
After:  [1, 1, 3, 1]

Before: [1, 1, 1, 3]
10 2 3 1
After:  [1, 3, 1, 3]

Before: [0, 2, 2, 3]
6 1 2 1
After:  [0, 4, 2, 3]

Before: [0, 2, 1, 1]
11 0 0 0
After:  [0, 2, 1, 1]

Before: [2, 1, 2, 0]
6 3 2 0
After:  [2, 1, 2, 0]

Before: [0, 1, 0, 3]
7 1 0 3
After:  [0, 1, 0, 1]

Before: [0, 1, 2, 1]
7 1 0 0
After:  [1, 1, 2, 1]

Before: [2, 2, 2, 3]
6 2 2 2
After:  [2, 2, 4, 3]

Before: [1, 0, 2, 3]
10 0 3 0
After:  [3, 0, 2, 3]

Before: [1, 2, 2, 0]
1 3 1 2
After:  [1, 2, 2, 0]

Before: [2, 1, 1, 2]
12 3 1 1
After:  [2, 3, 1, 2]

Before: [2, 1, 2, 0]
6 0 1 3
After:  [2, 1, 2, 3]

Before: [3, 2, 0, 0]
2 1 0 1
After:  [3, 1, 0, 0]

Before: [3, 1, 3, 1]
9 0 2 0
After:  [6, 1, 3, 1]

Before: [0, 3, 1, 2]
11 0 0 1
After:  [0, 0, 1, 2]

Before: [3, 1, 3, 1]
9 0 2 1
After:  [3, 6, 3, 1]

Before: [3, 3, 3, 3]
8 2 0 0
After:  [9, 3, 3, 3]

Before: [1, 2, 2, 3]
15 0 3 3
After:  [1, 2, 2, 0]

Before: [1, 2, 3, 3]
15 0 3 1
After:  [1, 0, 3, 3]

Before: [3, 1, 0, 2]
12 3 1 0
After:  [3, 1, 0, 2]

Before: [0, 0, 0, 0]
13 0 1 0
After:  [1, 0, 0, 0]

Before: [0, 1, 1, 2]
12 3 1 1
After:  [0, 3, 1, 2]

Before: [3, 2, 1, 0]
2 1 0 1
After:  [3, 1, 1, 0]

Before: [0, 1, 3, 1]
7 1 0 1
After:  [0, 1, 3, 1]

Before: [2, 1, 0, 3]
15 0 3 3
After:  [2, 1, 0, 0]

Before: [2, 3, 0, 3]
4 1 0 3
After:  [2, 3, 0, 1]

Before: [1, 1, 3, 1]
14 1 0 1
After:  [1, 1, 3, 1]

Before: [2, 1, 2, 2]
12 3 1 3
After:  [2, 1, 2, 3]

Before: [1, 0, 2, 2]
5 1 0 1
After:  [1, 1, 2, 2]

Before: [0, 0, 3, 1]
11 0 0 1
After:  [0, 0, 3, 1]

Before: [3, 2, 0, 0]
2 1 0 0
After:  [1, 2, 0, 0]

Before: [0, 3, 0, 3]
0 1 3 0
After:  [9, 3, 0, 3]

Before: [2, 2, 2, 2]
9 3 3 1
After:  [2, 4, 2, 2]

Before: [1, 3, 0, 2]
3 1 3 2
After:  [1, 3, 1, 2]

Before: [3, 2, 3, 3]
2 1 0 0
After:  [1, 2, 3, 3]

Before: [1, 3, 1, 0]
9 0 2 3
After:  [1, 3, 1, 2]

Before: [3, 1, 0, 2]
1 1 0 3
After:  [3, 1, 0, 3]

Before: [0, 1, 1, 3]
7 1 0 1
After:  [0, 1, 1, 3]

Before: [3, 2, 0, 2]
2 1 0 0
After:  [1, 2, 0, 2]

Before: [0, 0, 0, 2]
13 0 1 1
After:  [0, 1, 0, 2]

Before: [3, 2, 2, 1]
2 1 0 2
After:  [3, 2, 1, 1]

Before: [0, 2, 2, 0]
6 2 2 3
After:  [0, 2, 2, 4]

Before: [3, 3, 1, 1]
8 1 0 1
After:  [3, 9, 1, 1]

Before: [0, 2, 1, 0]
11 0 0 2
After:  [0, 2, 0, 0]

Before: [0, 1, 2, 2]
7 1 0 1
After:  [0, 1, 2, 2]

Before: [1, 1, 0, 0]
14 1 0 0
After:  [1, 1, 0, 0]

Before: [0, 2, 3, 2]
11 0 0 1
After:  [0, 0, 3, 2]

Before: [0, 1, 3, 1]
11 0 0 0
After:  [0, 1, 3, 1]

Before: [1, 0, 0, 2]
5 1 0 0
After:  [1, 0, 0, 2]

Before: [0, 0, 0, 1]
1 0 3 0
After:  [1, 0, 0, 1]

Before: [1, 0, 2, 0]
5 1 0 2
After:  [1, 0, 1, 0]

Before: [3, 2, 0, 2]
2 1 0 2
After:  [3, 2, 1, 2]

Before: [0, 3, 0, 2]
3 1 3 3
After:  [0, 3, 0, 1]

Before: [2, 3, 0, 0]
4 1 0 0
After:  [1, 3, 0, 0]

Before: [1, 1, 0, 1]
14 1 0 1
After:  [1, 1, 0, 1]

Before: [1, 1, 2, 1]
12 3 2 1
After:  [1, 3, 2, 1]

Before: [2, 0, 0, 3]
15 0 3 1
After:  [2, 0, 0, 3]

Before: [3, 2, 3, 1]
8 2 0 1
After:  [3, 9, 3, 1]

Before: [2, 2, 0, 1]
9 0 1 1
After:  [2, 4, 0, 1]

Before: [2, 3, 3, 3]
15 0 3 1
After:  [2, 0, 3, 3]

Before: [2, 0, 2, 3]
9 0 0 1
After:  [2, 4, 2, 3]

Before: [3, 2, 2, 2]
2 1 0 3
After:  [3, 2, 2, 1]

Before: [2, 3, 2, 2]
4 1 2 1
After:  [2, 1, 2, 2]

Before: [1, 2, 3, 3]
1 1 0 1
After:  [1, 3, 3, 3]

Before: [0, 3, 2, 3]
6 0 2 0
After:  [2, 3, 2, 3]

Before: [2, 1, 2, 3]
15 0 3 0
After:  [0, 1, 2, 3]

Before: [2, 1, 2, 0]
9 0 0 3
After:  [2, 1, 2, 4]

Before: [1, 2, 1, 3]
0 3 3 1
After:  [1, 9, 1, 3]

Before: [2, 2, 1, 1]
1 2 1 2
After:  [2, 2, 3, 1]

Before: [0, 3, 2, 1]
1 0 1 0
After:  [3, 3, 2, 1]

Before: [2, 1, 2, 1]
9 2 0 2
After:  [2, 1, 4, 1]

Before: [0, 0, 3, 0]
13 0 1 1
After:  [0, 1, 3, 0]

Before: [1, 1, 2, 3]
6 0 2 3
After:  [1, 1, 2, 3]

Before: [2, 1, 1, 2]
10 1 3 0
After:  [3, 1, 1, 2]

Before: [1, 1, 3, 2]
12 3 1 1
After:  [1, 3, 3, 2]

Before: [2, 3, 3, 1]
8 0 2 1
After:  [2, 6, 3, 1]

Before: [1, 0, 2, 3]
5 1 0 3
After:  [1, 0, 2, 1]

Before: [2, 3, 2, 1]
4 1 2 0
After:  [1, 3, 2, 1]

Before: [2, 2, 0, 2]
9 3 1 2
After:  [2, 2, 4, 2]

Before: [2, 2, 1, 3]
10 2 3 0
After:  [3, 2, 1, 3]

Before: [0, 0, 2, 3]
13 0 1 2
After:  [0, 0, 1, 3]

Before: [2, 1, 3, 1]
12 3 2 3
After:  [2, 1, 3, 3]

Before: [2, 1, 3, 1]
6 1 2 2
After:  [2, 1, 3, 1]

Before: [1, 2, 1, 0]
1 1 2 3
After:  [1, 2, 1, 3]

Before: [3, 1, 2, 2]
0 0 2 0
After:  [6, 1, 2, 2]

Before: [2, 0, 3, 3]
8 2 2 3
After:  [2, 0, 3, 9]

Before: [1, 0, 1, 1]
5 1 0 2
After:  [1, 0, 1, 1]

Before: [3, 3, 0, 2]
3 1 3 2
After:  [3, 3, 1, 2]

Before: [1, 3, 2, 3]
15 0 3 3
After:  [1, 3, 2, 0]

Before: [2, 3, 2, 3]
4 1 0 0
After:  [1, 3, 2, 3]

Before: [2, 0, 1, 3]
15 0 3 3
After:  [2, 0, 1, 0]

Before: [0, 0, 1, 2]
13 0 1 1
After:  [0, 1, 1, 2]

Before: [2, 1, 3, 3]
15 0 3 2
After:  [2, 1, 0, 3]

Before: [1, 1, 0, 3]
14 1 0 0
After:  [1, 1, 0, 3]

Before: [2, 3, 2, 2]
3 1 3 1
After:  [2, 1, 2, 2]

Before: [0, 1, 3, 2]
11 0 0 2
After:  [0, 1, 0, 2]

Before: [2, 3, 3, 2]
4 1 0 3
After:  [2, 3, 3, 1]

Before: [1, 1, 0, 3]
15 0 3 1
After:  [1, 0, 0, 3]

Before: [2, 1, 0, 1]
9 0 0 0
After:  [4, 1, 0, 1]

Before: [1, 3, 2, 2]
4 1 2 0
After:  [1, 3, 2, 2]

Before: [1, 2, 1, 3]
1 1 0 0
After:  [3, 2, 1, 3]

Before: [2, 2, 1, 2]
1 2 0 2
After:  [2, 2, 3, 2]

Before: [0, 1, 3, 2]
7 1 0 0
After:  [1, 1, 3, 2]

Before: [0, 1, 2, 2]
7 1 0 3
After:  [0, 1, 2, 1]

Before: [0, 2, 1, 0]
11 0 0 3
After:  [0, 2, 1, 0]

Before: [1, 1, 2, 1]
14 1 0 2
After:  [1, 1, 1, 1]

Before: [1, 0, 3, 1]
12 3 2 1
After:  [1, 3, 3, 1]

Before: [1, 3, 3, 2]
9 3 3 2
After:  [1, 3, 4, 2]

Before: [1, 0, 1, 3]
10 2 3 3
After:  [1, 0, 1, 3]

Before: [1, 0, 1, 0]
5 1 0 2
After:  [1, 0, 1, 0]

Before: [0, 1, 3, 2]
7 1 0 1
After:  [0, 1, 3, 2]

Before: [1, 1, 1, 2]
14 1 0 2
After:  [1, 1, 1, 2]

Before: [1, 3, 3, 3]
0 2 3 0
After:  [9, 3, 3, 3]

Before: [0, 0, 1, 0]
13 0 1 0
After:  [1, 0, 1, 0]

Before: [3, 1, 2, 2]
9 2 3 1
After:  [3, 4, 2, 2]

Before: [1, 0, 1, 1]
5 1 0 3
After:  [1, 0, 1, 1]

Before: [1, 0, 3, 3]
5 1 0 1
After:  [1, 1, 3, 3]

Before: [0, 1, 2, 1]
7 1 0 3
After:  [0, 1, 2, 1]

Before: [0, 3, 1, 1]
11 0 0 0
After:  [0, 3, 1, 1]

Before: [0, 2, 0, 3]
6 2 3 2
After:  [0, 2, 3, 3]

Before: [2, 1, 3, 2]
8 2 0 3
After:  [2, 1, 3, 6]

Before: [1, 0, 3, 0]
5 1 0 2
After:  [1, 0, 1, 0]

Before: [0, 0, 1, 3]
13 0 1 1
After:  [0, 1, 1, 3]

Before: [1, 1, 1, 2]
10 1 3 3
After:  [1, 1, 1, 3]

Before: [0, 1, 0, 2]
12 3 1 1
After:  [0, 3, 0, 2]

Before: [2, 3, 3, 2]
9 0 3 2
After:  [2, 3, 4, 2]

Before: [0, 2, 1, 0]
1 1 2 1
After:  [0, 3, 1, 0]

Before: [0, 2, 2, 2]
6 3 2 2
After:  [0, 2, 4, 2]

Before: [3, 1, 3, 0]
1 1 0 0
After:  [3, 1, 3, 0]

Before: [1, 0, 1, 1]
5 1 0 1
After:  [1, 1, 1, 1]

Before: [3, 2, 1, 3]
2 1 0 3
After:  [3, 2, 1, 1]

Before: [0, 1, 3, 1]
11 0 0 3
After:  [0, 1, 3, 0]

Before: [3, 1, 0, 2]
10 1 3 3
After:  [3, 1, 0, 3]

Before: [2, 3, 3, 0]
4 1 0 3
After:  [2, 3, 3, 1]

Before: [2, 3, 3, 2]
4 1 0 2
After:  [2, 3, 1, 2]

Before: [3, 2, 0, 0]
2 1 0 3
After:  [3, 2, 0, 1]

Before: [1, 0, 2, 1]
6 1 2 1
After:  [1, 2, 2, 1]

Before: [3, 2, 1, 0]
2 1 0 3
After:  [3, 2, 1, 1]

Before: [2, 2, 3, 0]
9 0 0 1
After:  [2, 4, 3, 0]

Before: [1, 1, 1, 0]
14 1 0 1
After:  [1, 1, 1, 0]

Before: [0, 2, 1, 2]
10 2 3 0
After:  [3, 2, 1, 2]

Before: [1, 1, 1, 3]
15 0 3 0
After:  [0, 1, 1, 3]

Before: [0, 1, 2, 0]
7 1 0 0
After:  [1, 1, 2, 0]

Before: [0, 1, 2, 3]
7 1 0 3
After:  [0, 1, 2, 1]

Before: [1, 0, 3, 1]
12 3 2 2
After:  [1, 0, 3, 1]

Before: [2, 3, 0, 0]
4 1 0 2
After:  [2, 3, 1, 0]

Before: [3, 2, 3, 3]
2 1 0 1
After:  [3, 1, 3, 3]

Before: [0, 0, 1, 3]
13 0 1 2
After:  [0, 0, 1, 3]

Before: [3, 2, 2, 1]
2 1 0 3
After:  [3, 2, 2, 1]

Before: [3, 1, 1, 3]
10 1 3 2
After:  [3, 1, 3, 3]

Before: [0, 0, 0, 1]
13 0 1 3
After:  [0, 0, 0, 1]

Before: [1, 3, 2, 2]
6 2 2 3
After:  [1, 3, 2, 4]

Before: [0, 1, 2, 1]
7 1 0 2
After:  [0, 1, 1, 1]

Before: [2, 0, 1, 1]
9 3 2 2
After:  [2, 0, 2, 1]

Before: [1, 2, 3, 1]
12 3 2 3
After:  [1, 2, 3, 3]

Before: [0, 3, 1, 0]
8 1 1 2
After:  [0, 3, 9, 0]

Before: [1, 0, 2, 3]
5 1 0 2
After:  [1, 0, 1, 3]

Before: [1, 1, 0, 0]
14 1 0 2
After:  [1, 1, 1, 0]

Before: [3, 3, 3, 3]
0 0 3 3
After:  [3, 3, 3, 9]

Before: [0, 1, 3, 0]
7 1 0 3
After:  [0, 1, 3, 1]

Before: [1, 1, 1, 2]
9 0 2 3
After:  [1, 1, 1, 2]

Before: [1, 0, 0, 3]
15 0 3 3
After:  [1, 0, 0, 0]

Before: [2, 3, 3, 3]
9 3 2 0
After:  [6, 3, 3, 3]

Before: [0, 0, 0, 3]
11 0 0 1
After:  [0, 0, 0, 3]

Before: [2, 1, 2, 2]
0 1 2 3
After:  [2, 1, 2, 2]

Before: [0, 0, 1, 1]
13 0 1 1
After:  [0, 1, 1, 1]

Before: [1, 2, 3, 3]
10 0 3 3
After:  [1, 2, 3, 3]

Before: [0, 2, 2, 0]
11 0 0 1
After:  [0, 0, 2, 0]

Before: [2, 3, 3, 0]
8 2 1 1
After:  [2, 9, 3, 0]

Before: [3, 2, 3, 2]
2 1 0 0
After:  [1, 2, 3, 2]

Before: [0, 2, 3, 2]
1 0 3 0
After:  [2, 2, 3, 2]

Before: [0, 1, 3, 3]
10 1 3 2
After:  [0, 1, 3, 3]

Before: [3, 3, 2, 0]
4 1 2 1
After:  [3, 1, 2, 0]

Before: [3, 3, 2, 1]
0 1 2 3
After:  [3, 3, 2, 6]

Before: [1, 3, 1, 2]
3 1 3 0
After:  [1, 3, 1, 2]

Before: [0, 1, 0, 0]
7 1 0 1
After:  [0, 1, 0, 0]

Before: [1, 0, 3, 1]
5 1 0 3
After:  [1, 0, 3, 1]

Before: [3, 3, 2, 3]
4 1 2 0
After:  [1, 3, 2, 3]

Before: [2, 0, 0, 2]
9 3 0 2
After:  [2, 0, 4, 2]

Before: [0, 2, 0, 2]
11 0 0 2
After:  [0, 2, 0, 2]

Before: [3, 2, 3, 3]
2 1 0 2
After:  [3, 2, 1, 3]

Before: [0, 1, 0, 2]
10 1 3 1
After:  [0, 3, 0, 2]

Before: [2, 3, 2, 1]
9 2 0 3
After:  [2, 3, 2, 4]

Before: [0, 0, 2, 3]
11 0 0 0
After:  [0, 0, 2, 3]

Before: [3, 2, 0, 1]
8 0 0 3
After:  [3, 2, 0, 9]

Before: [3, 2, 2, 1]
2 1 0 0
After:  [1, 2, 2, 1]

Before: [3, 1, 3, 1]
12 3 2 1
After:  [3, 3, 3, 1]

Before: [3, 1, 3, 3]
10 1 3 0
After:  [3, 1, 3, 3]

Before: [1, 1, 3, 3]
14 1 0 3
After:  [1, 1, 3, 1]

Before: [1, 1, 0, 1]
14 1 0 3
After:  [1, 1, 0, 1]

Before: [2, 1, 1, 3]
15 0 3 2
After:  [2, 1, 0, 3]

Before: [0, 0, 0, 1]
11 0 0 3
After:  [0, 0, 0, 0]

Before: [0, 1, 1, 2]
10 1 3 3
After:  [0, 1, 1, 3]

Before: [1, 0, 0, 3]
6 2 3 0
After:  [3, 0, 0, 3]

Before: [2, 3, 0, 2]
4 1 0 2
After:  [2, 3, 1, 2]

Before: [2, 3, 0, 0]
4 1 0 1
After:  [2, 1, 0, 0]

Before: [3, 3, 3, 2]
3 1 3 0
After:  [1, 3, 3, 2]

Before: [0, 0, 0, 0]
13 0 1 1
After:  [0, 1, 0, 0]

Before: [2, 1, 2, 3]
15 0 3 2
After:  [2, 1, 0, 3]

Before: [0, 0, 0, 1]
13 0 1 1
After:  [0, 1, 0, 1]

Before: [1, 2, 3, 2]
9 1 3 0
After:  [4, 2, 3, 2]

Before: [1, 3, 0, 3]
15 0 3 3
After:  [1, 3, 0, 0]

Before: [0, 2, 2, 0]
11 0 0 2
After:  [0, 2, 0, 0]

Before: [1, 1, 2, 3]
0 1 2 3
After:  [1, 1, 2, 2]

Before: [2, 3, 2, 3]
15 0 3 0
After:  [0, 3, 2, 3]

Before: [1, 0, 0, 0]
5 1 0 2
After:  [1, 0, 1, 0]

Before: [1, 3, 2, 3]
4 1 2 1
After:  [1, 1, 2, 3]

Before: [0, 1, 1, 3]
6 0 3 0
After:  [3, 1, 1, 3]

Before: [0, 1, 0, 2]
7 1 0 1
After:  [0, 1, 0, 2]

Before: [2, 0, 2, 1]
6 0 2 3
After:  [2, 0, 2, 4]

Before: [0, 0, 1, 0]
13 0 1 2
After:  [0, 0, 1, 0]

Before: [0, 3, 1, 2]
3 1 3 1
After:  [0, 1, 1, 2]

Before: [3, 2, 2, 2]
2 1 0 2
After:  [3, 2, 1, 2]

Before: [1, 2, 2, 3]
10 0 3 0
After:  [3, 2, 2, 3]

Before: [1, 2, 2, 0]
9 1 1 1
After:  [1, 4, 2, 0]

Before: [1, 1, 2, 1]
14 1 0 1
After:  [1, 1, 2, 1]

Before: [1, 3, 2, 1]
4 1 2 2
After:  [1, 3, 1, 1]

Before: [1, 1, 1, 3]
15 0 3 1
After:  [1, 0, 1, 3]

Before: [3, 2, 0, 3]
2 1 0 1
After:  [3, 1, 0, 3]

Before: [1, 1, 3, 2]
14 1 0 3
After:  [1, 1, 3, 1]

Before: [1, 1, 0, 3]
15 0 3 2
After:  [1, 1, 0, 3]

Before: [2, 2, 0, 2]
9 1 0 2
After:  [2, 2, 4, 2]

Before: [2, 0, 2, 3]
15 0 3 3
After:  [2, 0, 2, 0]

Before: [2, 0, 1, 2]
1 1 0 0
After:  [2, 0, 1, 2]

Before: [1, 1, 1, 3]
14 1 0 1
After:  [1, 1, 1, 3]

Before: [0, 3, 1, 0]
11 0 0 0
After:  [0, 3, 1, 0]

Before: [1, 3, 2, 3]
4 1 2 3
After:  [1, 3, 2, 1]

Before: [2, 0, 2, 3]
15 0 3 1
After:  [2, 0, 2, 3]

Before: [0, 1, 1, 0]
11 0 0 2
After:  [0, 1, 0, 0]

Before: [3, 2, 2, 2]
8 3 0 2
After:  [3, 2, 6, 2]

Before: [3, 2, 1, 3]
10 2 3 2
After:  [3, 2, 3, 3]

Before: [1, 1, 0, 2]
14 1 0 2
After:  [1, 1, 1, 2]

Before: [1, 3, 2, 3]
15 0 3 1
After:  [1, 0, 2, 3]

Before: [3, 3, 3, 3]
9 3 2 3
After:  [3, 3, 3, 6]

Before: [2, 3, 0, 3]
0 1 3 3
After:  [2, 3, 0, 9]

Before: [3, 0, 1, 2]
1 1 0 0
After:  [3, 0, 1, 2]

Before: [1, 1, 1, 3]
14 1 0 3
After:  [1, 1, 1, 1]

Before: [0, 3, 1, 1]
11 0 0 3
After:  [0, 3, 1, 0]

Before: [0, 0, 1, 1]
13 0 1 2
After:  [0, 0, 1, 1]

Before: [0, 2, 0, 1]
1 0 1 3
After:  [0, 2, 0, 2]

Before: [2, 1, 1, 2]
12 3 1 0
After:  [3, 1, 1, 2]

Before: [1, 3, 2, 0]
4 1 2 2
After:  [1, 3, 1, 0]

Before: [0, 0, 3, 2]
13 0 1 3
After:  [0, 0, 3, 1]

Before: [2, 2, 2, 3]
0 3 3 0
After:  [9, 2, 2, 3]

Before: [1, 2, 1, 3]
10 0 3 0
After:  [3, 2, 1, 3]

Before: [1, 1, 1, 3]
14 1 0 2
After:  [1, 1, 1, 3]

Before: [0, 3, 0, 2]
3 1 3 1
After:  [0, 1, 0, 2]

Before: [0, 0, 1, 2]
13 0 1 2
After:  [0, 0, 1, 2]

Before: [3, 2, 3, 0]
2 1 0 1
After:  [3, 1, 3, 0]

Before: [0, 0, 2, 3]
0 3 2 1
After:  [0, 6, 2, 3]

Before: [0, 2, 1, 0]
1 2 1 1
After:  [0, 3, 1, 0]

Before: [3, 3, 2, 2]
4 1 2 0
After:  [1, 3, 2, 2]

Before: [1, 1, 0, 0]
14 1 0 3
After:  [1, 1, 0, 1]

Before: [2, 1, 2, 1]
6 2 2 2
After:  [2, 1, 4, 1]

Before: [3, 2, 0, 3]
2 1 0 2
After:  [3, 2, 1, 3]

Before: [0, 2, 1, 3]
11 0 0 0
After:  [0, 2, 1, 3]

Before: [0, 0, 0, 3]
13 0 1 0
After:  [1, 0, 0, 3]

Before: [3, 1, 1, 2]
10 1 3 0
After:  [3, 1, 1, 2]

Before: [3, 2, 0, 1]
2 1 0 3
After:  [3, 2, 0, 1]

Before: [1, 0, 0, 2]
10 0 3 0
After:  [3, 0, 0, 2]

Before: [2, 3, 2, 2]
9 2 3 1
After:  [2, 4, 2, 2]

Before: [0, 0, 2, 1]
13 0 1 0
After:  [1, 0, 2, 1]

Before: [0, 3, 2, 3]
11 0 0 1
After:  [0, 0, 2, 3]

Before: [1, 1, 0, 1]
14 1 0 0
After:  [1, 1, 0, 1]

Before: [1, 2, 1, 2]
10 0 3 0
After:  [3, 2, 1, 2]

Before: [1, 0, 1, 0]
5 1 0 1
After:  [1, 1, 1, 0]

Before: [1, 3, 3, 1]
9 2 2 3
After:  [1, 3, 3, 6]

Before: [1, 2, 0, 2]
1 0 1 0
After:  [3, 2, 0, 2]

Before: [3, 1, 0, 0]
1 1 0 1
After:  [3, 3, 0, 0]

Before: [3, 0, 3, 0]
8 2 0 3
After:  [3, 0, 3, 9]

Before: [0, 0, 1, 2]
11 0 0 1
After:  [0, 0, 1, 2]

Before: [1, 0, 1, 2]
10 0 3 1
After:  [1, 3, 1, 2]

Before: [0, 0, 0, 1]
1 0 3 1
After:  [0, 1, 0, 1]

Before: [1, 1, 2, 2]
14 1 0 0
After:  [1, 1, 2, 2]

Before: [1, 1, 3, 2]
14 1 0 2
After:  [1, 1, 1, 2]

Before: [2, 1, 2, 1]
1 1 0 2
After:  [2, 1, 3, 1]

Before: [0, 0, 0, 3]
13 0 1 1
After:  [0, 1, 0, 3]

Before: [0, 0, 0, 2]
11 0 0 2
After:  [0, 0, 0, 2]

Before: [3, 2, 3, 0]
2 1 0 2
After:  [3, 2, 1, 0]

Before: [1, 1, 0, 3]
14 1 0 3
After:  [1, 1, 0, 1]

Before: [3, 2, 1, 2]
2 1 0 0
After:  [1, 2, 1, 2]

Before: [2, 3, 2, 2]
4 1 0 0
After:  [1, 3, 2, 2]

Before: [0, 1, 3, 1]
7 1 0 0
After:  [1, 1, 3, 1]

Before: [1, 2, 3, 1]
12 3 2 1
After:  [1, 3, 3, 1]

Before: [1, 0, 0, 1]
5 1 0 3
After:  [1, 0, 0, 1]

Before: [2, 3, 1, 2]
4 1 0 2
After:  [2, 3, 1, 2]

Before: [0, 0, 1, 1]
11 0 0 0
After:  [0, 0, 1, 1]

Before: [2, 1, 2, 2]
6 0 2 1
After:  [2, 4, 2, 2]

Before: [0, 3, 2, 2]
9 2 3 1
After:  [0, 4, 2, 2]

Before: [3, 2, 2, 3]
2 1 0 3
After:  [3, 2, 2, 1]

Before: [3, 1, 1, 3]
10 2 3 0
After:  [3, 1, 1, 3]

Before: [2, 1, 3, 2]
10 1 3 3
After:  [2, 1, 3, 3]

Before: [3, 3, 1, 2]
3 1 3 2
After:  [3, 3, 1, 2]

Before: [0, 2, 3, 1]
11 0 0 1
After:  [0, 0, 3, 1]

Before: [0, 0, 0, 2]
11 0 0 3
After:  [0, 0, 0, 0]

Before: [1, 3, 2, 2]
3 1 3 3
After:  [1, 3, 2, 1]

Before: [0, 1, 0, 1]
7 1 0 1
After:  [0, 1, 0, 1]

Before: [2, 1, 0, 2]
8 1 0 3
After:  [2, 1, 0, 2]

Before: [2, 2, 0, 3]
0 3 3 3
After:  [2, 2, 0, 9]

Before: [1, 2, 0, 2]
8 0 3 1
After:  [1, 2, 0, 2]

Before: [0, 0, 3, 3]
13 0 1 2
After:  [0, 0, 1, 3]

Before: [3, 3, 3, 2]
3 1 3 2
After:  [3, 3, 1, 2]

Before: [3, 2, 0, 1]
2 1 0 2
After:  [3, 2, 1, 1]

Before: [3, 2, 3, 2]
2 1 0 1
After:  [3, 1, 3, 2]

Before: [0, 1, 3, 0]
7 1 0 0
After:  [1, 1, 3, 0]

Before: [2, 0, 3, 3]
15 0 3 1
After:  [2, 0, 3, 3]

Before: [2, 1, 2, 3]
15 0 3 1
After:  [2, 0, 2, 3]

Before: [1, 1, 3, 1]
14 1 0 2
After:  [1, 1, 1, 1]

Before: [1, 1, 3, 0]
14 1 0 2
After:  [1, 1, 1, 0]

Before: [1, 0, 0, 3]
5 1 0 2
After:  [1, 0, 1, 3]

Before: [0, 1, 0, 1]
7 1 0 2
After:  [0, 1, 1, 1]

Before: [0, 3, 3, 3]
8 3 1 0
After:  [9, 3, 3, 3]

Before: [1, 0, 3, 3]
15 0 3 2
After:  [1, 0, 0, 3]

Before: [1, 0, 3, 3]
5 1 0 0
After:  [1, 0, 3, 3]

Before: [1, 3, 0, 2]
3 1 3 0
After:  [1, 3, 0, 2]

Before: [0, 1, 1, 1]
7 1 0 3
After:  [0, 1, 1, 1]

Before: [0, 0, 3, 2]
11 0 0 0
After:  [0, 0, 3, 2]

Before: [1, 1, 3, 3]
14 1 0 2
After:  [1, 1, 1, 3]

Before: [1, 1, 1, 2]
12 3 1 2
After:  [1, 1, 3, 2]

Before: [2, 3, 1, 2]
3 1 3 0
After:  [1, 3, 1, 2]

Before: [1, 0, 0, 2]
5 1 0 2
After:  [1, 0, 1, 2]

Before: [0, 3, 2, 3]
0 1 2 1
After:  [0, 6, 2, 3]

Before: [2, 3, 0, 2]
3 1 3 1
After:  [2, 1, 0, 2]

Before: [0, 0, 0, 1]
13 0 1 0
After:  [1, 0, 0, 1]

Before: [2, 3, 1, 2]
4 1 0 3
After:  [2, 3, 1, 1]

Before: [2, 3, 2, 2]
3 1 3 3
After:  [2, 3, 2, 1]

Before: [3, 2, 3, 0]
2 1 0 0
After:  [1, 2, 3, 0]

Before: [2, 3, 2, 0]
6 2 2 3
After:  [2, 3, 2, 4]

Before: [2, 1, 1, 0]
1 3 0 2
After:  [2, 1, 2, 0]

Before: [2, 2, 3, 3]
9 3 2 2
After:  [2, 2, 6, 3]

Before: [1, 0, 0, 2]
5 1 0 1
After:  [1, 1, 0, 2]

Before: [1, 1, 0, 2]
10 1 3 3
After:  [1, 1, 0, 3]

Before: [0, 1, 3, 2]
7 1 0 2
After:  [0, 1, 1, 2]

Before: [2, 1, 0, 2]
12 3 1 3
After:  [2, 1, 0, 3]

Before: [1, 3, 2, 3]
15 0 3 0
After:  [0, 3, 2, 3]

Before: [0, 2, 3, 3]
0 2 3 0
After:  [9, 2, 3, 3]

Before: [2, 1, 2, 1]
6 0 1 2
After:  [2, 1, 3, 1]

Before: [1, 0, 3, 2]
5 1 0 3
After:  [1, 0, 3, 1]

Before: [1, 2, 2, 3]
15 0 3 0
After:  [0, 2, 2, 3]

Before: [2, 2, 2, 3]
6 2 2 0
After:  [4, 2, 2, 3]

Before: [0, 1, 3, 3]
8 2 2 2
After:  [0, 1, 9, 3]

Before: [2, 0, 3, 1]
12 3 2 0
After:  [3, 0, 3, 1]

Before: [0, 3, 2, 0]
4 1 2 2
After:  [0, 3, 1, 0]

Before: [2, 3, 3, 3]
4 1 0 3
After:  [2, 3, 3, 1]

Before: [2, 1, 3, 2]
8 1 0 0
After:  [2, 1, 3, 2]

Before: [3, 2, 1, 1]
2 1 0 3
After:  [3, 2, 1, 1]

Before: [0, 1, 1, 0]
7 1 0 2
After:  [0, 1, 1, 0]

Before: [2, 0, 2, 3]
6 1 3 1
After:  [2, 3, 2, 3]

Before: [1, 0, 1, 3]
5 1 0 0
After:  [1, 0, 1, 3]

Before: [1, 3, 2, 3]
4 1 2 2
After:  [1, 3, 1, 3]

Before: [0, 1, 0, 3]
7 1 0 2
After:  [0, 1, 1, 3]

Before: [2, 3, 0, 2]
3 1 3 3
After:  [2, 3, 0, 1]

Before: [3, 2, 3, 1]
2 1 0 1
After:  [3, 1, 3, 1]

Before: [1, 1, 2, 2]
14 1 0 1
After:  [1, 1, 2, 2]

Before: [2, 3, 2, 2]
3 1 3 2
After:  [2, 3, 1, 2]

Before: [1, 1, 2, 2]
14 1 0 3
After:  [1, 1, 2, 1]

Before: [3, 3, 2, 2]
3 1 3 0
After:  [1, 3, 2, 2]

Before: [1, 0, 3, 1]
5 1 0 1
After:  [1, 1, 3, 1]

Before: [1, 0, 1, 2]
5 1 0 0
After:  [1, 0, 1, 2]

Before: [3, 1, 0, 2]
10 1 3 1
After:  [3, 3, 0, 2]

Before: [1, 0, 3, 2]
10 0 3 1
After:  [1, 3, 3, 2]

Before: [3, 2, 0, 2]
2 1 0 1
After:  [3, 1, 0, 2]

Before: [0, 3, 2, 1]
4 1 2 2
After:  [0, 3, 1, 1]

Before: [2, 1, 2, 1]
6 2 2 3
After:  [2, 1, 2, 4]

Before: [1, 0, 0, 1]
5 1 0 1
After:  [1, 1, 0, 1]

Before: [0, 0, 3, 1]
13 0 1 1
After:  [0, 1, 3, 1]

Before: [0, 0, 0, 1]
13 0 1 2
After:  [0, 0, 1, 1]

Before: [0, 2, 2, 3]
8 3 1 2
After:  [0, 2, 6, 3]

Before: [1, 0, 3, 1]
5 1 0 2
After:  [1, 0, 1, 1]

Before: [3, 3, 1, 0]
8 0 0 2
After:  [3, 3, 9, 0]

Before: [0, 0, 2, 0]
13 0 1 1
After:  [0, 1, 2, 0]

Before: [2, 3, 1, 0]
4 1 0 1
After:  [2, 1, 1, 0]

Before: [2, 2, 2, 1]
12 3 2 2
After:  [2, 2, 3, 1]

Before: [0, 0, 2, 0]
13 0 1 0
After:  [1, 0, 2, 0]

Before: [1, 2, 3, 1]
12 3 2 0
After:  [3, 2, 3, 1]

Before: [2, 1, 1, 1]
8 3 0 3
After:  [2, 1, 1, 2]

Before: [1, 0, 0, 0]
5 1 0 3
After:  [1, 0, 0, 1]

Before: [1, 0, 0, 3]
15 0 3 2
After:  [1, 0, 0, 3]

Before: [0, 1, 2, 3]
6 2 2 2
After:  [0, 1, 4, 3]

Before: [0, 2, 3, 3]
11 0 0 1
After:  [0, 0, 3, 3]

Before: [0, 0, 1, 3]
13 0 1 0
After:  [1, 0, 1, 3]

Before: [0, 1, 2, 2]
7 1 0 2
After:  [0, 1, 1, 2]

Before: [1, 2, 1, 1]
1 2 1 1
After:  [1, 3, 1, 1]

Before: [3, 1, 2, 1]
12 3 2 1
After:  [3, 3, 2, 1]

Before: [1, 1, 2, 3]
14 1 0 3
After:  [1, 1, 2, 1]

Before: [0, 1, 2, 3]
7 1 0 2
After:  [0, 1, 1, 3]

Before: [2, 3, 3, 3]
4 1 0 2
After:  [2, 3, 1, 3]

Before: [1, 0, 2, 3]
15 0 3 1
After:  [1, 0, 2, 3]

Before: [0, 0, 1, 1]
11 0 0 1
After:  [0, 0, 1, 1]

Before: [1, 3, 1, 3]
0 3 3 3
After:  [1, 3, 1, 9]

Before: [0, 0, 3, 2]
13 0 1 0
After:  [1, 0, 3, 2]

Before: [1, 0, 2, 3]
0 2 3 0
After:  [6, 0, 2, 3]

Before: [0, 1, 1, 3]
11 0 0 2
After:  [0, 1, 0, 3]

Before: [0, 0, 2, 2]
13 0 1 0
After:  [1, 0, 2, 2]

Before: [2, 3, 3, 2]
3 1 3 1
After:  [2, 1, 3, 2]

Before: [3, 1, 1, 2]
12 3 1 2
After:  [3, 1, 3, 2]

Before: [0, 2, 3, 3]
0 1 3 1
After:  [0, 6, 3, 3]

Before: [3, 2, 1, 3]
2 1 0 1
After:  [3, 1, 1, 3]

Before: [0, 1, 2, 1]
6 1 2 0
After:  [3, 1, 2, 1]

Before: [3, 2, 1, 1]
2 1 0 1
After:  [3, 1, 1, 1]

Before: [1, 0, 1, 0]
5 1 0 3
After:  [1, 0, 1, 1]

Before: [3, 2, 0, 2]
9 3 1 0
After:  [4, 2, 0, 2]

Before: [3, 0, 3, 1]
1 1 2 3
After:  [3, 0, 3, 3]

Before: [2, 3, 2, 3]
4 1 2 2
After:  [2, 3, 1, 3]

Before: [3, 2, 2, 3]
2 1 0 2
After:  [3, 2, 1, 3]

Before: [0, 3, 0, 1]
1 2 1 1
After:  [0, 3, 0, 1]

Before: [1, 2, 0, 2]
1 1 0 2
After:  [1, 2, 3, 2]

Before: [0, 3, 2, 3]
4 1 2 2
After:  [0, 3, 1, 3]

Before: [3, 2, 3, 1]
2 1 0 3
After:  [3, 2, 3, 1]

Before: [3, 0, 2, 2]
1 1 0 1
After:  [3, 3, 2, 2]

Before: [0, 2, 1, 2]
1 1 2 0
After:  [3, 2, 1, 2]

Before: [0, 3, 2, 3]
4 1 2 1
After:  [0, 1, 2, 3]

Before: [1, 3, 2, 1]
8 1 1 1
After:  [1, 9, 2, 1]

Before: [2, 3, 1, 1]
4 1 0 3
After:  [2, 3, 1, 1]

Before: [0, 3, 3, 2]
3 1 3 3
After:  [0, 3, 3, 1]

Before: [0, 2, 1, 2]
9 3 1 1
After:  [0, 4, 1, 2]

Before: [0, 1, 1, 1]
7 1 0 2
After:  [0, 1, 1, 1]

Before: [0, 3, 1, 2]
3 1 3 2
After:  [0, 3, 1, 2]

Before: [3, 2, 1, 0]
2 1 0 0
After:  [1, 2, 1, 0]

Before: [3, 2, 3, 3]
2 1 0 3
After:  [3, 2, 3, 1]

Before: [2, 0, 2, 1]
12 3 2 3
After:  [2, 0, 2, 3]

Before: [1, 1, 3, 1]
12 3 2 3
After:  [1, 1, 3, 3]

Before: [2, 0, 3, 3]
1 1 2 0
After:  [3, 0, 3, 3]

Before: [0, 3, 3, 2]
8 3 1 2
After:  [0, 3, 6, 2]

Before: [0, 1, 3, 1]
12 3 2 0
After:  [3, 1, 3, 1]

Before: [0, 0, 1, 2]
13 0 1 0
After:  [1, 0, 1, 2]

Before: [1, 2, 3, 3]
8 2 2 2
After:  [1, 2, 9, 3]

Before: [0, 3, 2, 2]
9 3 3 3
After:  [0, 3, 2, 4]

Before: [1, 0, 2, 3]
15 0 3 0
After:  [0, 0, 2, 3]



12 1 1 0
12 3 2 1
12 0 1 2
6 0 1 1
0 1 2 1
9 1 3 3
12 2 2 2
12 3 2 0
12 1 3 1
1 2 0 1
0 1 1 1
9 3 1 3
7 3 3 0
12 0 0 1
12 1 0 3
12 0 1 2
6 3 1 1
0 1 1 1
9 0 1 0
12 3 3 2
12 1 2 1
0 1 2 1
0 1 3 1
0 1 3 1
9 0 1 0
12 2 0 3
12 2 3 2
12 3 3 1
2 2 1 1
0 1 3 1
9 0 1 0
7 0 2 1
12 0 3 2
12 3 2 3
12 0 2 0
11 3 2 3
0 3 2 3
9 3 1 1
12 3 1 2
12 0 3 3
12 2 0 0
3 0 2 0
0 0 2 0
9 1 0 1
12 1 2 0
12 0 2 2
12 3 2 2
0 2 3 2
0 2 2 2
9 1 2 1
7 1 1 2
12 2 2 3
12 2 0 1
12 2 3 0
10 1 3 1
0 1 3 1
9 2 1 2
7 2 3 3
12 1 3 1
12 0 0 2
0 1 2 2
0 2 2 2
9 3 2 3
7 3 0 2
12 1 2 3
12 0 0 1
12 1 1 0
6 0 1 3
0 3 1 3
9 3 2 2
7 2 1 1
12 2 2 0
12 1 2 2
12 1 2 3
4 0 3 0
0 0 1 0
9 0 1 1
7 1 0 0
0 1 0 3
6 3 0 3
12 2 0 1
12 2 3 2
15 3 2 1
0 1 2 1
0 1 3 1
9 1 0 0
7 0 3 1
12 3 2 2
12 1 2 0
5 3 2 0
0 0 3 0
9 0 1 1
12 0 0 0
12 2 3 0
0 0 2 0
0 0 2 0
9 1 0 1
7 1 2 2
12 1 2 3
0 3 0 1
6 1 1 1
12 2 0 0
9 3 3 0
0 0 3 0
0 0 3 0
9 0 2 2
7 2 2 0
12 0 2 1
0 2 0 3
6 3 0 3
0 2 0 2
6 2 3 2
5 3 2 2
0 2 1 2
9 2 0 0
7 0 3 3
0 2 0 2
6 2 0 2
12 3 1 1
12 2 2 0
14 1 0 1
0 1 1 1
9 1 3 3
7 3 2 0
12 1 2 3
12 3 2 2
12 3 3 1
9 3 3 1
0 1 3 1
9 1 0 0
7 0 1 1
12 2 1 0
12 0 0 2
12 0 3 3
12 2 3 2
0 2 2 2
9 1 2 1
12 2 1 3
12 1 2 0
12 2 1 2
10 2 3 2
0 2 1 2
0 2 2 2
9 1 2 1
7 1 2 0
12 1 2 2
12 2 1 1
12 3 2 3
14 3 1 3
0 3 1 3
0 3 3 3
9 3 0 0
7 0 1 3
12 0 2 2
12 0 2 0
12 3 0 1
11 1 2 0
0 0 2 0
0 0 1 0
9 0 3 3
7 3 2 1
0 3 0 0
6 0 2 0
12 0 3 3
12 3 3 2
1 0 2 0
0 0 1 0
9 0 1 1
7 1 3 2
12 1 0 1
12 3 1 3
12 2 0 0
8 1 0 0
0 0 1 0
9 0 2 2
7 2 1 1
12 1 2 2
12 1 2 0
0 0 0 3
6 3 2 3
8 0 3 2
0 2 3 2
9 2 1 1
7 1 1 3
12 3 2 1
12 2 3 2
6 0 1 0
0 0 1 0
0 0 3 0
9 0 3 3
7 3 1 0
12 0 1 1
12 2 0 3
12 3 0 2
12 1 2 3
0 3 3 3
9 3 0 0
7 0 1 1
0 2 0 3
6 3 3 3
12 1 3 0
12 0 2 2
0 0 2 0
0 0 2 0
9 0 1 1
12 1 3 3
12 2 1 0
0 0 0 2
6 2 1 2
4 0 3 3
0 3 1 3
9 1 3 1
7 1 2 2
12 3 3 1
12 0 1 0
12 1 2 3
6 3 1 0
0 0 1 0
9 0 2 2
7 2 2 0
12 2 3 1
0 1 0 2
6 2 3 2
1 1 2 1
0 1 1 1
9 0 1 0
0 3 0 2
6 2 2 2
12 3 2 1
12 0 2 3
2 2 1 3
0 3 3 3
9 0 3 0
12 1 2 3
12 0 2 1
12 3 0 2
0 3 2 1
0 1 2 1
9 0 1 0
7 0 2 2
12 3 1 1
12 2 2 0
4 0 3 1
0 1 2 1
0 1 3 1
9 1 2 2
12 2 3 1
0 2 0 3
6 3 3 3
14 3 1 0
0 0 3 0
9 0 2 2
7 2 0 3
12 2 2 2
12 3 3 0
2 2 0 0
0 0 1 0
9 0 3 3
7 3 2 1
12 2 2 0
12 1 3 3
4 0 3 2
0 2 2 2
0 2 1 2
9 1 2 1
12 2 2 2
12 1 0 0
12 0 2 3
7 0 2 2
0 2 2 2
9 1 2 1
12 3 3 2
12 1 3 3
12 2 2 0
4 0 3 0
0 0 1 0
9 0 1 1
0 1 0 0
6 0 1 0
9 0 0 3
0 3 3 3
9 1 3 1
12 1 3 3
12 2 1 0
12 1 2 2
4 0 3 2
0 2 3 2
9 1 2 1
7 1 1 3
12 3 2 1
12 0 2 2
12 3 3 0
3 2 0 0
0 0 2 0
9 3 0 3
7 3 0 0
0 1 0 1
6 1 0 1
12 3 3 3
12 1 2 2
11 3 2 1
0 1 1 1
9 0 1 0
7 0 3 2
0 0 0 0
6 0 2 0
12 3 1 1
14 1 0 1
0 1 2 1
9 2 1 2
7 2 1 1
12 1 3 3
12 3 1 0
12 3 0 2
11 0 2 0
0 0 3 0
9 0 1 1
7 1 3 3
12 2 0 0
0 1 0 1
6 1 0 1
1 0 2 0
0 0 1 0
9 0 3 3
7 3 1 2
12 1 1 0
12 3 1 1
12 2 0 3
8 0 3 3
0 3 2 3
9 3 2 2
7 2 1 1
12 0 1 2
12 2 1 0
0 0 0 3
6 3 1 3
0 3 2 0
0 0 1 0
9 0 1 1
7 1 3 3
12 2 0 2
0 0 0 1
6 1 0 1
12 1 3 0
7 0 2 2
0 2 1 2
9 3 2 3
7 3 0 0
0 0 0 2
6 2 3 2
0 2 0 3
6 3 3 3
12 2 3 1
1 1 2 2
0 2 2 2
9 0 2 0
7 0 0 2
12 1 3 0
12 0 0 1
6 0 1 1
0 1 1 1
9 1 2 2
7 2 3 0
12 0 2 3
0 3 0 2
6 2 2 2
12 1 0 1
15 3 2 3
0 3 3 3
9 3 0 0
7 0 0 1
12 3 0 0
12 3 3 3
2 2 0 3
0 3 3 3
9 1 3 1
7 1 1 0
12 3 0 1
12 0 2 3
15 3 2 2
0 2 3 2
9 2 0 0
7 0 2 2
12 0 0 1
12 2 1 0
12 1 3 3
4 0 3 1
0 1 1 1
0 1 3 1
9 2 1 2
7 2 3 1
0 1 0 3
6 3 3 3
12 0 2 2
11 3 2 3
0 3 1 3
0 3 1 3
9 3 1 1
7 1 3 3
0 3 0 2
6 2 2 2
12 3 0 0
12 0 3 1
1 2 0 2
0 2 1 2
0 2 2 2
9 2 3 3
12 0 0 2
12 1 1 0
12 3 3 1
0 0 2 1
0 1 1 1
9 1 3 3
12 1 2 2
0 1 0 1
6 1 1 1
0 3 0 0
6 0 3 0
11 0 2 2
0 2 1 2
9 3 2 3
12 3 3 1
12 1 1 0
12 1 1 2
6 0 1 0
0 0 3 0
0 0 1 0
9 0 3 3
7 3 0 1
12 2 2 2
12 1 2 0
0 1 0 3
6 3 0 3
15 3 2 2
0 2 1 2
9 1 2 1
12 1 1 2
12 2 1 3
0 0 0 0
6 0 2 0
13 0 3 2
0 2 2 2
9 1 2 1
0 3 0 3
6 3 0 3
12 2 3 2
15 3 2 0
0 0 3 0
9 1 0 1
7 1 1 0
0 2 0 1
6 1 2 1
12 2 2 3
0 2 0 2
6 2 0 2
5 2 3 3
0 3 1 3
9 0 3 0
7 0 1 1
12 2 3 0
12 2 2 3
13 0 3 3
0 3 1 3
0 3 1 3
9 1 3 1
7 1 2 0
12 3 3 3
12 2 3 2
12 3 2 1
2 2 1 3
0 3 2 3
9 0 3 0
7 0 0 3
12 1 1 0
7 0 2 0
0 0 1 0
9 3 0 3
7 3 0 1
12 0 1 2
12 3 0 0
12 0 2 3
3 2 0 3
0 3 2 3
0 3 1 3
9 1 3 1
7 1 1 0
12 2 3 1
0 0 0 3
6 3 0 3
12 3 0 2
1 1 2 1
0 1 2 1
9 1 0 0
7 0 1 3
12 0 1 1
0 0 0 2
6 2 1 2
12 3 1 0
11 0 2 2
0 2 3 2
9 2 3 3
7 3 0 0
12 2 2 3
12 1 0 2
0 0 0 1
6 1 3 1
14 1 3 3
0 3 1 3
9 3 0 0
0 1 0 3
6 3 2 3
12 1 3 1
12 0 1 2
0 1 2 1
0 1 3 1
0 1 2 1
9 1 0 0
7 0 1 2
0 2 0 1
6 1 3 1
0 3 0 0
6 0 3 0
14 1 3 3
0 3 2 3
9 3 2 2
12 1 0 3
0 0 0 0
6 0 2 0
12 2 2 1
4 0 3 1
0 1 3 1
9 2 1 2
7 2 1 3
12 3 2 2
12 1 1 1
3 0 2 1
0 1 3 1
9 3 1 3
7 3 0 1
12 1 0 3
12 0 0 2
12 3 3 0
0 3 2 0
0 0 3 0
0 0 1 0
9 0 1 1
12 0 3 3
0 0 0 2
6 2 2 2
12 0 1 0
15 3 2 2
0 2 1 2
9 1 2 1
7 1 0 3
12 1 2 0
12 1 2 1
12 0 0 2
12 2 1 2
0 2 3 2
9 2 3 3
7 3 0 2
0 3 0 0
6 0 2 0
12 2 0 1
0 2 0 3
6 3 2 3
13 0 3 0
0 0 1 0
9 0 2 2
7 2 1 3
12 3 3 2
12 0 3 1
12 2 1 0
3 0 2 2
0 2 3 2
0 2 1 2
9 3 2 3
7 3 2 2
12 3 3 1
12 1 0 3
4 0 3 3
0 3 1 3
9 2 3 2
12 2 3 3
12 0 2 1
13 0 3 1
0 1 3 1
0 1 1 1
9 1 2 2
7 2 3 0
12 2 1 1
0 2 0 2
6 2 2 2
0 2 0 3
6 3 0 3
15 3 2 2
0 2 1 2
9 2 0 0
7 0 3 1
12 2 1 0
12 1 3 3
12 3 3 2
4 0 3 2
0 2 3 2
9 1 2 1
7 1 2 0
12 3 2 3
12 3 2 1
0 0 0 2
6 2 0 2
11 3 2 1
0 1 1 1
0 1 1 1
9 0 1 0
7 0 0 1
12 1 1 2
12 0 1 0
12 0 0 3
12 2 3 2
0 2 2 2
9 2 1 1
0 2 0 2
6 2 2 2
12 3 1 0
15 3 2 2
0 2 3 2
9 1 2 1
12 3 2 3
12 0 2 2
3 2 0 3
0 3 2 3
9 1 3 1
7 1 1 2
0 3 0 1
6 1 3 1
12 2 3 3
0 1 0 0
6 0 2 0
13 0 3 0
0 0 3 0
9 0 2 2
12 1 0 3
12 3 1 0
6 3 1 3
0 3 2 3
9 2 3 2
7 2 0 3
12 2 3 2
12 1 0 1
12 2 1 0
8 1 0 0
0 0 3 0
9 3 0 3
12 3 3 1
12 2 3 0
2 2 1 2
0 2 1 2
0 2 1 2
9 3 2 3
7 3 1 0
12 3 2 2
12 0 2 3
12 0 1 1
5 3 2 1
0 1 3 1
9 0 1 0
7 0 0 1
12 2 2 3
12 2 0 2
12 2 0 0
13 0 3 2
0 2 3 2
9 1 2 1
7 1 3 0
12 3 2 3
0 2 0 2
6 2 0 2
12 1 2 1
11 3 2 1
0 1 2 1
9 1 0 0
7 0 2 1
12 2 1 0
12 2 0 3
5 2 3 2
0 2 1 2
9 2 1 1
7 1 0 0
12 3 1 1
12 3 1 2
12 1 3 3
12 2 1 3
0 3 1 3
9 0 3 0
7 0 2 2
0 0 0 0
6 0 1 0
12 1 1 3
9 0 3 0
0 0 1 0
9 2 0 2
7 2 0 3
12 3 1 2
12 2 0 1
12 3 1 0
1 1 0 2
0 2 3 2
9 2 3 3
12 1 3 0
12 2 2 2
7 0 2 2
0 2 1 2
0 2 2 2
9 3 2 3
12 3 2 2
9 0 0 0
0 0 2 0
9 0 3 3
7 3 1 1
0 1 0 2
6 2 2 2
0 0 0 3
6 3 2 3
12 3 0 0
1 2 0 0
0 0 3 0
9 0 1 1
12 3 2 3
12 3 3 0
12 1 2 2
11 0 2 0
0 0 2 0
0 0 2 0
9 0 1 1
12 1 1 0
12 1 3 3
12 2 3 2
7 0 2 0
0 0 3 0
0 0 2 0
9 1 0 1
7 1 2 2
12 3 2 1
12 3 1 0
12 2 0 3
14 1 3 0
0 0 2 0
0 0 2 0
9 0 2 2
7 2 3 0
12 3 1 3
12 0 1 2
12 2 0 1
12 2 3 2
0 2 1 2
9 0 2 0
7 0 0 2
12 3 2 1
12 1 3 0
6 0 1 1
0 1 2 1
9 1 2 2
7 2 0 3
12 3 1 2
12 2 3 1
12 3 0 0
1 1 2 2
0 2 1 2
0 2 3 2
9 2 3 3
12 1 1 1
0 0 0 2
6 2 3 2
12 0 3 0
0 1 2 2
0 2 3 2
0 2 3 2
9 2 3 3
12 0 2 2
12 3 0 1
0 1 0 0
6 0 3 0
3 2 0 2
0 2 1 2
0 2 1 2
9 3 2 3
7 3 0 2
12 1 2 0
12 1 1 3
9 0 0 1
0 1 3 1
0 1 2 1
9 1 2 2
7 2 3 0
12 3 0 2
0 3 0 3
6 3 0 3
0 3 0 1
6 1 1 1
5 3 2 3
0 3 1 3
0 3 3 3
9 3 0 0
7 0 3 3
12 3 2 1
12 2 3 2
12 0 2 0
2 2 1 0
0 0 1 0
9 0 3 3
7 3 3 1
12 0 0 3
12 3 2 0
1 2 0 0
0 0 2 0
0 0 3 0
9 0 1 1
7 1 1 3
12 3 1 0
12 2 0 1
2 2 0 1
0 1 2 1
9 3 1 3
7 3 2 0
12 3 3 2
12 1 2 3
0 3 0 1
6 1 1 1
0 1 2 3
0 3 3 3
9 0 3 0
7 0 0 2
12 2 2 0
12 0 0 1
0 2 0 3
6 3 1 3
4 0 3 0
0 0 3 0
0 0 3 0
9 0 2 2
7 2 0 0
12 3 3 1
0 3 0 2
6 2 1 2
6 3 1 2
0 2 3 2
0 2 3 2
9 2 0 0
7 0 0 1
12 0 2 3
12 3 3 0
12 2 3 2
15 3 2 3
0 3 2 3
9 1 3 1
7 1 1 0
12 2 2 1
12 3 0 2
12 1 0 3
0 3 2 1
0 1 2 1
9 1 0 0
7 0 0 2
12 2 0 1
12 2 0 3
0 0 0 0
6 0 1 0
10 1 3 0
0 0 3 0
0 0 3 0
9 0 2 2
12 3 3 0
12 1 0 3
12 0 2 1
6 3 1 0
0 0 3 0
0 0 3 0
9 2 0 2
7 2 0 1
12 1 1 0
12 2 2 2
12 2 3 3
10 2 3 3
0 3 1 3
9 1 3 1
7 1 3 0
12 2 1 3
12 0 0 1
0 0 0 2
6 2 0 2
5 2 3 2
0 2 3 2
0 2 2 2
9 0 2 0
7 0 1 2
12 1 3 0
12 1 3 1
8 0 3 3
0 3 1 3
9 3 2 2
7 2 2 3
12 2 0 1
12 3 2 0
0 1 0 2
6 2 1 2
11 0 2 1
0 1 1 1
9 1 3 3
7 3 3 0
12 1 1 1
0 3 0 3
6 3 2 3
0 2 0 2
6 2 0 2
5 2 3 3
0 3 1 3
0 3 3 3
9 3 0 0
7 0 1 3
0 0 0 0
6 0 3 0
12 2 1 1
12 3 3 2
1 1 0 1
0 1 3 1
0 1 2 1
9 1 3 3
7 3 2 1
12 1 2 0
12 0 2 3
0 1 0 2
6 2 2 2
15 3 2 0
0 0 1 0
9 0 1 1
7 1 1 2
0 1 0 0
6 0 1 0
12 3 2 1
6 0 1 1
0 1 3 1
0 1 1 1
9 2 1 2
12 2 0 0
0 2 0 1
6 1 3 1
0 1 0 3
6 3 3 3
2 0 1 0
0 0 1 0
9 0 2 2
12 3 2 0
12 0 2 3
12 2 3 1
10 1 3 0
0 0 3 0
9 0 2 2
7 2 3 1
12 1 2 0
12 2 0 3
12 3 2 2
0 0 2 2
0 2 2 2
0 2 1 2
9 1 2 1
7 1 3 0`
