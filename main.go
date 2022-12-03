package main

import (
	"fmt"
	"os"

	"github.com/phootip/advent-of-code-2022/day1"
	"github.com/phootip/advent-of-code-2022/day2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input number to select which day you want")
		return
	}
	switch os.Args[1] {
	case "1":
		fmt.Println(day1.Sol1())
		fmt.Println(day1.Sol2())
	case "2":
		fmt.Println(day2.Sol1())
		fmt.Println(day2.Sol2())
	}
}