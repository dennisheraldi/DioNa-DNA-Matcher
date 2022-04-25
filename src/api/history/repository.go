package prediksi

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Prediksi, error)
	FindByID(id int) (Prediksi, error)
	Create(prediksi Prediksi) (Prediksi, error)
	Update(prediksi Prediksi) (Prediksi, error)
	Delete(prediksi Prediksi) (Prediksi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([][]Prediksi, error) {
	var prediksis [][]Prediksi
	err := r.db.Find(&prediksis).Error
	return prediksis, err
}

