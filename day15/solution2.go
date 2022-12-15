package day15

import "github.com/phootip/advent-of-code-2022/utils"

// 8168290671045 too low
// 13543690671045
func Sol2() (ans int) {
	pp.Println("Starting Day15 Solution2...")
	raw := utils.ReadFile("./day15/input.txt")
	limit := 4_000_000
	// raw := utils.ReadFile("./day15/example.txt")
	// limit := 20
	raw = raw[:len(raw)-1]
	map1 := rawToMap(raw)
	pp.Println(map1)
	beacon := findBeacon2(map1,limit)
	// beacon := [2]int{3385922, 2671045}
	pp.Println(beacon)
	return beacon[0]*4000000 + beacon[1]

}

func findBeacon(map1 []pair, limit int) (beacon [2]int) {
	for i:=0; i <= limit; i++ {
		for j:=0; j <= limit; j++ {
			if i % 100_000 == 0 && i != 0 {
				pp.Println(i,j)
			}
			if !reachable([2]int{i,j}, map1){
				return [2]int{i,j}
			}
		}
	}
	return beacon
}

func findBeacon2(map1 []pair, limit int) (beacon [2]int) {
	for i:=0; i <= limit; i++ {
		result := mergeRange(genPointsAtY(map1, i))
		// pp.Println(i)
		// pp.Println("result: ",result)
		if len(result) > 1 {
			if result[0][1] > result[1][0] {
				result[0], result[1] = result[1], result[0]
			}
			return [2]int{(result[0][1] + result[1][0])/2,i}
		}
	}
	return beacon
}

func mergeRange(points [][2]int) [][2]int {
	// mem := make([][2]int, len(points))
	mem := [][2]int{points[0]}
	points = points[1:]
	// pp.Println("mem", mem)
	// pp.Println("points", points)
	// pp.Println()
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
		// pp.Println("mem", mem)
		// pp.Println("points", points)
	}
	return mem
}

func reachable(now [2]int, points []pair) bool {
	for _, point := range points {
		if point.maxDistance >= distance(now, point.sensor) {
			return true
		}
	}
	return false
}

// func rawToMap(raw []string) (map1 []pair) {
// 	for _, line := range raw {
// 		splited := strings.Split(line, ":")
// 		sensor, beacon := parseCoordinate(splited)
// 		map1 = append(map1, pair{sensor, beacon, distance(sensor, beacon)})
// 	}
// 	return map1
// }
