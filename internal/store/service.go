package store

type Service interface {
	Get(key string) (string, bool)
	Put(key string, value string)
	Remove(key string)
	Flush() error
	Init() error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return service{repo: repository}
}

func (s service) Get(key string) (string, bool) {
	return s.repo.Get(key)
}

func (s service) Put(key string, value string) {
	s.repo.Put(key, value)
}

func (s service) Remove(key string) {
	s.repo.Remove(key)
}

func (s service) Flush() error {
	return s.repo.Flush()
}

func (s service) Init() error {
	return s.repo.Init()
}
