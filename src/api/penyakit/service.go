package penyakit

type Service interface {
	FindAll() ([]Penyakit, error)
	FindByID(id int) (Penyakit, error)
	FindByName(name string) (Penyakit, error)
	Create(penyakitRequest PenyakitRequest) (Penyakit, error)
	Update(ID int, penyakitRequest PenyakitRequest) (Penyakit, error)
	Delete(ID int) (Penyakit, error)
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

func (s *service) FindByID(id int) (Penyakit, error) {
	return s.repository.FindByID(id)
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

func (s *service) Update(ID int, penyakitRequest PenyakitRequest) (Penyakit, error) {
	penyakit, err := s.repository.FindByID(ID)
	if err != nil {
		return penyakit, err
	}
	penyakit.NamaPenyakit = penyakitRequest.NamaPenyakit
	penyakit.DNAPenyakit = penyakitRequest.DNAPenyakit

	return s.repository.Update(penyakit)
}

func (s *service) Delete(ID int) (Penyakit, error) {
	penyakit, err := s.repository.FindByID(ID)
	if err != nil {
		return penyakit, err
	}
	return s.repository.Delete(penyakit)
}