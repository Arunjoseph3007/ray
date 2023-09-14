package main

import (
	"ray-tracing/camera"
	"ray-tracing/hit"
	"ray-tracing/vec"
)

func main() {
	world := hit.HitList{}

	material_ground := &hit.Lambertian{Albedo: *vec.New(0.8, 0.8, 0.1)}
	material_center := &hit.Lambertian{Albedo: *vec.New(0.7, 0.3, 0.3)}
	material_left := &hit.Metal{Albedo: *vec.New(0.8, 0.8, 0.8), Fuzz: 0.3}
	material_right := &hit.Metal{Albedo: *vec.New(0.8, 0.6, 0.2), Fuzz: .8}

	world.Add(hit.Shpere{Center: *vec.New(0.0, -100.5, -1.0), Radius: 100.0, Material: material_ground})
	world.Add(hit.Shpere{Center: *vec.New(0.0, 0.0, -1.0), Radius: 0.5, Material: material_center})
	world.Add(hit.Shpere{Center: *vec.New(-1.0, 0.0, -1.0), Radius: 0.5, Material: material_left})
	world.Add(hit.Shpere{Center: *vec.New(1.0, 0.0, -1.0), Radius: 0.5, Material: material_right})

	cam := camera.New(
		400,
		16.0/9.0,
		4,
		20,
	)

	cam.Adjust(
		90,
		*vec.New(-2, 2, 1),
		*vec.New(0, 0, -1),
		*vec.New(0, 1, 0),
	)

	cam.Render(world)
}
