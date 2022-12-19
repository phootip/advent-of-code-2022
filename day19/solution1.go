package day19

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

// var mem map[[4]int]map[[4]int]int
var mem map[*Blueprint]int

func Sol1() (ans int) {
	fmt.Println("Starting Day19 Solution1...")
	// raw := utils.ReadFile("./day19/input.txt")
	raw := utils.ReadFile("./day19/example.txt")
	raw = raw[:len(raw)-1]
	blueprints := rawToBlueprint(raw)
	ans = bestBlueprint(blueprints)
	return ans
}

func bestBlueprint(blueprints []*Blueprint) (ans int) {
	for _, blueprint := range blueprints {
		mem = make(map[*Blueprint]int)
		ans = utils.Max(ans, bestGeodes(blueprint))
		break
	}
	return ans
}

func bestGeodes(blueprint *Blueprint) (result int) {
	fmt.Println("calculating blueprint: ", blueprint.id)
	time := 0
	result = process(blueprint, time)
	fmt.Println("result: ", result)
	return result
}

func process(blueprint *Blueprint, time int) int {
	if mem[blueprint] != 0 {
		return  mem[blueprint]
	}
	if time > 24 {
		// fmt.Println(len(mem))
		memBlueprint(blueprint, blueprint.resource["geode"])
		return blueprint.resource["geode"]
	}
	// fmt.Println("time: ", time)
	// build
	result := 0
	for _, robot := range blueprint.robots {
		if robot.buildable(blueprint.resource) {
			// fmt.Println("try building robot: ", robot)
			newBlueprint := copyBlueprint(blueprint)
			newBlueprint.build(robot.robotType)
			newBlueprint.updateResource()
			result = utils.Max(result, process(newBlueprint, time+1))
		}
	}
	// if time == 3 {
	// 	fmt.Println("result: ", result)
	// 	panic("debug")
	// }
	// collect
	blueprint.updateResource()
	result = utils.Max(result, process(blueprint, time+1))
	memBlueprint(blueprint, result)
	return result
}

func memBlueprint(blueprint *Blueprint, result int) {
		fmt.Println(&blueprint)
		mem[blueprint] = result
		if mem[blueprint] == 0 {
			mem[blueprint] = -1
		}
}

func mapToArr(a map[string]int) [4]int {
	result := [4]int{}
	for i, k := range [4]string{"ore", "clay", "obsidian", "geode"} {
		result[i] = a[k]
	}
	return result
}

func (blueprint *Blueprint) updateResource() {
	for k := range blueprint.resource {
		blueprint.resource[k] += blueprint.robotCounts[k]
		// fmt.Println("robotCount: ", k, blueprint.robotCounts[k])
		// fmt.Println("resource: ", k, blueprint.resource[k])
	}
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
