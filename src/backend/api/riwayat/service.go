package riwayat

import "time"

type Service interface {
	FindAll() ([]Riwayat, error)
	FindByPenyakit(string) ([]Riwayat, error)
	FindByTanggal(string) ([]Riwayat, error)
	FindByTanggalPenyakit(string, string) ([]Riwayat, error)
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

func (s *service) FindByPenyakit(nama_penyakit string) ([]Riwayat, error) {
	return s.repository.FindByPenyakit(nama_penyakit)
}

func (s *service) FindByTanggal(tanggal_pred string) ([]Riwayat, error) {
	return s.repository.FindByTanggal(tanggal_pred)
}

func (s *service) FindByTanggalPenyakit(tanggal_pred, nama_penyakit string) ([]Riwayat, error) {
	return s.repository.FindByTanggalPenyakit(tanggal_pred, nama_penyakit)
}

func (s *service) Create(riwayatRequest RiwayatRequest) (Riwayat, error) {
	format := "2006-01-02"
	tanggalPred, _ := time.Parse(format, riwayatRequest.TanggalPred)

	riwayat := Riwayat{
		TanggalPred:  tanggalPred,
		NamaPasien:   riwayatRequest.NamaPasien,
		NamaPenyakit: riwayatRequest.NamaPenyakit,
		Similarity:   riwayatRequest.Similarity,
		Status:       riwayatRequest.Status,
	}
	return s.repository.Create(riwayat)
}