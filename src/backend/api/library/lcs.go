package library

import "math"

func LcsCount(p string, q string, x int, y int) float64 {
	if x == 0 || y == 0 {
		return 0
	} else if p[x-1] == q[y-1] {
		return 1 + (LcsCount(p, q, x-1, y-1))
	} else {
		n := LcsCount(p, q, x-1, y)
		m := LcsCount(p, q, x, y-1)
		if n > m {
			return n
		} else {
			return m
		}
	}
}

func LcsResult(text, pattern string) float64 {
	p := pattern
	q := text
	var result float64 = LcsCount(p, q, len(p), len(q)) / float64(len(p)) * 100
	return math.Round(result*100) / 100
}
