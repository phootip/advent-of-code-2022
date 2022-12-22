package utils

import "strconv"

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}


func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Contains[T int | rune | string](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// func sliceToSet[T string | int](s []T) (result map[T]bool) {
// 	for 
	
// }

func Filter[T int | rune | string](s []T, e T) []T {
	s2 := make([]T, len(s))
	copy(s2, s)
	for i, a := range s {
		if a == e {
			return append(s2[:i],s2[i+1:]...)
		}
	}
	return s2
}

func Difference(slice1 []string, slice2 []string) []string{
	result := make([]string,0)
	for _,s := range slice1 {
		if !Contains(slice2, s) {
			result = append(result, s)
		}
	}
	return result
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
			if a == e {
					return true
			}
	}
	return false
}

func ReverseS(s string) (result string) {
	for _, r := range s {
		result = string(r) + result
	}
	return result
}

func SToInt(ss []string) []int {
	result := make([]int, len(ss))
	for i, s := range ss {
		r, err := strconv.Atoi(s)
		Check(err)
		result[i] = r
	}
	return result
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func SliceEqual[T [2]int](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i][0] != slice2[i][0] || slice1[i][1] != slice2[i][1] {
			return false
		}
	}
	return true
}

func CopyMap[T map[string]int](map1 T, map2 T) T{
	for k,v := range map2 {
		map1[k] = v
	}
	return map1
}

func SliceIndex[T int](slice []T, item T) int {
	for i, item2 := range slice {
		if item == item2 {
			return i
		}
	}
	return -1
}
