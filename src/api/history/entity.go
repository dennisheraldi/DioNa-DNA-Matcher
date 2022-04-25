package prediksi

import "time"

type Prediksi struct {
	ID           int
	Tanggal 	 time.Time
	NamaPasien 	 string
	NamaPenyakit string
	Status 		 string
}