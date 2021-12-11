package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mv struct {
	dir string
	val int
}

func main() {
	path := "day2/input.txt"
	var mvs []mv
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		v, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		mvs = append(mvs, mv{split[0], v})
	}

	var h, d, aim int

	for _, i := range mvs {
		switch i.dir {
		case "forward":
			h += i.val
			d += (aim * i.val)
		case "down":
			aim += i.val
		case "up":
			aim -= i.val
		}
	}
	fmt.Print(h * d)
}
