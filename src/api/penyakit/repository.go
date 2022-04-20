package penyakit

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Penyakit, error)
	FindByID(id int) (Penyakit, error)
	Create(penyakit Penyakit) (Penyakit, error)
	Update(penyakit Penyakit) (Penyakit, error)
	Delete(penyakit Penyakit) (Penyakit, error)
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

func (r *repository) FindByID(id int) (Penyakit, error) {
	var penyakit Penyakit
	err := r.db.First(&penyakit, id).Error
	return penyakit, err
}

func (r *repository) Create(penyakit Penyakit) (Penyakit, error) {
	err := r.db.Create(&penyakit).Error
	return penyakit, err
}

func (r *repository) Update(penyakit Penyakit) (Penyakit, error) {
	err := r.db.Save(&penyakit).Error
	return penyakit, err
}

func (r *repository) Delete(penyakit Penyakit) (Penyakit, error) {
	err := r.db.Delete(&penyakit).Error
	return penyakit, err
}