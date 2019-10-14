package primitives

import (
	"image/color"
)

type Hit struct {
	OriginalColor color.RGBA

	T      float64
	P      Vector
	Normal Vector
	Object Object
}

func (h *Hit) GetColorVector(ray Ray) Vector {
	if h.T > 0.0 {
		return h.ratioHitPoint(ray)
	}
	return h.ratioUnitVector(ray)
}

func (h *Hit) ratioUnitVector(ray Ray) Vector {
	unit := ray.Direction.Unit()
	t := 0.5 * (unit.Y + 1.0)
	v2 := Vec1.MultiplyScalar(1.0 - t)
	v3 := Vector{0.5, 0.7, 1.0}
	v := v2.Add(v3.MultiplyScalar(t))
	return v
}

func (h *Hit) ratioHitPoint(ray Ray) Vector {
	return h.Normal.Add(Vec1).MultiplyScalar(0.5)
}
