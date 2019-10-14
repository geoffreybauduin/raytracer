package main

import (
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/geoffreybauduin/raytracer/engine"
	"github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	scene := engine.NewScene(600, 300, color.RGBA{0, 0, 255, 255})
	if err := scene.Render(); err != nil {
		panic(err)
	}
}
