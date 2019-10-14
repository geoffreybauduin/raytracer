package primitives

type Ray struct {
	Origin    Vector
	Direction Vector
}

func (ray Ray) PointAt(t float64) Vector {
	return ray.Origin.Add(ray.Direction.MultiplyScalar(t))
}
