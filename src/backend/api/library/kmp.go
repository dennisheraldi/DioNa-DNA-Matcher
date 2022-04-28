package library

// KMP is a string matching algorithm that uses the Knuth-Morris-Pratt algorithm.

/*
	brief
		Mencari indeks pertama dari pattern p di string s
		dengan algoritma Knuth-Morris-Pratt
	param t teks yang akan ditelusuri
	param p pattern yang akan dicari
	return indeks pertama dari pattern p di string s
*/

func KMP(t, p string) int {
	
	var n int = len(t)
	var m int = len(p)

	var fail[] int = computeFail(p);
	var i int = 0
	var j int = 0

	for i < n {
		if t[i] == p[j] {
			if (j == m - 1) {
				return i - m + 1 // match
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j - 1]
		} else {
			i++
		}
	}
	return -1 // no match
}

/*
	brief
		Mencari fail function dari pattern p
	
	param p pattern yang akan dicari
	return fail function dari pattern p

*/
func computeFail(p string) []int {
	var m int = len(p)
	var fail[] int = make([]int, m)
	fail[0] = 0
	
	var j int = 0
	var i int = 1

	for i < m {
		if p[i] == p[j] { // terjadi kecocokan pada j+1
			fail[i] = j + 1
			i++
			j++
		} else if j > 0 { // j cocok dengan prefiks p[0..j-1]
			j = fail[j - 1]
		} else { // tidak terjadi kecocokan
			fail[i] = 0
			i++
		}
	}
	return fail
}