package day15

import (
	"strconv"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var pp = message.NewPrinter(language.English)

// 4353354 too low
func Sol1() int {
	pp.Println("Starting Day15 Solution1...")
	raw := utils.ReadFile("./day15/input.txt")
	yGoal := 2_000_000
	// raw := utils.ReadFile("./day15/example.txt")
	// yGoal := 10
	raw = raw[:len(raw)-1]
	map1 := rawToMap(raw)
	points := genPointsAtY(map1, yGoal)
	// ans := answer(points)
	ans := answer2(points)
	return ans
}

func answer2(points [][2]int) (ans int) {
	// mem := make([][2]int, len(points))
	mem := [][2]int{points[0]}
	points = points[1:]
	pp.Println("mem", mem)
	pp.Println("points", points)
	pp.Println()
	for len(points) > 0 {
		// temp := points
		temp := make([][2]int, len(points))
		copy(temp, points)
		for i := len(points) - 1; i >= 0; i-- {
			point := points[i]
			last := &mem[len(mem)-1]
			if calOverlap(*last, point) > 0 {
				last[0] = utils.Min(last[0], point[0])
				last[1] = utils.Max(last[1], point[1])
				temp = remove(temp, i)
			}
		}
		if len(points) != len(temp) {
			points = temp
		} else {
			mem = append(mem, points[0])
			points = points[1:]
		}
		pp.Println("mem", mem)
		pp.Println("points", points)
	}
	for _, point := range mem {
		ans += point[1] - point[0]
	}
	return ans
}

func remove(slice [][2]int, s int) [][2]int {
	return append(slice[:s], slice[s+1:]...)
}

func answer(points [][2]int) (ans int) {
	for i, point := range points {
		// pp.Println("point: ", point)
		overlap := 0
		for j := i - 1; j >= 0; j-- {
			overlap += calOverlap(point, points[j])
		}
		ans += point[1] - point[0] + 1 - overlap
		// pp.Println("ans: ", ans)
		// pp.Println()
	}

	return ans
}

// [16 24] [14 18]
func calOverlap(point1 [2]int, point2 [2]int) (overlap int) {
	// pp.Println("check overlap: ", point1, point2)
	if point1[1] < point2[0] || point2[1] < point1[0] {
		// pp.Println("no overlap")
		return 0
	}
	overlap = utils.Min(point1[1], point2[1]) - utils.Max(point1[0], point2[0]) + 1
	// pp.Println("overlap: ", overlap)

	return overlap
}

func genPointsAtY(map1 []pair, yGoal int) (points [][2]int) {
	for _, pair1 := range map1 {
		// pp.Println()
		distanceToY := distance(pair1.sensor, [2]int{pair1.sensor[0], yGoal})
		// pp.Println(pair1)
		// pp.Println("maxDistance: ", pair1.maxDistance)
		if distanceToY > pair1.maxDistance {
			// pp.Println("skip: ", pair1)
			continue
		}
		midPoint := pair1.sensor[0]
		restDistance := pair1.maxDistance - distanceToY
		data := [2]int{midPoint - restDistance, midPoint + restDistance}
		// pp.Println("range: ", data)
		points = append(points, data)
	}

	return points
}

func distance(point1 [2]int, point2 [2]int) int {
	return utils.Abs(point1[0]-point2[0]) + utils.Abs(point1[1]-point2[1])
}

func rawToMap(raw []string) (map1 []pair) {
	for _, line := range raw {
		splited := strings.Split(line, ":")
		sensor, beacon := parseCoordinate(splited)
		map1 = append(map1, pair{sensor, beacon, distance(sensor, beacon)})
	}
	return map1
}

func parseCoordinate(coor []string) ([2]int, [2]int) {
	sensorString := strings.Split(coor[0][10:], ", ")
	beaconString := strings.Split(coor[1][22:], ", ")
	sensor := [2]int{}
	beacon := [2]int{}
	for i := range sensorString {
		int1, err1 := strconv.Atoi(sensorString[i][2:])
		int2, err2 := strconv.Atoi(beaconString[i][2:])
		utils.Check(err1)
		utils.Check(err2)
		sensor[i] = int1
		beacon[i] = int2
	}
	return sensor, beacon

}

func Map(vs []string, f func(string) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

type pair struct {
	sensor      [2]int
	beacon      [2]int
	maxDistance int
}
