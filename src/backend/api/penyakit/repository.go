package penyakit

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Penyakit, error)
	FindByName(name string) (Penyakit, error)
	Create(penyakit Penyakit) (Penyakit, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Penyakit, error) {
	var penyakits []Penyakit
	err := r.db.Find(&penyakits).Error
	return penyakits, err
}

func (r *repository) FindByName(name string) (Penyakit, error) {
	var penyakit Penyakit
	err := r.db.Where("nama_penyakit = ?", name).Find(&penyakit).Error
	return penyakit, err
}

func (r *repository) Create(penyakit Penyakit) (Penyakit, error) {
	err := r.db.Create(&penyakit).Error
	return penyakit, err
}