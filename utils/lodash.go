package utils

func Contains[T int | rune](s []T, e T) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
