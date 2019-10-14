package engine

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/geoffreybauduin/raytracer/primitives"
)

type Camera struct {
	Width, Height int
	Horizontal    primitives.Vector
	Vertical      primitives.Vector
	Origin        primitives.Vector
	Image         *image.RGBA
	LowerLeft     primitives.Vector
}

// NewCamera creates a new camera
func NewCamera(width, height int) *Camera {
	camera := &Camera{
		Width:      width,
		Height:     height,
		Horizontal: primitives.Vector{X: 4.0, Y: 0, Z: 0},
		Vertical:   primitives.Vector{X: 0, Y: 2.0, Z: 0},
		Origin:     primitives.Vector{X: 0.0, Y: 0.0, Z: 0.0},
		LowerLeft:  primitives.Vector{X: -2.0, Y: -1.0, Z: -1.0},
		Image:      image.NewRGBA(image.Rect(0, 0, width, height)),
	}
	return camera
}

// Snapshot creates a file ./result.png with the current view of the camera
func (c *Camera) Snapshot() error {
	f, err := os.Create("./result.png")
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, c.Image)
}

// RayAt returns a ray
func (c Camera) RayAt(x, y float64) primitives.Ray {
	u := x / float64(c.Width)
	v := y / float64(c.Height)

	horizontal := c.Horizontal.MultiplyScalar(u)
	vertical := c.Vertical.MultiplyScalar(v)
	direction := c.LowerLeft.Add(horizontal).Add(vertical).Subtract(c.Origin)
	return primitives.Ray{
		Origin:    c.Origin,
		Direction: direction,
	}
}

// SetColor sets the color of a point on the camera
func (c Camera) SetColor(x, y int, color_ color.RGBA) {
	c.Image.SetRGBA(x, y, color_)
}
