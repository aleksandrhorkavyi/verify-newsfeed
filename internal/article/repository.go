package article

type Repository interface {
	All() []*Article
}

type repository struct {
}

func (r repository) All() []*Article {
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
