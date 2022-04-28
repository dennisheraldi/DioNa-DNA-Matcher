package penyakit

type Service interface {
	FindAll() ([]Penyakit, error)
	FindByName(name string) (Penyakit, error)
	Create(penyakitRequest PenyakitRequest) (Penyakit, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Penyakit, error) {
	return s.repository.FindAll()
}

func (s *service) FindByName(name string) (Penyakit, error) {
	return s.repository.FindByName(name)
}

func (s *service) Create(penyakitRequest PenyakitRequest) (Penyakit, error) {
	penyakit := Penyakit{
		NamaPenyakit: penyakitRequest.NamaPenyakit,
		DNAPenyakit:  penyakitRequest.DNAPenyakit,
	}
	return s.repository.Create(penyakit)
}