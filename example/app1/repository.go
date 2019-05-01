package app1

// Repository is the interface
type Repository interface {
	NotiPromotion(model Promotion) error
}

type repository struct {
	Ds DataSource
}

// NewRepository is new instance
func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}

func (r *repository) NotiPromotion(model Promotion) error {
	return r.Ds.Publish("promotion", &model)
}
