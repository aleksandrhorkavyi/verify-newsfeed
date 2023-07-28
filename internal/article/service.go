package article

type Service interface {
	All() []*Article
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) All() []*Article {
	return nil
}
