package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")
	p1()
	// p2()
	p2alt()
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

func p1() {

	t := time.Now()

	var program []Inst
	ip := 0

	for _, s := range strings.Split(INPUT_TAEL, "\n") {

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
	ptr := register[ip]

	for ptr >= 0 && ptr < len(program) {

		code := program[ptr]
		ops[code.Name](code.Op, register)
		register[ip]++
		ptr = register[ip]

		// fmt.Println(register, ptr)
		// var input string
		// fmt.Scanln(&input)

	}

	fmt.Println("Done:", register, time.Since(t))
}

func p2() {

	t := time.Now()

	var program []Inst
	ip := 0

	for _, s := range strings.Split(INPUT, "\n") {

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
	register[0] = 1
	ptr := register[ip]

	hijack := 0
	for ptr >= 0 && ptr < len(program) {

		code := program[ptr]

		if code == (Inst{"addi", [3]int{2, 1, 2}}) {

			code = Inst{"setr", [3]int{1, 0, 2}}
			ops[code.Name](code.Op, register)

			if hijack%2 == 0 && register[1] != register[4] {
				register[2] /= register[4]
			} else {
				register[2]++
			}

			hijack++
		} else {
			ops[code.Name](code.Op, register)
		}

		// if register[0] >= 18869760 || code == (Inst{"addr", [3]int{4, 0, 0}}) {
		// fmt.Println(code, register, register[ip]+1)
		// var input string
		// fmt.Scanln(&input)
		// }

		register[ip]++
		ptr = register[ip]
	}

	// Sum of factors of big number in register needed
	fmt.Println("Done 2:", register, time.Since(t))
}

func p2alt() {

	t := time.Now()

	var program []Inst
	ip := 0

	for _, s := range strings.Split(INPUT_TAEL, "\n") {

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
	register[0] = 1
	ptr := register[ip]

	result := 0

	for ptr >= 0 && ptr < len(program) {

		code := program[ptr]

		if code.Name == "eqrr" {
			result = divisorSum(register[code.Op[1]])
			break
		}

		ops[code.Name](code.Op, register)

		register[ip]++
		ptr = register[ip]
	}

	// Sum of divisors of big number in register needed
	fmt.Println("Done 2alt:", result, time.Since(t))
}

func divisorSum(n int) int {

	sum := 1 + n

	i := 0
	for i = 2; i*i < n; i++ {
		if n%i == 0 {
			sum += i + n/i
		}
	}

	if i*i == n {
		sum += i
	}

	return sum
}

const INPUT1 = `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`

const INPUT = `#ip 5
addi 5 16 5
seti 1 1 4
seti 1 8 2
mulr 4 2 3
eqrr 3 1 3
addr 3 5 5
addi 5 1 5
addr 4 0 0
addi 2 1 2
gtrr 2 1 3
addr 5 3 5
seti 2 6 5
addi 4 1 4
gtrr 4 1 3
addr 3 5 5
seti 1 4 5
mulr 5 5 5
addi 1 2 1
mulr 1 1 1
mulr 5 1 1
muli 1 11 1
addi 3 7 3
mulr 3 5 3
addi 3 8 3
addr 1 3 1
addr 5 0 5
seti 0 9 5
setr 5 8 3
mulr 3 5 3
addr 5 3 3
mulr 5 3 3
muli 3 14 3
mulr 3 5 3
addr 1 3 1
seti 0 4 0
seti 0 3 5`

const INPUT_TAEL = `#ip 4
addi 4 16 4
seti 1 5 1
seti 1 2 2
mulr 1 2 3
eqrr 3 5 3
addr 3 4 4
addi 4 1 4
addr 1 0 0
addi 2 1 2
gtrr 2 5 3
addr 4 3 4
seti 2 7 4
addi 1 1 1
gtrr 1 5 3
addr 3 4 4
seti 1 9 4
mulr 4 4 4
addi 5 2 5
mulr 5 5 5
mulr 4 5 5
muli 5 11 5
addi 3 1 3
mulr 3 4 3
addi 3 18 3
addr 5 3 5
addr 4 0 4
seti 0 3 4
setr 4 2 3
mulr 3 4 3
addr 4 3 3
mulr 4 3 3
muli 3 14 3
mulr 3 4 3
addr 5 3 5
seti 0 4 0
seti 0 5 4`
