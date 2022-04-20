package penyakit

type PenyakitResponse struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	RantaiDNA string `json:"rantai_dna"`
}