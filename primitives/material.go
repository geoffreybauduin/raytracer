package primitives

type Material interface {
	GetAttenuation() Vector
	Scatter(Ray, *Hit) *Ray
}
