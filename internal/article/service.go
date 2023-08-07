package article

type Service interface {
	DoSomething() map[string]any
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) DoSomething() map[string]any {
	return s.repo.One()
}
