package penyakit

type PenyakitRequest struct {
	NamaPenyakit string `json:"nama_penyakit"`
	DNAPenyakit  string `json:"dna_penyakit"`
}
