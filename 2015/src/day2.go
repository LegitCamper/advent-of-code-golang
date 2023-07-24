package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var paper int = 0
	var ribbon int = 0

	for scanner.Scan() {
		paper += calc_paper(scanner.Text())
		ribbon += calc_ribbon(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("You need ", paper, "paper")
	fmt.Println("You need ", ribbon, "ribbon")
}

func calc_ribbon(box string) int {
	box_split := strings.Split(box, "x")
	var box_lens []int
	for i := 0; i < len(box_split); i++ {
		var l, l_err = strconv.Atoi(box_split[i])
		if l_err != nil {
			panic(l_err)
		} else {
			box_lens = append(box_lens, l)
		}
	}
	sort.Ints(box_lens)
	var ribbon int = (box_lens[0] * 2) + (box_lens[1] * 2) + (box_lens[0] * box_lens[1] * box_lens[2])

	return ribbon
}

func calc_paper(box string) int {
	box_split := strings.Split(box, "x")
	var box_lens []int
	for i := 0; i < len(box_split); i++ {
		var l, l_err = strconv.Atoi(box_split[i])
		if l_err != nil {
			panic(l_err)
		} else {
			box_lens = append(box_lens, l)
		}
	}

	var paper int = 0
	var shortest int = 0

	for i := 0; i < len(box_lens); i++ {
		var leng int
		if i == len(box_lens)-1 {
			leng = (box_lens[i] * box_lens[0])
		} else {
			leng = (box_lens[i] * box_lens[i+1])
		}
		paper += leng * 2

		if shortest == 0 {
			shortest = leng
		} else {
			if leng < shortest {
				shortest = leng
			}
		}

	}
	return paper + shortest
}
