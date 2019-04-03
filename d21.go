package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")
	p1(INPUT_TAEL)
	// p2()
	p2alt(INPUT_TAEL) // 11316540
	p2alt(INPUT)      // 9547924

	// sim()
}

func sim(number int) int {

	var cache [16777216]bool

	var r1, r3, r4, r5 int

	r3 = r5 | 65536
	r5 = number

	previous := r5
	r5 += r3 & 255
	r5 &= 16777215
	r5 *= 65899
	r5 &= 16777215

	for {
		if r3 < 256 {
			if !cache[r5] {
				cache[r5] = true
				previous = r5
			} else {
				break
			}

			r3 = r5 | 65536
			r5 = number

			r5 += r3 & 255
			r5 &= 16777215
			r5 *= 65899
			r5 &= 16777215

		} else {

			r1, r4 = r3/256-1, 256
			for r4 <= r3 {
				r1++
				r4 = (r1 + 1) * 256
			}

			r3 = r1

			r5 += r3 & 255
			r5 &= 16777215
			r5 *= 65899
			r5 &= 16777215
		}
	}

	return previous
}

func p1(input string) {

	t := time.Now()

	var program []Inst
	ip := 0

	for _, s := range strings.Split(input, "\n") {

		if len(s) == 0 {
			continue
		}

		if s[0] == '#' {
			ip, _ = strconv.Atoi(strings.Split(s, " ")[1])
			continue
		}

		var instruction Inst
		scanned, err := fmt.Sscanf(s, "%s %d %d %d", &instruction.Name, &instruction.Op[0], &instruction.Op[1], &instruction.Op[2])
		if err != nil || scanned != 4 {
			panic("Couldn't scan! " + s)
		}

		program = append(program, instruction)
	}

	register := make([]int, 6)
	register[0] = 0
	ptr := register[ip]

	var result int

	for ptr >= 0 && ptr < len(program) {

		code := program[ptr]

		if code.Name == "eqrr" {
			result = register[code.Op[0]]
			// fmt.Println(result)
			break
		}

		ops[code.Name](code.Op, register)
		register[ip]++
		ptr = register[ip]
	}

	fmt.Println("Done:", result, time.Since(t))
}

func p2(input string) {

	t := time.Now()

	var program []Inst
	ip := 0

	for _, s := range strings.Split(input, "\n") {

		if len(s) == 0 {
			continue
		}

		if s[0] == '#' {
			ip, _ = strconv.Atoi(strings.Split(s, " ")[1])
			continue
		}

		var instruction Inst
		scanned, err := fmt.Sscanf(s, "%s %d %d %d", &instruction.Name, &instruction.Op[0], &instruction.Op[1], &instruction.Op[2])
		if err != nil || scanned != 4 {
			panic("Couldn't scan! " + s)
		}

		program = append(program, instruction)
	}

	register := make([]int, 6)
	register[0] = 0
	ptr := register[ip]

	seen := make(map[int]bool)

	count := 1

	previous := -1

	for ptr >= 0 && ptr < len(program) {

		code := program[ptr]

		if code.Name == "eqrr" {
			if !seen[register[code.Op[0]]] {
				seen[register[code.Op[0]]] = true
				previous = register[code.Op[0]]
			} else {
				break
			}
		}

		ops[code.Name](code.Op, register)
		register[ip]++
		ptr = register[ip]

		count++
	}

	fmt.Println("Done 2:", previous, time.Since(t))
}

func p2alt(input string) {

	t := time.Now()

	var program []Inst
	ip := 0

	for _, s := range strings.Split(input, "\n") {

		if len(s) == 0 {
			continue
		}

		if s[0] == '#' {
			ip, _ = strconv.Atoi(strings.Split(s, " ")[1])
			continue
		}

		var instruction Inst
		scanned, err := fmt.Sscanf(s, "%s %d %d %d", &instruction.Name, &instruction.Op[0], &instruction.Op[1], &instruction.Op[2])
		if err != nil || scanned != 4 {
			panic("Couldn't scan! " + s)
		}

		program = append(program, instruction)
	}

	register := make([]int, 6)
	register[0] = 0
	ptr := register[ip]

	result := -1

	for ptr >= 0 && ptr < len(program) {

		code := program[ptr]

		if code.Name == "seti" {
			if code.Op[0] > 1000 {
				result = sim(code.Op[0])
				break
			}
		}

		ops[code.Name](code.Op, register)
		register[ip]++
		ptr = register[ip]
	}

	fmt.Println("Done 2:", result, time.Since(t))
}

type Inst struct {
	Name string
	Op   [3]int
}

func addr(op [3]int, r []int) {
	r[op[2]] = r[op[0]] + r[op[1]]
}

func addi(op [3]int, r []int) {
	r[op[2]] = r[op[0]] + op[1]
}

func mulr(op [3]int, r []int) {
	r[op[2]] = r[op[0]] * r[op[1]]
}

func muli(op [3]int, r []int) {
	r[op[2]] = r[op[0]] * op[1]
}

func banr(op [3]int, r []int) {
	r[op[2]] = r[op[0]] & r[op[1]]
}

func bani(op [3]int, r []int) {
	r[op[2]] = r[op[0]] & op[1]
}

func borr(op [3]int, r []int) {
	r[op[2]] = r[op[0]] | r[op[1]]
}

func bori(op [3]int, r []int) {
	r[op[2]] = r[op[0]] | op[1]
}

func setr(op [3]int, r []int) {
	r[op[2]] = r[op[0]]
}

func seti(op [3]int, r []int) {
	r[op[2]] = op[0]
}

func gtir(op [3]int, r []int) {
	if op[0] > r[op[1]] {
		r[op[2]] = 1
	} else {
		r[op[2]] = 0
	}
}

func gtri(op [3]int, r []int) {
	if r[op[0]] > op[1] {
		r[op[2]] = 1
	} else {
		r[op[2]] = 0
	}
}

func gtrr(op [3]int, r []int) {
	if r[op[0]] > r[op[1]] {
		r[op[2]] = 1
	} else {
		r[op[2]] = 0
	}
}

func eqir(op [3]int, r []int) {
	if op[0] == r[op[1]] {
		r[op[2]] = 1
	} else {
		r[op[2]] = 0
	}
}

func eqri(op [3]int, r []int) {
	if r[op[0]] == op[1] {
		r[op[2]] = 1
	} else {
		r[op[2]] = 0
	}
}

func eqrr(op [3]int, r []int) {
	if r[op[0]] == r[op[1]] {
		r[op[2]] = 1
	} else {
		r[op[2]] = 0
	}
}

var ops map[string]func([3]int, []int)

func init() {

	ops = map[string]func([3]int, []int){
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

const INPUT1 = ``

const INPUT = `#ip 2
seti 123 0 5
bani 5 456 5
eqri 5 72 5
addr 5 2 2
seti 0 0 2
seti 0 3 5
bori 5 65536 3
seti 9010242 6 5
bani 3 255 1
addr 5 1 5
bani 5 16777215 5
muli 5 65899 5
bani 5 16777215 5
gtir 256 3 1
addr 1 2 2
addi 2 1 2
seti 27 6 2
seti 0 8 1
addi 1 1 4
muli 4 256 4
gtrr 4 3 4
addr 4 2 2
addi 2 1 2
seti 25 5 2
addi 1 1 1
seti 17 7 2
setr 1 3 3
seti 7 2 2
eqrr 5 0 1
addr 1 2 2
seti 5 2 2`

const INPUT_TAEL = `#ip 4
seti 123 0 1
bani 1 456 1
eqri 1 72 1
addr 1 4 4
seti 0 0 4
seti 0 6 1
bori 1 65536 3
seti 6780005 8 1
bani 3 255 2
addr 1 2 1
bani 1 16777215 1
muli 1 65899 1
bani 1 16777215 1
gtir 256 3 2
addr 2 4 4
addi 4 1 4
seti 27 5 4
seti 0 5 2
addi 2 1 5
muli 5 256 5
gtrr 5 3 5
addr 5 4 4
addi 4 1 4
seti 25 4 4
addi 2 1 2
seti 17 7 4
setr 2 1 3
seti 7 3 4
eqrr 1 0 2
addr 2 4 4
seti 5 4 4`
