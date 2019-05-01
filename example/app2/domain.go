package app2

// UseCase is the interface
type UseCase interface {
	SubNotiPromotion(callback func(model *Promotion)) error
}

type useCase struct {
	Repo Repository
}

// NewUseCase is new instance
func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}

func (uc *useCase) SubNotiPromotion(callback func(model *Promotion)) error {
	return uc.Repo.SubNotiPromotion(callback)
}
