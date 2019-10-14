package object

import (
	"image/color"

	"github.com/geoffreybauduin/raytracer/primitives"
)

type Default struct {
	Color    color.RGBA
	Material primitives.Material
}

func (d Default) GetMaterial() primitives.Material {
	return d.Material
}
