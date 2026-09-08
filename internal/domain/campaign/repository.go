package campaign

type Repository interface {
	Save(campaign *Campaign) error
	// FindByID(id string) (*Campaign, error)
}
