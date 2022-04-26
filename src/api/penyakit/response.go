package penyakit

type PenyakitResponse struct {
	NamaPenyakit string `json:"nama_penyakit"`
	DNAPenyakit  string `json:"dna_penyakit"`
}