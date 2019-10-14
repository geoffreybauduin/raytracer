package object

import (
	"math"

	"github.com/geoffreybauduin/raytracer/primitives"
)

type Sphere struct {
	Default

	Center primitives.Vector
	Radius float64
}

func (sphere Sphere) Hit(ray primitives.Ray) *primitives.Hit {
	oc := ray.Origin.Subtract(sphere.Center)
	a := ray.Direction.DotProduct(ray.Direction)
	b := oc.DotProduct(ray.Direction)
	c := oc.DotProduct(oc) - sphere.Radius*sphere.Radius
	discriminant := b*b - a*c
	if discriminant >= 0.0 {
		sqrtDiscriminant := math.Sqrt(discriminant)
		mB := -1.0 * b
		twoA := a
		minus := (mB - sqrtDiscriminant) / twoA
		add := (mB + sqrtDiscriminant) / twoA
		var t float64
		if add < 0.0 {
			t = minus
		} else if minus < 0.0 {
			t = add
		} else {
			t = math.Min(add, minus)
		}
		p := ray.PointAt(t)
		return &primitives.Hit{
			Object:        sphere,
			OriginalColor: sphere.Color,
			P:             p,
			T:             t,
			Normal:        p.Subtract(sphere.Center).Divide(sphere.Radius),
		}
	}
	return nil
}
