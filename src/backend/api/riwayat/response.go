package riwayat

type RiwayatResponse struct {
	TanggalPred  string  `json:"tanggal_pred"`
	NamaPasien   string  `json:"nama_pasien"`
	NamaPenyakit string  `json:"nama_penyakit"`
	Similarity   float64 `json:"similarity"`
	Status       string  `json:"status"`
}