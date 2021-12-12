package main

import (
	"fmt"
	"strings"
)

type Cave struct {
	name string
	Big  bool
	con  []string
}

var test_data = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end",
}

var data = []string{
	"HF-qu",
	"end-CF",
	"CF-ae",
	"vi-HF",
	"vt-HF",
	"qu-CF",
	"hu-vt",
	"CF-pk",
	"CF-vi",
	"qu-ae",
	"ae-hu",
	"HF-start",
	"vt-end",
	"ae-HF",
	"end-vi",
	"vi-vt",
	"hu-start",
	"start-ae",
	"CS-hu",
	"CF-vt",
}

func IsUpperCase(s string) bool {
	return s == string(strings.ToUpper(string(s)))
}

func newCave(name string, con []string) *Cave {
	return &Cave{name, IsUpperCase(name), con}
}

func main() {
	caves := make(map[string]*Cave)
	for _, line := range data {
		parts := strings.Split(line, "-")

		if _, ok := caves[parts[0]]; !ok {
			caves[parts[0]] = newCave(parts[0], []string{parts[1]})
		} else {
			caves[parts[0]].con = append(caves[parts[0]].con, parts[1])
		}
		parts[0], parts[1] = parts[1], parts[0]
		if _, ok := caves[parts[0]]; !ok {
			caves[parts[0]] = newCave(parts[0], []string{parts[1]})
		} else {
			caves[parts[0]].con = append(caves[parts[0]].con, parts[1])
		}
	}
	fmt.Println(FindPaths(caves, "start", make(map[string]int)))
	fmt.Println(FindPathsTwice(caves, "start", make(map[string]int), false))

}

func FindPaths(cv map[string]*Cave, currentCave string, visits map[string]int) int {
	visits[currentCave] += 1
	sum := 0
	for _, c := range cv[currentCave].con {
		if c == "end" {
			sum++
			continue
		}
		if cv[c].Big || visits[c] == 0 {
			nextVisits := visits
			if !cv[c].Big {
				nextVisits = copyMap(visits)
			}
			sum += FindPaths(cv, c, nextVisits)
		}
	}
	return sum
}

func FindPathsTwice(cv map[string]*Cave, currentCave string, visits map[string]int, hasVisitedTwice bool) int {

	if visits[currentCave] > 0 && !cv[currentCave].Big {
		hasVisitedTwice = true
	}

	visits[currentCave] += 1
	sum := 0
	for _, c := range cv[currentCave].con {
		if c == "end" {
			sum++
			continue
		}
		if cv[c].name != "start" && (cv[c].Big || visits[c] < 1 || !hasVisitedTwice) {
			nextVisits := visits
			if !cv[c].Big {
				nextVisits = copyMap(visits)
			}
			sum += FindPathsTwice(cv, c, nextVisits, hasVisitedTwice)
		}
	}
	return sum
}

func copyMap(a map[string]int) map[string]int {
	out := make(map[string]int, len(a))
	for k, v := range a {
		out[k] = v
	}
	return out
}
