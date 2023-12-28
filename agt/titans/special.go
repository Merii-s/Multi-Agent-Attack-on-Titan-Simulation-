package titans

type SpecialTitanI interface {
	TitanI
	Transform()
	Capacity()
}

type SpecialTitan struct {
	attributes Titan
}
