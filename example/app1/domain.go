package app1

// UseCase is the interface
type UseCase interface {
	NotiPromotion(model Promotion) error
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

func (uc *useCase) NotiPromotion(model Promotion) error {
	return uc.Repo.NotiPromotion(model)
}
