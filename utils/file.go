package utils

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) []string {
	b, err := os.ReadFile(filename)
	s := string(b)
	Check(err)
	data := strings.Split(s, "\n")
	return data

}
