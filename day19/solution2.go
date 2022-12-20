package day19

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/phootip/advent-of-code-2022/utils"
)

var uniTime int
var globalBest = 0

func Sol2() (ans int) {
	fmt.Println("Starting Day19 Solution2...")
	allType = [4]string{"ore", "clay", "obsidian", "geode"}
	raw := utils.ReadFile("./day19/input.txt")
	// raw := utils.ReadFile("./day19/example.txt")
	uniTime = 32
	// uniTime = 28
	// uniTime = 24
	raw = raw[:len(raw)-1]
	blueprints := rawToBlueprint(raw)
	// blueprints = blueprints[:1]
	blueprints = blueprints[:3]
	ans = bestBlueprint2(blueprints)
	return ans
}

func bestBlueprint2(blueprints []*Blueprint) (ans int) {
	ans = 1
	for _, blueprint := range blueprints {
		globalBest = 0
		// ans += bestGeodes2(blueprint)*blueprint.id
		ans *= bestGeodes2(blueprint)
		fmt.Println(ans)
	}
	return ans
}
func bestGeodes2(blueprint *Blueprint) (result int) {
	fmt.Println("calculating blueprint: ", blueprint.id)
	// fmt.Println("blueprint: ", blueprint)
	start := time.Now()
	debugPath := []string{"clay", "clay", "obsidian", "clay", "obsidian", "geode", "geode"}
	// debugPath := []string{"ore","clay","clay","clay","clay","clay","clay", "obsidian","obsidian","obsidian","obsidian", "clay", "obsidian", "geode", "obsidian", "geode","obsidian", "geode"}
	_ = debugPath
	mem = make(map[string]int)
	result, bestPath := process2(blueprint, 0, "clay", []string{}, debugPath)
	fmt.Println("bestPath: ", bestPath)
	fmt.Println("len(mem):", len(mem))
	fmt.Println("result1: ", result)
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
	mem = make(map[string]int)
	result2, bestPath2 := process2(blueprint, 0, "ore", []string{}, debugPath)
	fmt.Println("bestPath2: ", bestPath2)
	fmt.Println("len(mem):", len(mem))
	fmt.Println("result2: ", result2)
	elapsed = time.Since(start)
	log.Printf("Execution took %s", elapsed)
	_ = bestPath
	// return result
	_ = bestPath2
	return utils.Max(result, result2)
}

func process2(oldBlueprint *Blueprint, time int, target string, path []string, debugPath []string) (int, []string) {
	if mem[strings.Join(path, ",")] != 0 {
		// fmt.Println("path already cal: ",path)
		return mem[strings.Join(path, ",")], path
	}
	currentBest := oldBlueprint.assumeBest(uniTime-time)
	if currentBest < globalBest {
		mem[strings.Join(path, ",")] = -1
		return -1, path
	}
	blueprint := copyBlueprint(oldBlueprint)
	// fmt.Println("time: ", time)
	result := 0
	robot := blueprint.robots[target]
	waitTime := robot.timeTillBuildable(blueprint)
	endtime := time + waitTime
	timeLeft := uniTime - endtime
	// fmt.Println("target: ", target)
	// fmt.Println("path: ", path)
	// fmt.Println("debugPath: ", debugPath)
	// fmt.Println("waitTime: ", waitTime)
	if endtime > uniTime {
		result = blueprint.resource["geode"]
		mem[strings.Join(path, ",")] = result
		if result == 0 {
			mem[strings.Join(path, ",")] = -1
		}
		globalBest = utils.Max(globalBest, result)
		return result, path
	}
	blueprint.updateResource2(waitTime)
	blueprint.build2(target, timeLeft)
	path = append(path, target)

	// fmt.Println("endtime: ", endtime)
	// blueprint.debugResource()
	// if len(debugPath) > 0 {
	// 	newBlueprint := copyBlueprint(blueprint)
	// 	nextTarget := debugPath[0]
	// 	debugPath = debugPath[1:]
	// 	return process2(newBlueprint, endtime, nextTarget, path, debugPath)
	// }
	// if len(debugPath) == 0 {
	// 	panic("debug")
	// 	return process2(blueprint, endtime, "geode", path, debugPath)
	// }
	// panic("don't come here")
	bestPath := []string{}
	_ = bestPath
	nextBuilds := blueprint.getNextBuilds(uniTime - endtime)
	for _, t := range nextBuilds {
		newBlueprint := copyBlueprint(blueprint)
		newResult, newPath := process2(newBlueprint, endtime, t, path, debugPath)
		if newResult >= result {
			result = newResult
			bestPath = newPath
		}
	}
	mem[strings.Join(path, ",")] = result
	if result == 0 {
		mem[strings.Join(path, ",")] = -1
	}
	globalBest = utils.Max(globalBest, result)
	return result, bestPath
}


func (blueprint *Blueprint) assumeBest(timeLeft int) int {
	currentBest := blueprint.resource["geode"]
	for timeLeft > 0 {
		currentBest += timeLeft
		timeLeft--
	}
	return currentBest
}

func (blueprint *Blueprint) build2(robotType string, timeLeft int) {
	blueprint.robotCounts[robotType] += 1
	robot := blueprint.robots[robotType]
	for k := range robot.costs {
		blueprint.resource[k] -= robot.costs[k]
	}
	if robotType == "geode"{
		// fmt.Println("updating geode:", timeLeft)
		blueprint.resource["geode"] += timeLeft
	}
}

func (blueprint *Blueprint) updateResource2(waitTime int) {
	for k := range blueprint.resource {
		if k == "geode" {
			continue
		}
		blueprint.resource[k] += blueprint.robotCounts[k] * waitTime
	}
}

func (blueprint *Blueprint) getNextBuilds(timeLeft int) []string {
	result := []string{}
	// result := []string{"geode"}
	for _, t := range []string{"ore","clay","obsidian"} {
		if blueprint.robotCounts[t] >= blueprint.robotLimits[t] {
			continue
		}
		if blueprint.robotCounts[t]*timeLeft + blueprint.resource[t] >= timeLeft*blueprint.robotLimits[t] {
			continue
		}
		result = append(result, t)
	}
	result = append(result, "geode")
	return result
}
