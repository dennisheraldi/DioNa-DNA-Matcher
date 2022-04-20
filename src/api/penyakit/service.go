package penyakit

type Service interface {
	FindAll() ([]Penyakit, error)
	FindByID(id int) (Penyakit, error)
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

func (s *service) Create(penyakitRequest PenyakitRequest) (Penyakit, error) {
	penyakit := Penyakit{
		Nama:      penyakitRequest.Nama,
		RantaiDNA: penyakitRequest.RantaiDNA,
	}
	return s.repository.Create(penyakit)
}

func (s *service) Update(ID int, penyakitRequest PenyakitRequest) (Penyakit, error) {
	penyakit, err := s.repository.FindByID(ID)
	if err != nil {
		return penyakit, err
	}
	penyakit.Nama = penyakitRequest.Nama
	penyakit.RantaiDNA = penyakitRequest.RantaiDNA

	return s.repository.Update(penyakit)
}

func (s *service) Delete(ID int) (Penyakit, error) {
	penyakit, err := s.repository.FindByID(ID)
	if err != nil {
		return penyakit, err
	}
	return s.repository.Delete(penyakit)
}