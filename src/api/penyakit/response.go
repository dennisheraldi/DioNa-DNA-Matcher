package penyakit

type PenyakitResponse struct {
	NamaPenyakit string `json:"nama_penyakit"`
	DNASeq       string `json:"dna_seq"`
}