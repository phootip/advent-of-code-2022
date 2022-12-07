package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day7 Solution1...")
	raw := utils.ReadFile("./day7/input.txt")
	// raw := utils.ReadFile("./day7/example.txt")
	dirMap := genDir(raw)
	fmt.Println(dirMap)

	fmt.Println("Calculating each dir size...")
	// calTotalSizeLoop(dirMap)
	calTotalSize(dirMap, "/")
	return calAns(dirMap)
	// return 0
}
func calAns(dirMap map[string]any) int {
	ans := 0
	for k, v := range dirMap {
		fmt.Println(k, v.(map[string]any)["totalSize"])
		size := v.(map[string]any)["totalSize"].(int)
		if size <= 100000 {
			ans += size
		}
	}
	return ans

}

func genDir(raw []string) map[string]any {
	pwd := []string{}
	dirMap := map[string]any{}
	for _, line := range raw {
		if line == "" {
			return dirMap
		}
		tokenized := strings.Split(line, " ")
		if tokenized[0] == "$" {
			pwd = parseCommand(tokenized[1:], pwd)
		} else {
			parseList(pwd, tokenized, dirMap)
		}
	}
	return dirMap
}

func calTotalSizeLoop(dirMap map[string]any) {
	count := len(dirMap)
	for count > 0 {
		fmt.Println(count)
		for position, dirList := range dirMap {
			if mapContainsKey(dirList, "totalSize") {
				continue
			}
			// fmt.Println(position)
			// if count == 41 {
			// 	fmt.Println(position)
			// }
			totalSize := sumSize(dirMap, dirList)
			if totalSize != -1 {
				count--
				dirMap[position].(map[string]any)["totalSize"] = totalSize
				// fmt.Println(position)
			} else {
				fmt.Println(totalSize)
				fmt.Println(position)
				fmt.Println(dirMap[position])
			}
		}
	}
}

func sumSize(dirMap map[string]any, m any) int {
	totalSize := 0
	for file, size := range m.(map[string]any) {
		if size == "dir" {
			if !mapContainsKey(dirMap[file], "totalSize") {
				return -1
			} else {
				totalSize += dirMap[file].(map[string]any)["totalSize"].(int)
			}
		} else {
			totalSize += size.(int)
		}
	}
	return totalSize
}

func mapContainsKey(m any, k string) bool {
	_, ok := m.(map[string]any)[k]
	return ok
}

func calTotalSize(dirMap map[string]any, position string) int {
	if val, ok := dirMap[position].(map[string]any)["totalSize"]; ok {
		// fmt.Println("already have totalSize: ", position)
		return val.(int)
	}
	dirList := dirMap[position].(map[string]any)
	ans := 0
	for k,v := range dirList {
		if v != "dir" {
			ans += v.(int)
		} else {
			ans += calTotalSize(dirMap, position+"/"+k)
		}
	}
	dirMap[position].(map[string]any)["totalSize"] = ans
	return ans
}

func parseList(pwd []string, tokenized []string, dirMap map[string]any) {
	// position := pwd[len(pwd)-1]
	position := strings.Join(pwd, "/")
	fileType := tokenized[0]
	name := tokenized[1]
	// fmt.Println(position, fileType, name)
	if dirMap[position] == nil {
		dirMap[position] = map[string]any{}
	}
	if fileType == "dir" {
		dirMap[position].(map[string]any)[name] = "dir"
		// add filetype dir handler
	} else {
		fileSize, err := strconv.Atoi(fileType)
		utils.Check(err)
		dirMap[position].(map[string]any)[name] = fileSize
	}
	// fmt.Println("dirMap", dirMap)
}

func parseCommand(cmd []string, pwd []string) []string {
	if cmd[0] == "cd" {
		if cmd[1] == ".." {
			return pwd[:len(pwd)-1]
		}
		return append(pwd, cmd[1])
	}
	return pwd
}
