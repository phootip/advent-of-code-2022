package main

import (
	"fmt"
	"os"

	"github.com/phootip/advent-of-code-2022/day1"
	"github.com/phootip/advent-of-code-2022/day10"
	"github.com/phootip/advent-of-code-2022/day11"
	"github.com/phootip/advent-of-code-2022/day12"
	"github.com/phootip/advent-of-code-2022/day13"
	"github.com/phootip/advent-of-code-2022/day14"
	"github.com/phootip/advent-of-code-2022/day15"
	"github.com/phootip/advent-of-code-2022/day16"
	"github.com/phootip/advent-of-code-2022/day17"
	"github.com/phootip/advent-of-code-2022/day18"
	"github.com/phootip/advent-of-code-2022/day19"
	"github.com/phootip/advent-of-code-2022/day2"
	"github.com/phootip/advent-of-code-2022/day3"
	"github.com/phootip/advent-of-code-2022/day4"
	"github.com/phootip/advent-of-code-2022/day5"
	"github.com/phootip/advent-of-code-2022/day6"
	"github.com/phootip/advent-of-code-2022/day7"
	"github.com/phootip/advent-of-code-2022/day8"
	"github.com/phootip/advent-of-code-2022/day9"
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
		fmt.Println("answer: ", day5.Sol1())
		fmt.Println("answer: ", day5.Sol2())
	case "6":
		fmt.Println("answer: ", day6.Sol1())
		fmt.Println("answer: ", day6.Sol2())
	case "7":
		fmt.Println("answer: ", day7.Sol1())
		fmt.Println("answer: ", day7.Sol2())
	case "8":
		fmt.Println("answer: ", day8.Sol1())
		fmt.Println("answer: ", day8.Sol2())
	case "9":
		fmt.Println("answer: ", day9.Sol1())
		fmt.Println("answer: ", day9.Sol2())
	case "10":
		fmt.Println("answer: ", day10.Sol1())
		fmt.Print("answer:\n", day10.Sol2())
	case "11":
		fmt.Println("answer: ", day11.Sol1())
		fmt.Println("answer: ", day11.Sol2())
	case "12":
		fmt.Println("answer: ", day12.Sol1())
		fmt.Println("answer: ", day12.Sol2())
	case "13":
		fmt.Println("answer: ", day13.Sol1())
		fmt.Println("answer: ", day13.Sol2())
	case "14":
		fmt.Println("answer: ", day14.Sol1())
		fmt.Println("answer: ", day14.Sol2())
	case "15":
		fmt.Println("answer: ", day15.Sol1())
		fmt.Println("answer: ", day15.Sol2())
	case "16":
		fmt.Println("answer: ", day16.Sol1())
		fmt.Println("answer: ", day16.Sol2())
	case "17":
		fmt.Println("answer: ", day17.Sol1())
		fmt.Println("answer: ", day17.Sol2())
	case "18":
		fmt.Println("answer: ", day18.Sol1())
		fmt.Println("answer: ", day18.Sol2())
	case "19":
		// fmt.Println("answer: ", day19.Sol1())
		fmt.Println("answer: ", day19.Sol2())
	}

}
