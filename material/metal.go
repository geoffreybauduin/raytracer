package material

import (
	"github.com/geoffreybauduin/raytracer/primitives"
)

type Metal struct {
	Attenuation primitives.Vector
}

func (metal Metal) GetAttenuation() primitives.Vector {
	return metal.Attenuation
}

func (metal Metal) Scatter(ray primitives.Ray, hit *primitives.Hit) *primitives.Ray {
	unit := ray.Direction.Unit()
	reflected := metal.reflect(unit, hit.Normal)
	dot := reflected.DotProduct(hit.Normal)
	if dot > 0.0 {
		return &primitives.Ray{Origin: hit.P, Direction: reflected}
	}
	return nil
}

func (_ Metal) reflect(v1, v2 primitives.Vector) primitives.Vector {
	return v1.Subtract(v2.MultiplyScalar(2 * v1.DotProduct(v2)))
}
