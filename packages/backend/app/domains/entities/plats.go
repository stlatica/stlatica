package entities

// Plat is the entity for plat.
type Plat struct {
	PlatBase
	Images []string
}

// GetImages returns images.
func (p *Plat) GetImages() []string {
	return p.Images
}
