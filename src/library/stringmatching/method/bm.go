package method

import (
	"math"
)

// BM is a string matching algorithm that uses the Boyer-Moore algorithm.

/*
	brief
		Mencari indeks pertama dari pattern p di string s
		dengan algoritma Boyer-Moore
	param t teks yang akan ditelusuri
	param p pattern yang akan dicari
	return indeks pertama dari pattern p di string s
*/

func BM(t, p string) int {
	var last[] int = buildLast(p)
	var n int = len(t)
	var m int = len(p)
	var i int = m-1

	if (i > n-1) { // tidak ada match jika panjang pattern lebih besar dari panjang string
		return -1
	}

	var j int = m-1

	for f := true; f; f = (i <= n-1) {
		if p[j] == t[i] {
			if j == 0 {
				return i // match
			} // looking-glass technique
			i--
			j--
		} else { // character jump technique
			var lo int = last[t[i]]
			i = i + m - int(math.Min(float64(j), float64(1+lo)))
			j = m - 1
		}
	}

	return -1 // no match
	
}

/*
	brief
		Mencari last function dari pattern p
	param 
		p pattern yang akan dicari
	return
		array storing index of last occurence 
		of each ASCII char in pattern p
*/
func buildLast(p string) []int {
	var m int = len(p)
	var last[] int = make([]int, 128)
	for i := 0; i < 128; i++ { // inisialisasi array last
		last[i] = -1
	}
	for i := 0; i < m; i++ {
		last[p[i]] = i
	}
	return last
}