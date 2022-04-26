package riwayat

type RiwayatRequest struct {
	TanggalPred  string `json:"tanggal_pred"`
	NamaPasien   string `json:"nama_pasien"`
	NamaPenyakit string `json:"nama_penyakit"`
	Status       string `json:"status"`
}

type RiwayatSubmit struct {
	NamaPasien   string `json:"nama_pasien"`
	DNAPasien    string `json:"dna_pasien"`
	NamaPenyakit string `json:"nama_penyakit"`
}