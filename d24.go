package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {

	fmt.Println("Starting...")
	p1(INPUT) // INPUT = 16006/6221, TAEL = 20753/3013
	p2(INPUT)
	// printer(INPUT, 0)
}

func p1(input string) {

	t := time.Now()

	g := load(input, 0)

	for len(g.Immune) > 0 && len(g.Infection) > 0 {
		order := g.Targeting()
		progress := g.Attacking(order, false)

		if !progress {
			fmt.Println("Stalemate!")
			break
		}
	}

	result := 0
	for u := range g.Immune {
		result += u.size
	}
	for u := range g.Infection {
		result += u.size
	}

	fmt.Println("Done:", result, time.Since(t))
}

func p2(input string) {

	t := time.Now()

	result := 0

	boost := 0
	survival := 1000000000
	death := 0

	for {
		g := load(input, boost)

		for len(g.Immune) > 0 && len(g.Infection) > 0 {
			order := g.Targeting()
			progress := g.Attacking(order, false)

			if !progress {
				// fmt.Println("Stalemate at", boost)
				break
			}
		}

		if len(g.Infection) == 0 {

			result = 0
			for u := range g.Immune {
				result += u.size
			}

			boost, survival = death+(boost-death)/2, boost
			// fmt.Println("Survival:", survival, "Trying:", boost)

		} else {
			boost, death = min((boost+1)*2, boost+(survival-boost)/2), boost
			// fmt.Println("Death:", death, "Trying:", boost)
		}

		if death+1 >= survival {
			fmt.Println("Survived at", survival, "with", result, "units left")
			break
		}

		// time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Done 2:", result, time.Since(t))
}

func printer(input string, boost int) {

	t := time.Now()

	g := load(input, boost)

	round := 0

	for len(g.Immune) > 0 && len(g.Infection) > 0 {
		round++

		fmt.Println("- ROUND", round, "-")
		fmt.Println("Immune System:")
		for u := range g.Immune {
			fmt.Printf("Group %d contains %d units\n", u.id, u.size)
		}

		fmt.Println("Infection:")
		for u := range g.Infection {
			fmt.Printf("Group %d contains %d units\n", u.id, u.size)
		}

		fmt.Println("")

		order := g.Targeting()
		progress := g.Attacking(order, true)

		fmt.Println("")

		if !progress {
			fmt.Println("Stalemate!")
			break
		}

		fmt.Println("")

		// time.Sleep(100 * time.Millisecond)
	}

	result := 0
	for u := range g.Immune {
		result += u.size
	}
	for u := range g.Infection {
		result += u.size
	}

	fmt.Println("Done:", result, time.Since(t))
}

const (
	BLUDGEONING = 1 << iota
	COLD
	FIRE
	RADIATION
	SLASHING
)

var mapping map[string]int

func init() {
	mapping = map[string]int{
		"bludgeoning": BLUDGEONING,
		"cold":        COLD,
		"fire":        FIRE,
		"radiation":   RADIATION,
		"slashing":    SLASHING,
	}
}

type Unit struct {
	side       string
	id         int
	size       int
	hp         int
	ap         int
	initiative int
	immunity   int
	weakness   int
	attack     int
	target     *Unit
}

type Game struct {
	Immune    map[*Unit]bool
	Infection map[*Unit]bool
}

func ep(u *Unit) int {
	return u.size * u.ap
}

func load(input string, boost int) *Game {

	g := &Game{}

	for _, s := range strings.SplitN(input, "\n\n", 2) {

		var units map[*Unit]bool
		var side string
		var buff int

		if strings.HasPrefix(s, "Immune") {
			g.Immune = make(map[*Unit]bool)
			units = g.Immune
			side = "Immune System"
			buff = boost
		} else if strings.HasPrefix(s, "Infection") {
			g.Infection = make(map[*Unit]bool)
			units = g.Infection
			side = "Infection"
			buff = 0
		} else {
			panic("Could not parse side:" + s)
		}

		id := 0

		for _, unit := range strings.Split(s, "\n")[1:] {
			id++

			var u Unit
			u.id = id
			u.side = side

			u.size, _ = strconv.Atoi(strings.Fields(unit)[0])
			u.hp, _ = strconv.Atoi(strings.Fields(unit)[4])

			if strings.Contains(unit, "immune to") {
				for _, kind := range strings.FieldsFunc(strings.Split(unit, "immune to")[1], func(c rune) bool { return !unicode.IsLetter(c) }) {
					if k, ok := mapping[kind]; ok {
						u.immunity |= k
					} else {
						break
					}
				}
			}

			if strings.Contains(unit, "weak to") {
				for _, kind := range strings.FieldsFunc(strings.Split(unit, "weak to")[1], func(c rune) bool { return !unicode.IsLetter(c) }) {
					if k, ok := mapping[kind]; ok {
						u.weakness |= k
					} else {
						break
					}
				}
			}

			unit = strings.SplitN(unit, "attack that does", 2)[1]

			u.ap, _ = strconv.Atoi(strings.Fields(unit)[0])
			u.ap += buff
			u.attack |= mapping[strings.Fields(unit)[1]]
			u.initiative, _ = strconv.Atoi(strings.Fields(unit)[5])

			units[&u] = true
		}
	}

	return g
}

func sortAttackers(s []*Unit) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].initiative > s[j].initiative
	})
}

func sortSelectors(s []*Unit) {
	sort.Slice(s, func(i, j int) bool {
		if ep(s[i]) == ep(s[j]) {
			return s[i].initiative > s[j].initiative
		}
		return ep(s[i]) > ep(s[j])
	})
}

func sortTargets(s []*Unit) {
	sort.Slice(s, func(i, j int) bool {
		if ep(s[i]) == ep(s[j]) {
			return s[i].initiative > s[j].initiative
		}
		return ep(s[i]) > ep(s[j])
	})
}

// All units choose their targets
func (g *Game) Targeting() []*Unit {

	selectors := make([]*Unit, 0, len(g.Immune)+len(g.Infection))
	for u := range g.Immune {
		selectors = append(selectors, u)
	}

	for u := range g.Infection {
		selectors = append(selectors, u)
	}

	sortSelectors(selectors)

	// targets := append([]*Unit{}, selectors...)
	// targets := make([]*Unit, len(selectors))
	// copy(targets, selectors)

	// sortTargets(targets)

	taken := make(map[*Unit]bool)

	for _, selector := range selectors {
		selector.target = nil

		// for _, target := range targets {
		for _, target := range selectors {
			if selector.side == target.side ||
				selector.attack&target.immunity > 0 ||
				taken[target] {
				continue
			}

			if selector.target == nil ||
				(selector.attack&target.weakness > 0 &&
					selector.attack&selector.target.weakness == 0) {
				taken[selector.target] = false
				selector.target = target
				taken[target] = true
			}
		}
	}

	return selectors
}

// All units execute their attack
func (g *Game) Attacking(attackers []*Unit, printing bool) bool {

	sortAttackers(attackers)

	progress := false

	for _, attacker := range attackers {
		if attacker.target == nil || attacker.size == 0 {
			continue
		}

		t := attacker.target

		damage := ep(attacker)
		if attacker.attack&t.weakness > 0 {
			damage *= 2
		}

		if printing {
			fmt.Printf("%s group %d attacks defending group %d, killing %d/%d units\n",
				attacker.side,
				attacker.id,
				t.id,
				min(damage/t.hp, t.size),
				t.size,
			)
		}

		if damage >= t.hp {
			progress = true
		}

		t.size = max(0, t.size-damage/t.hp)
		if t.size == 0 {
			if t.side == "Infection" {
				delete(g.Infection, t)
			} else {
				delete(g.Immune, t)
			}
		}
	}

	return progress
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

const INPUT1 = `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`

const INPUT = `Immune System:
3020 units each with 3290 hit points with an attack that does 10 radiation damage at initiative 16
528 units each with 6169 hit points with an attack that does 113 fire damage at initiative 9
4017 units each with 2793 hit points (weak to radiation) with an attack that does 6 slashing damage at initiative 1
2915 units each with 7735 hit points with an attack that does 26 cold damage at initiative 4
3194 units each with 1773 hit points (immune to radiation; weak to fire) with an attack that does 5 cold damage at initiative 13
1098 units each with 4711 hit points with an attack that does 36 radiation damage at initiative 7
2530 units each with 3347 hit points (immune to slashing) with an attack that does 12 bludgeoning damage at initiative 5
216 units each with 7514 hit points (immune to cold, slashing; weak to bludgeoning) with an attack that does 335 slashing damage at initiative 15
8513 units each with 9917 hit points (immune to slashing; weak to cold) with an attack that does 10 fire damage at initiative 14
1616 units each with 3771 hit points with an attack that does 19 bludgeoning damage at initiative 10

Infection:
1906 units each with 37289 hit points (immune to radiation; weak to fire) with an attack that does 28 radiation damage at initiative 3
6486 units each with 32981 hit points with an attack that does 9 bludgeoning damage at initiative 18
489 units each with 28313 hit points (immune to radiation, bludgeoning) with an attack that does 110 bludgeoning damage at initiative 6
1573 units each with 44967 hit points (weak to bludgeoning, cold) with an attack that does 42 slashing damage at initiative 12
2814 units each with 11032 hit points (immune to fire, slashing; weak to radiation) with an attack that does 7 slashing damage at initiative 2
1588 units each with 18229 hit points (weak to slashing; immune to radiation, cold) with an attack that does 20 radiation damage at initiative 19
608 units each with 39576 hit points (immune to bludgeoning) with an attack that does 116 slashing damage at initiative 20
675 units each with 48183 hit points (immune to cold, slashing, bludgeoning) with an attack that does 138 slashing damage at initiative 8
685 units each with 11702 hit points with an attack that does 32 fire damage at initiative 17
1949 units each with 32177 hit points with an attack that does 32 radiation damage at initiative 11`

const INPUT_TAEL = `Immune System:
597 units each with 4458 hit points with an attack that does 73 slashing damage at initiative 6
4063 units each with 9727 hit points (weak to radiation) with an attack that does 18 radiation damage at initiative 9
2408 units each with 5825 hit points (weak to slashing; immune to fire, radiation) with an attack that does 17 slashing damage at initiative 2
5199 units each with 8624 hit points (immune to fire) with an attack that does 16 radiation damage at initiative 15
1044 units each with 4485 hit points (weak to bludgeoning) with an attack that does 41 radiation damage at initiative 3
4890 units each with 9477 hit points (immune to cold; weak to fire) with an attack that does 19 slashing damage at initiative 7
1280 units each with 10343 hit points with an attack that does 64 cold damage at initiative 19
609 units each with 6435 hit points with an attack that does 86 cold damage at initiative 17
480 units each with 2750 hit points (weak to cold) with an attack that does 57 fire damage at initiative 11
807 units each with 4560 hit points (immune to fire, slashing; weak to bludgeoning) with an attack that does 56 radiation damage at initiative 8

Infection:
1237 units each with 50749 hit points (weak to radiation; immune to cold, slashing, bludgeoning) with an attack that does 70 radiation damage at initiative 12
4686 units each with 25794 hit points (immune to cold, slashing; weak to bludgeoning) with an attack that does 10 bludgeoning damage at initiative 14
1518 units each with 38219 hit points (weak to slashing, fire) with an attack that does 42 radiation damage at initiative 16
4547 units each with 21147 hit points (weak to fire; immune to radiation) with an attack that does 7 slashing damage at initiative 4
1275 units each with 54326 hit points (immune to cold) with an attack that does 65 cold damage at initiative 20
436 units each with 36859 hit points (immune to fire, cold) with an attack that does 164 fire damage at initiative 18
728 units each with 53230 hit points (weak to radiation, bludgeoning) with an attack that does 117 fire damage at initiative 5
2116 units each with 21754 hit points with an attack that does 17 bludgeoning damage at initiative 10
2445 units each with 21224 hit points (immune to cold) with an attack that does 16 cold damage at initiative 13
3814 units each with 22467 hit points (weak to bludgeoning, radiation) with an attack that does 10 cold damage at initiative 1`
