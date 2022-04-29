package library

func CheckDNA(pasien, penyakit string) (string, float64) {
	if KMP(pasien, penyakit) != -1 { // Pengecekan terlebih dahulu dengan string matching metode KMP
		return "True", 100.00
	} else { // Jika tidak cocok, cek similarity dengan metode LCS
		similarity := LcsResult(pasien, penyakit)
		if similarity >= 80 { // Jika similarity lebih dari 80%, maka cocok
			return "True", similarity
		} else {
			return "False", similarity
		}
	}
}