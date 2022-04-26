package penyakit

type PenyakitRequest struct {
	NamaPenyakit string `json:"nama_penyakit"`
	DNASeq       string `json:"dna_seq"`
}
