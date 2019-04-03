package main

import (
	"bytes"
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

	target := 540561

	list := []byte{3, 7}
	elfA := 0
	elfB := 1

	for len(list) < 10+target {
		sum := list[elfA] + list[elfB]

		if sum >= 10 {
			list = append(list, sum/10)
		}
		list = append(list, sum%10)

		elfA = (elfA + 1 + int(list[elfA])) % len(list)
		elfB = (elfB + 1 + int(list[elfB])) % len(list)

		// fmt.Println(elfA, elfB, list)
	}

	var buf bytes.Buffer
	for i := 0; i < 10; i++ {
		buf.WriteByte(list[target+i] + '0')
	}

	fmt.Println("Done:", buf.String(), time.Since(t))
}

func p2() {

	t := time.Now()

	// target := []byte{5, 1, 5, 8, 9}
	// target := []byte{0, 1, 2, 4, 5}
	// target := []byte{9, 2, 5, 1, 0}
	// target := []byte{5, 9, 4, 1, 4}
	target := []byte{5, 4, 0, 5, 6, 1}

	// target := []byte{9, 3, 9, 6, 0, 1}

	list := []byte{3, 7}
	elfA := 0
	elfB := 1

	position := 0

	for {
		sum := list[elfA] + list[elfB]

		if sum >= 10 {
			list = append(list, sum/10)
		}
		list = append(list, sum%10)

		elfA = (elfA + 1 + int(list[elfA])) % len(list)
		elfB = (elfB + 1 + int(list[elfB])) % len(list)

		startPoint := len(list) - len(target) - 1
		if startPoint < 0 {
			startPoint = 0
		}

		position = bytes.Index(list[startPoint:], target)
		if position != -1 {
			position = startPoint + position
			break
		}
		// fmt.Println(elfA, elfB, list)
	}

	fmt.Println("Done 2:", position, time.Since(t))
}
