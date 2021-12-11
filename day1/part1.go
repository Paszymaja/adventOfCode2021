package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var incNum int

	dt, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(dt)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	prevDepth, _ := strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		if depth > prevDepth {
			incNum++
		}
		prevDepth = depth
	}
	fmt.Println(incNum)
}
