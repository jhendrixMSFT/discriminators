package farm

// FishType provides polymorphic access to various types of fish.
type FishType interface {
	GetFish() *Fish
}

// Fish is the "base" discriminated type.
type Fish struct {
	ID       int
	FishType string // discriminator
}

// GetFish implements the FishType interface.
func (f *Fish) GetFish() *Fish {
	return f
}

// Tuna inherits Fish.
type Tuna struct {
	Fish
	Color string
}

// SalmonType provides polymorphic access to various types of salmon.
type SalmonType interface {
	FishType
	GetSalmon() *Salmon
}

// Salmon inherits Fish.
type Salmon struct {
	Fish
	Breed string
}

// GetSalmon implements the SalmonType interface.
func (f *Salmon) GetSalmon() *Salmon {
	return f
}

// SmartSalmon inherits Salmon.
type SmartSalmon struct {
	Salmon
	IQ int
}

// GetFish returns some type of fish.
func GetFish() FishType {
	return &Tuna{
		Color: "white",
		Fish: Fish{
			ID:       42,
			FishType: "salmon",
		},
	}
}

// GetSalmon returns some type of salmon.
func GetSalmon() SalmonType {
	return &SmartSalmon{
		Salmon: Salmon{
			Breed: "king",
			Fish: Fish{
				ID:       98,
				FishType: "smart_salmon",
			},
		},
		IQ: 151,
	}
}
