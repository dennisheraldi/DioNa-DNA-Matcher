package riwayat

import "time"

type Riwayat struct {
	TanggalPred 	time.Time
	NamaPasien 	 	string
	DNAPasien 		string
	NamaPenyakit 	string
	Status 		 	string
}