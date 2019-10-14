package engine

import (
	"image/color"
	"math/rand"

	"github.com/gammazero/workerpool"
	"github.com/geoffreybauduin/raytracer/material"
	"github.com/geoffreybauduin/raytracer/object"
	"github.com/geoffreybauduin/raytracer/primitives"
	"github.com/schollz/progressbar"
)

type Scene struct {
	Width, Height   int
	Camera          *Camera
	Objects         []primitives.Object
	BackgroundColor color.RGBA

	antiAliasingFactor int
}

func NewScene(width, height int, backgroundColor color.RGBA) Scene {
	scene := Scene{
		Width:  width,
		Height: height,
		Camera: NewCamera(width, height),
		Objects: []primitives.Object{
			object.Sphere{
				Center: primitives.Vector{X: 0, Y: 0, Z: -1},
				Radius: 0.5,
				Default: object.Default{
					Color:    color.RGBA{255, 0, 0, 255},
					Material: material.Metal{Attenuation: primitives.Vector{0.8, 0.3, 0.3}},
				},
			},
			object.Sphere{
				Center: primitives.Vector{X: 0, Y: -100.5, Z: -1},
				Radius: 100,
				Default: object.Default{
					Color:    color.RGBA{255, 0, 0, 255},
					Material: material.Metal{Attenuation: primitives.Vector{0.8, 0.8, 0.0}},
				},
			},
			object.Sphere{
				Center: primitives.Vector{X: 1, Y: 0, Z: -1},
				Radius: 0.5,
				Default: object.Default{
					Color:    color.RGBA{255, 0, 0, 255},
					Material: material.Metal{Attenuation: primitives.Vector{0.8, 0.6, 0.2}},
				},
			},
			object.Sphere{
				Center: primitives.Vector{X: -1, Y: 0, Z: -1},
				Radius: 100,
				Default: object.Default{
					Color:    color.RGBA{255, 0, 0, 255},
					Material: material.Metal{Attenuation: primitives.Vector{0.8, 0.8, 0.8}},
				},
			},
		},
		BackgroundColor:    backgroundColor,
		antiAliasingFactor: 10,
	}
	return scene
}

func (s Scene) Render() error {
	pool := workerpool.New(16)
	bar := progressbar.New(s.Height * s.Width)
	for y := s.Height - 1; y >= 0; y-- {
		for x := 0; x < s.Width; x++ {
			xx := x
			yy := y
			pool.Submit(func() {
				s.renderAt(xx, yy)
				bar.Add(1)
			})
		}
	}
	pool.StopWait()
	return s.Camera.Snapshot()
}

func (s Scene) renderAt(x, y int) {
	var colorVec primitives.Vector
	for d := 0; d < s.antiAliasingFactor; d++ {
		ray := s.Camera.RayAt(float64(x)+rand.Float64(), float64(y)+rand.Float64())
		colorVec = colorVec.Add(s.getColorWithRay(ray, 0))
	}
	colorVec = colorVec.Divide(float64(s.antiAliasingFactor))
	s.Camera.SetColor(x, y, colorVec.ToColor())
}

func (s Scene) getColorWithRay(ray primitives.Ray, depth int) primitives.Vector {
	var recordedHit *primitives.Hit
	for _, object := range s.Objects {
		if hit := object.Hit(ray); hit != nil {
			if hitIsBetterThan(hit, recordedHit) {
				recordedHit = hit
			}
		}
	}
	if recordedHit == nil {
		recordedHit = &primitives.Hit{OriginalColor: s.BackgroundColor}
		return recordedHit.GetColorVector(ray)
	}
	if depth < 5 {
		mat := recordedHit.Object.GetMaterial()
		scatteredRay := mat.Scatter(ray, recordedHit)
		if scatteredRay != nil {
			attenuation := mat.GetAttenuation()
			colorVec := s.getColorWithRay(*scatteredRay, depth+1)
			return colorVec.Multiply(attenuation)
		}
	}
	return recordedHit.GetColorVector(ray)
}

func hitIsBetterThan(h1 *primitives.Hit, current *primitives.Hit) bool {
	if h1.T < 0.0001 {
		return false
	} else if current == nil {
		return true
	}
	return h1.T < current.T
}
