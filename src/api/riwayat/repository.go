package riwayat

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Riwayat, error)
	// FindByNamaPenyakit (string) ([]Riwayat, error)
	// FindByTanggalPred (string) ([]Riwayat, error)
	// FindByNamaTanggal (string, string) ([]Riwayat, error)
	Create(riwayat Riwayat) (Riwayat, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Riwayat, error) {
	var Riwayats []Riwayat
	err := r.db.Find(&Riwayats).Error
	return Riwayats, err
}

func (r *repository) Create(riwayat Riwayat) (Riwayat, error) {
	err := r.db.Create(&riwayat).Error
	return riwayat, err
}

