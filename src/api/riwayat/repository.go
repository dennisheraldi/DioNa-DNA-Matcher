package riwayat

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Riwayat, error)
	FindByPenyakit(string) ([]Riwayat, error)
	FindByTanggal(string) ([]Riwayat, error)
	FindByTanggalPenyakit(string, string) ([]Riwayat, error)
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

func (r *repository) FindByPenyakit(penyakit string) ([]Riwayat, error) {
	var Riwayats []Riwayat
	err := r.db.Where("nama_penyakit = ?", penyakit).Find(&Riwayats).Error
	return Riwayats, err
}

func (r *repository) FindByTanggal(tanggal string) ([]Riwayat, error) {
	var Riwayats []Riwayat
	err := r.db.Where("tanggal_pred = ?", tanggal).Find(&Riwayats).Error
	return Riwayats, err
}

func (r *repository) FindByTanggalPenyakit(tanggal string, penyakit string) ([]Riwayat, error) {
	var Riwayats []Riwayat
	err := r.db.Where("tanggal_pred = ? AND nama_penyakit = ?", tanggal, penyakit).Find(&Riwayats).Error
	return Riwayats, err
}

func (r *repository) Create(riwayat Riwayat) (Riwayat, error) {
	err := r.db.Create(&riwayat).Error
	return riwayat, err
}

