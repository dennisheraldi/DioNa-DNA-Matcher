package riwayat

import "time"

type Riwayat struct {
	TanggalPred 	time.Time
	NamaPasien 	 	string
	NamaPenyakit 	string
	Similarity 		float64
	Status 		 	string
}