package app2

// Repository is the interface
type Repository interface {
	SubNotiPromotion(callback func(model *Promotion)) error
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

func (r *repository) SubNotiPromotion(callback func(model *Promotion)) error {
	_, err := r.Ds.Subscribe("promotion", callback)
	return err
}
