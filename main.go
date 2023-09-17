package main

import (
	"math/rand"
	"os"
	"ray-tracing/camera"
	"ray-tracing/hit"
	"ray-tracing/texture"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

func getTestWorld() {
	world := hit.NewHitList()

	material_ground := hit.NewLambertianFromColor(*vec.New(0.8, 0.8, 0.0))
	material_center := hit.NewLambertianFromColor(*vec.New(0.1, 0.2, 0.5))
	material_left := &hit.Dielectric{RefractIndex: 1.5}
	material_right := &hit.Metal{Albedo: *vec.New(0.8, 0.6, 0.2), Fuzz: 0.0}

	world.Add(hit.NewSphere(*vec.New(0.0, -100.5, -1.0), 100.0, material_ground))
	world.Add(hit.NewSphere(*vec.New(0.0, 0.0, -1.0), 0.5, material_center))
	world.Add(hit.NewSphere(*vec.New(-1.0, 0.0, -1.0), 0.5, material_left))
	world.Add(hit.NewSphere(*vec.New(-1.0, 0.0, -1.0), -0.4, material_left))
	world.Add(hit.NewSphere(*vec.New(1.0, 0.0, -1.0), 0.5, material_right))
	cam := camera.New(
		400,
		16.0/9.0,
		20,
		20,
	)

	cam.Adjust(
		20,
		*vec.New(-2, 2, 1),
		*vec.New(0, 0, -1),
		*vec.New(0, 1, 0),
	)

	cam.Render(world)
}

func getAwesomeWorld() {
	world := hit.NewHitList()

	ground_material := hit.NewLambertianFromColor(*vec.New(0.5, 0.5, 0.5))
	world.Add(hit.NewSphere(*vec.New(0, -1000, 0), 1000, ground_material))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			choose_mat := rand.Float64()
			center := vec.New(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if vec.Sub(*center, *vec.New(4, 0.2, 0)).Length() > 0.9 {

				if choose_mat < 0.8 {
					// diffuse
					albedo := vec.Random(-1, 1)
					sphere_material := hit.NewLambertianFromColor(albedo)
					world.Add(hit.NewSphere(*center, 0.2, sphere_material))
				} else if choose_mat < 0.95 {
					// metal
					albedo := vec.Random(0.5, 1)
					fuzz := utils.Rand(0, 0.5)
					sphere_material := &hit.Metal{Albedo: albedo, Fuzz: fuzz}
					world.Add(hit.NewSphere(*center, 0.2, sphere_material))
				} else {
					// glass
					sphere_material := &hit.Dielectric{RefractIndex: 1.5}
					world.Add(hit.NewSphere(*center, 0.2, sphere_material))
				}
			}
		}
	}

	material1 := &hit.Dielectric{RefractIndex: 1.5}
	world.Add(hit.NewSphere(*vec.New(0.0, 1, 0), 1.0, material1))

	material2 := hit.NewLambertianFromColor(*vec.New(0.4, 0.2, 0.1))
	world.Add(hit.NewSphere(*vec.New(-4, 1, 0), 1.0, material2))

	material3 := &hit.Metal{Albedo: *vec.New(0.7, 0.6, 0.5), Fuzz: 0.0}
	world.Add(hit.NewSphere(*vec.New(4, 1, 0), 1.0, material3))

	worldNode := hit.NodeFromHitables(world.Objects)
	newWorld := hit.NewHitList()
	newWorld.Add(worldNode)

	cam := camera.New(
		600,
		16.0/9.0,
		8,
		20,
	)

	cam.Adjust(
		20,
		*vec.New(13, 2, 3),
		*vec.New(0, 0, 0),
		*vec.New(0, 1, 0),
	)
	println("Rendering ", len(world.Objects), "objects")
	cam.Render(newWorld)
}

func getTwoCheckerSphere() {
	world := hit.NewHitList()

	checker := texture.NewCheckerFromColors(0.8, *vec.New(.2, .3, .1), *vec.New(.9, .9, .9))

	world.Add(hit.NewSphere(*vec.New(0, -10, 0), 10, hit.NewLambertianFromtexture(checker)))
	world.Add(hit.NewSphere(*vec.New(0, 10, 0), 10, hit.NewLambertianFromtexture(checker)))

	cam := camera.New(
		400,
		16.0/9.0,
		20,
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

func main() {
	arg := os.Args[1]

	switch arg {
	case "awesome":
		getAwesomeWorld()
	case "test":
		getTestWorld()
	case "checker":
		getTwoCheckerSphere()
	default:
		println("Please provide scene")
	}
}
