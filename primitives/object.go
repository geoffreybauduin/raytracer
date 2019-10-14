package primitives

type Object interface {
	Hit(Ray) *Hit
	GetMaterial() Material
}
