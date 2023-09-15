package main

import (
	"math/rand"
	"ray-tracing/camera"
	"ray-tracing/hit"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

func getTestWorld() hit.HitList {
	world := hit.HitList{}

	material_ground := &hit.Lambertian{Albedo: *vec.New(0.8, 0.8, 0.0)}
	material_center := &hit.Lambertian{Albedo: *vec.New(0.1, 0.2, 0.5)}
	material_left := &hit.Dielectric{RefractIndex: 1.5}
	material_right := &hit.Metal{Albedo: *vec.New(0.8, 0.6, 0.2), Fuzz: 0.0}

	world.Add(hit.Sphere{Center: *vec.New(0.0, -100.5, -1.0), Radius: 100.0, Material: material_ground})
	world.Add(hit.Sphere{Center: *vec.New(0.0, 0.0, -1.0), Radius: 0.5, Material: material_center})
	world.Add(hit.Sphere{Center: *vec.New(-1.0, 0.0, -1.0), Radius: 0.5, Material: material_left})
	world.Add(hit.Sphere{Center: *vec.New(-1.0, 0.0, -1.0), Radius: -0.4, Material: material_left})
	world.Add(hit.Sphere{Center: *vec.New(1.0, 0.0, -1.0), Radius: 0.5, Material: material_right})
	// cam := camera.New(
	// 	400,
	// 	16.0/9.0,
	// 	8,
	// 	20,
	// )

	// cam.Adjust(
	// 	20,
	// 	*vec.New(-2, 2, 1),
	// 	*vec.New(0, 0, -1),
	// 	*vec.New(0, 1, 0),
	// )

	return world
}

func getAwesomeWorld() hit.HitList {
	world := hit.HitList{}

	ground_material := &hit.Lambertian{Albedo: *vec.New(0.5, 0.5, 0.5)}
	world.Add(&hit.Sphere{Center: *vec.New(0, -1000, 0), Radius: 1000, Material: ground_material})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			choose_mat := rand.Float64()
			center := vec.New(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if vec.Sub(*center, *vec.New(4, 0.2, 0)).Length() > 0.9 {

				if choose_mat < 0.8 {
					// diffuse
					albedo := vec.Random(-1, 1)
					sphere_material := &hit.Lambertian{Albedo: albedo}
					world.Add(hit.Sphere{Center: *center, Radius: 0.2, Material: sphere_material})
				} else if choose_mat < 0.95 {
					// metal
					albedo := vec.Random(0.5, 1)
					fuzz := utils.Rand(0, 0.5)
					sphere_material := &hit.Metal{Albedo: albedo, Fuzz: fuzz}
					world.Add(hit.Sphere{Center: *center, Radius: 0.2, Material: sphere_material})
				} else {
					// glass
					sphere_material := &hit.Dielectric{RefractIndex: 1.5}
					world.Add(hit.Sphere{Center: *center, Radius: 0.2, Material: sphere_material})
				}
			}
		}
	}

	material1 := &hit.Dielectric{RefractIndex: 1.5}
	world.Add(hit.Sphere{Center: *vec.New(0.0, 1, 0), Radius: 1.0, Material: material1})

	material2 := &hit.Lambertian{Albedo: *vec.New(0.4, 0.2, 0.1)}
	world.Add(hit.Sphere{Center: *vec.New(-4, 1, 0), Radius: 1.0, Material: material2})

	material3 := &hit.Metal{Albedo: *vec.New(0.7, 0.6, 0.5), Fuzz: 0.0}
	world.Add(hit.Sphere{Center: *vec.New(4, 1, 0), Radius: 1.0, Material: material3})

	return world

	// cam := camera.New(
	// 	800,
	// 	16.0/9.0,
	// 	8,
	// 	20,
	// )

	// cam.Adjust(
	// 	20,
	// 	*vec.New(13, 2, 3),
	// 	*vec.New(0, 0, 0),
	// 	*vec.New(0, 1, 0),
	// )
}

func main() {
	world := getAwesomeWorld()

	cam := camera.New(
		600,
		16.0/9.0,
		10,
		20,
	)

	cam.Adjust(
		20,
		*vec.New(13, 2, 3),
		*vec.New(0, 0, 0),
		*vec.New(0, 1, 0),
	)

	cam.Render(world)
}
