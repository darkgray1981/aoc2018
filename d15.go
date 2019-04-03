package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")

	// verifyP1()
	// verifyP2()

	p1()
	p2()

	// test()
}

func verifyP1() {

	t := time.Now()

	if verifyP1Run(INPUT1) != 27730 {
		panic("1")
	}

	if verifyP1Run(INPUT2) != 36334 {
		panic("2")
	}
	if verifyP1Run(INPUT3) != 39514 {
		panic("3")
	}
	if verifyP1Run(INPUT4) != 27755 {
		panic("4")
	}
	if verifyP1Run(INPUT5) != 28944 {
		panic("5")
	}
	if verifyP1Run(INPUT6) != 18740 {
		panic("6")
	}

	if verifyP1Run(INPUT) != 188576 {
		panic("REAL")
	}

	if verifyP1Run(INPUT_TAEL) != 257954 {
		panic("TAEL")
	}
	if verifyP1Run(INPUT_REDDIT1) != 215168 {
		panic("REDDIT1")
	}

	if r := verifyP1Run(INPUT_REDDIT3); r != 229798 {
		panic(fmt.Sprintf("REDDIT3: %d != %d", r, 229798))
	}

	fmt.Println("All P1 tests passed in", time.Since(t))
}

func verifyP2() {

	t := time.Now()

	if r := verifyP2Run(INPUT1); r != 4988 {
		panic(fmt.Sprintf("1: %d != %d", r, 4988))
	}

	if r := verifyP2Run(INPUT3); r != 31284 {
		panic(fmt.Sprintf("3: %d != %d", r, 31284))
	}
	if r := verifyP2Run(INPUT4); r != 3478 {
		panic(fmt.Sprintf("3: %d != %d", r, 3478))
	}
	if r := verifyP2Run(INPUT5); r != 6474 {
		panic(fmt.Sprintf("3: %d != %d", r, 6474))
	}
	if r := verifyP2Run(INPUT6); r != 1140 {
		panic(fmt.Sprintf("6: %d != %d", r, 1140))
	}

	if r := verifyP2Run(INPUT); r != 57112 {
		panic(fmt.Sprintf("REAL: %d != %d", r, 57112))
	}

	if r := verifyP2Run(INPUT_TAEL); r != 51041 {
		panic(fmt.Sprintf("TAEL: %d != %d", r, 51041))
	}
	if r := verifyP2Run(INPUT_REDDIT1); r != 52374 {
		panic(fmt.Sprintf("REDDIT1: %d != %d", r, 52374))
	}
	if r := verifyP2Run(INPUT_REDDIT2); r != 67022 {
		panic(fmt.Sprintf("REDDIT2: %d != %d", r, 67022))
	}

	fmt.Println("All P2 tests passed in", time.Since(t))
}

func verifyP1Run(input string) int {
	field, units := newField(input)
	moved := true
	round := -1

outer:
	for moved {
		round++

		// fmt.Println()

		bodies := values(units)
		sortLocation(bodies)

		moved = false
		for _, b := range bodies {
			if !b.alive {
				continue
			}

			elves := countTeam('E', units)
			if elves == 0 || elves == len(units) {
				break outer
			}

			// fmt.Println("Moving", b.team, b.pos.x, b.pos.y)
			moved = b.Move(field, units) || moved
			b.Attack(field, units)
		}
	}

	hpsum := 0
	for _, u := range units {
		hpsum += u.hp
	}

	return round * hpsum
}

func verifyP2Run(input string) int {

	var round, hpsum int

	for apbuff := 4; hpsum == 0; apbuff++ {
		field, units := newField(input)
		moved := true
		round = -1

		everyone := 0
		for _, u := range units {
			if u.team == 'E' {
				u.ap = apbuff
				everyone++
			}
		}

	outer:
		for moved {
			round++

			// fmt.Println()

			bodies := values(units)
			sortLocation(bodies)

			moved = false
			for _, b := range bodies {
				if !b.alive {
					continue
				}

				elves := countTeam('E', units)
				if elves != everyone || elves == len(units) {
					break outer
				}

				// fmt.Println("Moving", b.team, b.pos.x, b.pos.y)
				moved = b.Move(field, units) || moved
				b.Attack(field, units)
			}
		}

		if countTeam('E', units) != everyone {
			continue
		}

		hpsum = 0
		for _, u := range units {
			hpsum += u.hp
		}
	}

	return round * hpsum
}

func p1() {

	t := time.Now()

	field, units := newField(INPUT)

	fmt.Println("Initially:")
	draw(field, units)
	fmt.Println()

	moved := true
	round := -1

outer:
	for moved {
		round++

		// fmt.Println()

		bodies := values(units)
		sortLocation(bodies)

		moved = false
		for _, b := range bodies {
			if !b.alive {
				continue
			}

			elves := countTeam('E', units)
			if elves == 0 || elves == len(units) {
				break outer
			}

			// fmt.Println("Moving", b.team, b.pos.x, b.pos.y)
			moved = b.Move(field, units) || moved
			b.Attack(field, units)
		}

		// fmt.Println("\nAfter", round+1, "round(s):")
		// draw(field, units)
		// time.Sleep(200 * time.Millisecond)
		// var input string
		// fmt.Scanln(&input)
	}

	fmt.Println("\nAfter", round+1, "round(s):")
	draw(field, units)

	hpsum := 0
	for _, u := range units {
		hpsum += u.hp
	}

	fmt.Println("Done:", round, round*hpsum, time.Since(t))
}

func p2() {

	t := time.Now()

	r := verifyP2Run(INPUT)

	fmt.Println("Done 2:", r, time.Since(t))
}

type Pos struct {
	x int
	y int
}

type Unit struct {
	pos   Pos
	hp    int
	ap    int
	team  byte
	alive bool
}

// Calculate the number of living units on a specified team
func countTeam(team byte, units map[Pos]*Unit) int {
	result := 0
	for _, u := range units {
		if u.team == team {
			result++
		}
	}

	return result
}

// Generate list of adjacent positions (up, left, right, down)
func adjacents(p Pos) []Pos {
	return []Pos{Pos{p.x, p.y - 1}, Pos{p.x - 1, p.y}, Pos{p.x + 1, p.y}, Pos{p.x, p.y + 1}}
}

// Move a unit to its desired position
func (u *Unit) Move(field [][]byte, units map[Pos]*Unit) bool {

	// Check for adjacent targets
	for _, p := range adjacents(u.pos) {
		if units[p] != nil && units[p].team != u.team {
			return true
		}
	}

	// Filter out allies
	targets := values(units)

	for i := 0; i < len(targets); i++ {
		if targets[i].team != u.team {
			continue
		}

		targets[i], targets[len(targets)-1] = targets[len(targets)-1], targets[i]
		targets = targets[:len(targets)-1]
		i--
	}

	// Give up if no targets left alive
	if len(targets) == 0 {
		return false
	}

	// Find candidate locations
	loc := make(map[Pos]bool)

	for _, t := range targets {
		for _, p := range adjacents(t.pos) {
			if fieldClear(p, field, units) {
				loc[p] = true
			}
		}
	}

	if len(loc) == 0 {
		return false
	}

	// Find closest location by path
	var frontier [][2]Pos

	for _, p := range adjacents(u.pos) {
		if fieldClear(p, field, units) {
			frontier = append(frontier, [2]Pos{p, p})
		}
	}

	var candidates [][2]Pos
	visited := make(map[Pos]bool)
	visited[u.pos] = true

	used := make(map[[2]Pos]bool)

	// Best first search for routes
	for len(frontier) > 0 && len(candidates) == 0 {

		for count := len(frontier); count > 0; count-- {

			r := frontier[0]
			frontier = frontier[1:]

			p := r[1]
			visited[p] = true

			if loc[p] {
				candidates = append(candidates, r)
				continue
			}

			for _, cp := range adjacents(p) {
				if !visited[cp] && !used[[2]Pos{r[0], cp}] && fieldClear(cp, field, units) {
					frontier = append(frontier, [2]Pos{r[0], cp})
					used[[2]Pos{r[0], cp}] = true
				}
			}
		}
	}

	// Check if we have any possible routes
	if len(candidates) == 0 {
		return false
	}

	sortRoute(candidates)

	// Make best possible move
	delete(units, u.pos)
	u.pos = candidates[0][0]
	units[u.pos] = u

	return true
}

// Let a unit attack adjacent enemy
func (u *Unit) Attack(field [][]byte, units map[Pos]*Unit) bool {

	targets := make([]*Unit, 0, 4)

	for _, p := range adjacents(u.pos) {
		targets = append(targets, units[p])
	}

	// Filter out allies
	for i := 0; i < len(targets); i++ {
		if targets[i] != nil && targets[i].team != u.team {
			continue
		}

		targets[i], targets[len(targets)-1] = targets[len(targets)-1], targets[i]
		targets = targets[:len(targets)-1]
		i--
	}

	// Give up if no targets available
	if len(targets) == 0 {
		return false
	}

	// Pick best target
	sortAttack(targets)

	// Stab him!
	foe := targets[0]
	if foe.hp > u.ap {
		foe.hp -= u.ap
	} else {
		foe.hp = 0
		foe.alive = false
		delete(units, foe.pos)
	}

	return true
}

// Check if field position is clear of obstacles
func fieldClear(p Pos, field [][]byte, units map[Pos]*Unit) bool {
	return field[p.y][p.x] == '.' && units[p] == nil
}

// Load a new game field
func newField(raw string) ([][]byte, map[Pos]*Unit) {
	var field [][]byte
	units := make(map[Pos]*Unit)

	for y, s := range strings.Split(raw, "\n") {

		row := make([]byte, len(s))

		for x, r := range s {
			c := byte(r)
			switch c {
			case 'E', 'G':
				units[Pos{x, y}] = newUnit(x, y, c)
				row[x] = '.'
			default:
				row[x] = c
			}
		}

		field = append(field, row)
	}

	return field, units
}

// Print out a game field and its units
func draw(field [][]byte, units map[Pos]*Unit) {

	var buf bytes.Buffer

	for y, row := range field {
		var here []*Unit

		for x, c := range row {
			unit := units[Pos{x, y}]
			if unit != nil {
				c = unit.team
				here = append(here, unit)
			}
			buf.WriteString(fmt.Sprintf("%c", c))
		}

		for i, u := range here {
			padding := ", "
			if i == 0 {
				padding = "   "
			}
			buf.WriteString(fmt.Sprintf("%s%c(%d)", padding, u.team, u.hp))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

// Produce a new unit
func newUnit(x, y int, team byte) *Unit {
	return &Unit{
		Pos{x, y},
		200,
		3,
		team,
		true,
	}
}

// Generate list of values in a Unit map
func values(m map[Pos]*Unit) []*Unit {
	s := make([]*Unit, 0, len(m))

	for _, u := range m {
		s = append(s, u)
	}

	return s
}

// Check if one position sorts lower than another
func (p Pos) Less(op Pos) bool {
	if p.y == op.y {
		return p.x < op.x
	}
	return p.y < op.y
}

// Sort slice of units by their HP and Reading Order positions
func sortAttack(s []*Unit) {
	sort.Slice(s, func(i, j int) bool {
		if s[i].hp == s[j].hp {
			return s[i].pos.Less(s[j].pos)
		}
		return s[i].hp < s[j].hp
	})
}

// Sort slice of units by Reading Order positions
func sortLocation(s []*Unit) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].pos.Less(s[j].pos)
	})
}

// Sort slice of positions by Reading Order
func sortPos(s []Pos) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Less(s[j])
	})
}

// Sort routes by Reading Order of last Position
func sortRoute(s [][2]Pos) {
	sort.Slice(s, func(i, j int) bool {

		if s[i][1].y == s[j][1].y {
			if s[i][1].x == s[j][1].x {
				return s[i][0].Less(s[j][0])
			}
			return s[i][1].x < s[j][1].x
		}

		return s[i][1].y < s[j][1].y
	})
}

const INPUT1 = `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

const INPUT2 = `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`

const INPUT3 = `#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`

const INPUT4 = `#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`

const INPUT5 = `#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`

const INPUT6 = `#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`

const INPUT = `################################
#############..#################
#############..#.###############
############G..G.###############
#############....###############
##############.#...#############
################..##############
#############G.##..#..##########
#############.##.......#..######
#######.####.G##.......##.######
######..####.G.......#.##.######
#####.....#..GG....G......######
####..###.....#####.......######
####.........#######..E.G..#####
####.G..G...#########....E.#####
#####....G.G#########.#...######
###........G#########....#######
##..#.......#########....##.E.##
##.#........#########.#####...##
#............#######..#.......##
#.G...........#####........E..##
#....G........G..G.............#
#..................E#...E...E..#
#....#...##...G...E..........###
#..###...####..........G###E.###
#.###########..E.......#########
#.###############.......########
#################.......########
##################....#..#######
##################..####.#######
#################..#####.#######
################################`

const INPUT_TAEL = `################################
##########..........############
########G..................#####
#######..G.GG...............####
#######....G.......#......######
########.G.G...............#E..#
#######G.................#.....#
########.......................#
########G.....G....#.....##....#
########.....#....G.........####
#########..........##....E.E#.##
##########G..G..........#####.##
##########....#####G....####E.##
######....G..#######.....#.....#
###....#....#########......#####
####........#########..E...#####
###.........#########......#####
####G....G..#########......#####
####..#.....#########....#######
######.......#######...E.#######
###.G.....E.G.#####.....########
#.....G........E.......#########
#......#..#..####....#.#########
#...#.........###.#..###########
##............###..#############
######.....E####..##############
######...........###############
#######....E....################
######...####...################
######...###....################
###.....###..##..###############
################################`

const INPUT_REDDIT1 = `################################
####.#######..G..########.....##
##...........G#..#######.......#
#...#...G.....#######..#......##
########.......######..##.E...##
########......G..####..###....##
#...###.#.....##..G##.....#...##
##....#.G#....####..##........##
##..#....#..#######...........##
#####...G.G..#######...G......##
#########.GG..G####...###......#
#########.G....EG.....###.....##
########......#####...##########
#########....#######..##########
#########G..#########.##########
#########...#########.##########
######...G..#########.##########
#G###......G#########.##########
#.##.....G..#########..#########
#............#######...#########
#...#.........#####....#########
#####.G..................#######
####.....................#######
####.........E..........########
#####..........E....E....#######
####....#.......#...#....#######
####.......##.....E.#E...#######
#####..E...####.......##########
########....###.E..E############
#########.....##################
#############.##################
################################`

const INPUT_REDDIT2 = `################################
###########################..###
##########################...###
#########################..#####
####...##################.######
#####..################...#.####
#..G...G#########.####G.....####
#.......########.....G.......###
#.....G....###G....#....E.....##
####...##......##.............##
####G...#.G...###.G...........##
####G.......................####
####.........G#####.........####
####...GG#...#######.......#####
###.........#########G....######
###.G.......#########G...#######
###.G.......#########......#####
####.....G..#########....E..####
#####.......#########..E....####
######...##G.#######........####
######.#.#.G..#####.....##..####
########....E...........##..####
########....E#######........####
########......######E....##..E.#
########......#####.....#......#
########.....######............#
##################...#.E...E...#
##################.............#
###################.......E#####
####################....#...####
####################.###########
################################`

const INPUT_REDDIT3 = `################################
###############.##...###########
##############..#...G.#..#######
##############.............#####
###############....G....G......#
##########..........#..........#
##########................##..##
######...##..G...G.......####..#
####..G..#G...............####.#
#######......G....G.....G#####E#
#######.................E.######
########..G...............######
######....G...#####E...G....####
######..G..G.#######........####
###.........#########.......E.##
###..#..#...#########...E.....##
######......#########.......####
#####...G...#########.....######
#####G......#########.....######
#...#G..G....#######......######
###...##......#####.......######
####..##..G........E...E..######
#####.####.....######...########
###########..#...####...E.######
###############...####..#...####
###############...###...#.E.####
#####################.#E....####
#####################.#...######
###################...##.#######
##################..############
##################...###########
################################`
