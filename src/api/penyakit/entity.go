package penyakit

type Penyakit struct {
	NamaPenyakit string `gorm: "primaryKey"`
	DNASeq       string
}