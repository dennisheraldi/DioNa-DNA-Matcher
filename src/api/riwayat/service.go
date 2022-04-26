package riwayat

import "time"

type Service interface {
	FindAll() ([]Riwayat, error)
	// FindByNamaPenyakit(string) ([]Riwayat, error)
	// FindByTanggalPred(string) ([]Riwayat, error)
	// FindByNamaTanggal(string, string) ([]Riwayat, error)
	Create(riwayatRequest RiwayatRequest) (Riwayat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Riwayat, error) {
	return s.repository.FindAll()
}

func (s *service) Create(riwayatRequest RiwayatRequest) (Riwayat, error) {
	format := "2006-01-02"
	tanggalPred, _ := time.Parse(format, riwayatRequest.TanggalPred)

	riwayat := Riwayat{
		TanggalPred:  tanggalPred,
		NamaPasien:   riwayatRequest.NamaPasien,
		NamaPenyakit: riwayatRequest.NamaPenyakit,
		Status:       riwayatRequest.Status,
	}
	return s.repository.Create(riwayat)
}