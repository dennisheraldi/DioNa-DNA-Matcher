package penyakit

type PenyakitRequest struct {
	Nama      string `json:"nama"		validate:"required"`
	RantaiDNA string `json:"rantai_dna"	validate:"required"`
}
