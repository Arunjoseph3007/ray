package main

import (
	"ray-tracing/camera"
	"ray-tracing/hit"
	"ray-tracing/vec"
)

func main() {
	world := hit.HitList{}
	world.Add(hit.Shpere{
		Center: *vec.New(0, 0, -1),
		Radius: .5,
	})
	world.Add(hit.Shpere{
		Center: *vec.New(0, -100.5, -1),
		Radius: 100,
	})

	cam := camera.Camera{}

	cam.Initialize()
	cam.Render(world)
}
