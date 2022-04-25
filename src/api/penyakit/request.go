package penyakit

type PenyakitRequest struct {
	NamaPenyakit string `json:"nama_penyakit"		validate:"required"`
	DNASeq       string `json:"dna_seq"	validate:"required"`
}
