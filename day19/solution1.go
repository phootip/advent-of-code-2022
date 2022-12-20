package day19

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

// var mem map[[4]int]map[[4]int]int
var mem map[string]int
var allType [4]string

func Sol1() (ans int) {
	fmt.Println("Starting Day19 Solution1...")
	allType = [4]string{"ore", "clay", "obsidian", "geode"}
	// raw := utils.ReadFile("./day19/input.txt")
	raw := utils.ReadFile("./day19/example.txt")
	raw = raw[:len(raw)-1]
	blueprints := rawToBlueprint(raw)
	ans = bestBlueprint(blueprints)
	return ans
}

func bestBlueprint(blueprints []*Blueprint) (ans int) {
	for _, blueprint := range blueprints {
		mem = make(map[string]int)
		ans = utils.Max(ans, bestGeodes(blueprint))
		// fmt.Println(mem)
	}
	return ans
}

func bestGeodes(blueprint *Blueprint) (result int) {
	fmt.Println("calculating blueprint: ", blueprint.id)
	time := 0
	bestPath := []string{"clay","clay","obsidian","clay","obsidian","geode","geode"}
	_ = bestPath
	// result = utils.Max(process(blueprint, time, "ore", []string{}, bestPath), process(blueprint, time, "clay", []string{}, bestPath))
	fmt.Println(blueprint)
	result, bestPath = process(blueprint, time, "clay", []string{}, bestPath)
	fmt.Println("result1: ", result)
	fmt.Println("bestPath: ", bestPath)
	fmt.Println(blueprint)
	result2, bestPath2 := process(blueprint, time, "ore", []string{}, bestPath)
	fmt.Println("result2: ", result2)
	fmt.Println("bestPath2: ", bestPath2)
	// return result
	return utils.Max(result, result2)
}

func process(oldBlueprint *Blueprint, time int, target string, path []string, debugPath []string) (int,[]string) {
	if mem[strings.Join(path,",")] != 0 {
		return mem[strings.Join(path,",")], path
	}
	blueprint := copyBlueprint(oldBlueprint)
	// fmt.Println("time: ", time)
	// fmt.Println("target: ", target)
	// fmt.Println("path: ", path)
	result := 0
	robot := blueprint.robots[target]
	// fmt.Println(robot)
	waitTime := robot.timeTillBuildable(blueprint)
	if waitTime > 24-time {
		timeLeft := 24-time
		result = blueprint.resource["geode"]
		result += timeLeft*blueprint.robotCounts["geode"]
		// if result == 5 {
		// 	fmt.Println("path: ", path)
		// 	fmt.Println("result: ", result)
		// }
		mem[strings.Join(path,",")] = result
		if result == 0 {
			mem[strings.Join(path,",")] = -1
		}
		return result, path
	}
	// newBlueprint := copyBlueprint(blueprint)
	blueprint.updateResource(waitTime)
	blueprint.build(target)
	path = append(path, target)
	endtime := time + waitTime
	// fmt.Println("endtime: ", endtime)
	// blueprint.debugResource()
	// if len(debugPath) > 0 {
	// 	newBlueprint := copyBlueprint(blueprint)
	// 	nextTarget := debugPath[0]
	// 	debugPath = debugPath[1:]
	// 	result = utils.Max(result, process(newBlueprint, endtime, nextTarget, path, debugPath))
	// 	return result
	// }
	// if len(debugPath) == 0 {
	// 	return process(blueprint, endtime, "geode", path, debugPath)
	// }
	bestPath := []string{}
	_ = bestPath
	for _, t := range allType {
		newBlueprint := copyBlueprint(blueprint)
		newResult, newPath := process(newBlueprint, endtime, t, path, debugPath)
		if newResult > result {
			result = newResult
			bestPath = newPath
		}
	}
	// if time > 0 {
	// 	fmt.Println("bestPath: ", bestPath)
	// }
	mem[strings.Join(path,",")] = result
	if result == 0 {
		mem[strings.Join(path,",")] = -1
	}
	// fmt.Println("bestPath :",bestPath)
	return result, bestPath
}

func (robot *Robot) timeTillBuildable(blueprint *Blueprint) int {
	for cost := range robot.costs {
		if blueprint.robotCounts[cost] == 0 {
			return 9999
		}
	}
	maxTime := 0
	for cost := range robot.costs {
		costLeft := robot.costs[cost] - blueprint.resource[cost]
		// fmt.Println("cost Type: ", cost)
		// fmt.Println("robot Cost: ", robot.costs[cost])
		// fmt.Println("resource now: ", blueprint.resource[cost])
		// fmt.Println("cost left: ", costLeft)
		// fmt.Println("robotCount: ", blueprint.robotCounts[cost])
		time := costLeft / blueprint.robotCounts[cost]
		if costLeft%blueprint.robotCounts[cost] != 0 {
			time++
		}
		maxTime = utils.Max(maxTime, time)
	}
	// fmt.Println("waitTime: ", maxTime)
	return maxTime+1
}

// func memBlueprint(blueprint *Blueprint, result int) {
// 	fmt.Println(&blueprint)
// 	mem[blueprint] = result
// 	if mem[blueprint] == 0 {
// 		mem[blueprint] = -1
// 	}
// }

func mapToArr(a map[string]int) [4]int {
	result := [4]int{}
	for i, k := range [4]string{"ore", "clay", "obsidian", "geode"} {
		result[i] = a[k]
	}
	return result
}

func (blueprint *Blueprint) updateResource(waitTime int) {
	for k := range blueprint.resource {
		blueprint.resource[k] += blueprint.robotCounts[k] * waitTime
		// fmt.Println("robotCount: ", k, blueprint.robotCounts[k])
		// fmt.Println("resource: ", k, blueprint.resource[k])
	}
}

func (blueprint *Blueprint) debugResource() {
	for _, k := range allType {
		fmt.Println("robotCount: ", k, blueprint.robotCounts[k])
		// fmt.Println("resource: ", k, blueprint.resource[k])
	}
	for _, k := range allType {
		fmt.Println("resource: ", k, blueprint.resource[k])
	}
	fmt.Println()
}

func (blueprint *Blueprint) build(robotType string) {
	blueprint.robotCounts[robotType] += 1
	robot := blueprint.robots[robotType]
	for k := range robot.costs {
		blueprint.resource[k] -= robot.costs[k]
	}
}

func copyBlueprint(blueprint *Blueprint) *Blueprint {
	newBlueprint := Blueprint{}
	newBlueprint.robots = blueprint.robots
	newBlueprint.robotCounts = make(map[string]int)
	newBlueprint.resource = make(map[string]int)
	utils.CopyMap(newBlueprint.robotCounts, blueprint.robotCounts)
	utils.CopyMap(newBlueprint.resource, blueprint.resource)
	return &newBlueprint
}

func (robot *Robot) buildable(resource map[string]int) bool {
	for cost := range robot.costs {
		if robot.costs[cost] > resource[cost] {
			return false
		}
	}
	return true
}

func rawToBlueprint(raw []string) (blueprints []*Blueprint) {
	for _, line := range raw {
		line2 := strings.Split(line, " ")
		blueprint := Blueprint{}
		blueprint.id = utils.StringToInt(line2[1][:len(line2[1])-1])
		blueprint.robots = map[string]*Robot{}
		blueprint.robotCounts = map[string]int{}
		blueprint.resource = map[string]int{}
		blueprint.robots["ore"] = &Robot{robotType: "ore", costs: map[string]int{"ore": utils.StringToInt(line2[6])}}
		blueprint.robots["clay"] = &Robot{robotType: "clay", costs: map[string]int{"ore": utils.StringToInt(line2[12])}}
		blueprint.robots["obsidian"] = &Robot{robotType: "obsidian", costs: map[string]int{"ore": utils.StringToInt(line2[18]), "clay": utils.StringToInt(line2[21])}}
		blueprint.robots["geode"] = &Robot{robotType: "geode", costs: map[string]int{"ore": utils.StringToInt(line2[27]), "obsidian": utils.StringToInt(line2[30])}}
		for k := range blueprint.robots {
			blueprint.robotCounts[k] = 0
			blueprint.resource[k] = 0
		}
		blueprint.robotCounts["ore"] = 1
		blueprints = append(blueprints, &blueprint)
	}
	return blueprints
}

type Blueprint struct {
	id          int
	robots      map[string]*Robot
	robotCounts map[string]int
	resource    map[string]int
}

type Robot struct {
	robotType string
	costs     map[string]int
}
