package main

import (
	"fmt"
	"os"

	"github.com/phootip/advent-of-code-2022/day1"
	"github.com/phootip/advent-of-code-2022/day2"
	"github.com/phootip/advent-of-code-2022/day3"
	"github.com/phootip/advent-of-code-2022/day4"
	"github.com/phootip/advent-of-code-2022/day5"
	"github.com/phootip/advent-of-code-2022/draw"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input number to select which day you want")
		return
	}
	switch os.Args[1] {
	case "draw":
		draw.Draw()
	case "1":
		fmt.Println(day1.Sol1())
		fmt.Println(day1.Sol2())
	case "2":
		fmt.Println(day2.Sol1())
		fmt.Println(day2.Sol2())
	case "3":
		fmt.Println("answer: ", day3.Sol1())
		fmt.Println("answer: ", day3.Sol2())
	case "4":
		fmt.Println("answer: ", day4.Sol1())
		fmt.Println("answer: ", day4.Sol2())
	case "5":
		// fmt.Println("answer: ", day5.Sol1())
		fmt.Println("answer: ", day5.Sol2())
	}
}
