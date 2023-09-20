package main

import (
	"math"
	"math/rand"
	"os"
	"ray-tracing/camera"
	"ray-tracing/hit"
	"ray-tracing/texture"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

func TestScene() {
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

	cam.Render(world,"test")
}

func AwesomeScene() {
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
		50,
		50,
	)

	cam.Adjust(
		20,
		*vec.New(13, 2, 3),
		*vec.New(0, 0, 0),
		*vec.New(0, 1, 0),
	)
	println("Rendering ", len(world.Objects), "objects")
	cam.Render(newWorld, "awesome")
}

func CheckerSphereScene() {
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

	cam.Render(world, "checker")
}

func EarthScene() {
	world := hit.NewHitList()

	earth_texture := texture.NewImageTexture("assets/earthmap.png")
	earth_surface := hit.NewLambertianFromtexture(earth_texture)

	world.Add(hit.NewSphere(*vec.New(0, 0, 0), 2, earth_surface))

	cam := camera.New(
		400,
		16.0/9.0,
		20,
		20,
	)

	cam.Adjust(
		20,
		*vec.New(0, 0, 12),
		*vec.New(0, 0, 0),
		*vec.New(0, 1, 0),
	)

	cam.Render(world,"earth")
}

func SimpleLightScene() {
	world := hit.NewHitList()

	material_ground := hit.NewLambertianFromColor(*vec.New(0.05, 0.05, 0.05))
	material_center := hit.NewLambertianFromColor(*vec.New(0.1, 0.2, 0.5))
	material_left := &hit.Dielectric{RefractIndex: 1.5}
	material_right := &hit.Metal{Albedo: *vec.New(0.8, 0.6, 0.2), Fuzz: 0.0}
	light_material := hit.NewLightFromColor(*vec.WHITE)

	world.Add(hit.NewSphere(*vec.New(0.0, -100.5, -1.0), 100.0, material_ground))
	world.Add(hit.NewSphere(*vec.New(0.0, 0.0, -1.0), 0.5, material_right))
	world.Add(hit.NewSphere(*vec.New(-1.0, 0.0, -1.0), 0.5, material_left))
	world.Add(hit.NewSphere(*vec.New(-1.0, 0.0, -1.0), -0.4, material_left))
	world.Add(hit.NewSphere(*vec.New(1.0, 0.0, -1.0), 0.5, light_material))
	world.Add(hit.NewSphere(*vec.New(2.0, 0.0, -1.0), 0.5, material_center))
	cam := camera.New(
		400,
		16.0/9.0,
		50,
		100,
	)

	cam.Background = *vec.New(100/255,100/255,100/255)

	cam.Adjust(
		30,
		*vec.New(-2, 1, 1),
		*vec.New(0, 0, -1),
		*vec.New(0, 1, 0),
	)

	cam.Render(world, "light")
}

func CornelScene() {
	world := hit.NewHitList()
	red := hit.NewLambertianFromColor(*vec.New(.65, .05, .05))
	white := hit.NewLambertianFromColor(*vec.New(.73, .73, .73))
	green := hit.NewLambertianFromColor(*vec.New(.12, .45, .15))
	light := hit.NewLightFromColor(*vec.New(15, 15, 15))

	world.Add(hit.NewQuad(*vec.New(555, 0, 0), *vec.New(0, 555, 0), *vec.New(0, 0, 555), green))
	world.Add(hit.NewQuad(*vec.New(0, 0, 0), *vec.New(0, 555, 0), *vec.New(0, 0, 555), red))
	world.Add(hit.NewQuad(*vec.New(343, 554, 332), *vec.New(-130, 0, 0), *vec.New(0, 0, -105), light))
	world.Add(hit.NewQuad(*vec.New(0, 0, 0), *vec.New(555, 0, 0), *vec.New(0, 0, 555), white))
	world.Add(hit.NewQuad(*vec.New(555, 555, 555), *vec.New(-555, 0, 0), *vec.New(0, 0, -555), white))
	world.Add(hit.NewQuad(*vec.New(0, 0, 555), *vec.New(555, 0, 0), *vec.New(0, 555, 0), white))

	world.Add(getBox(*vec.New(130, 0, 65), *vec.New(295, 165, 230), white))
	world.Add(getBox(*vec.New(265, 0, 295), *vec.New(430, 330, 460), white))

	cam := camera.New(
		300,
		1,
		100,
		20,
	)
	cam.Background = *vec.BLACK
	cam.Adjust(
		40,
		*vec.New(278, 278, -800),
		*vec.New(278, 278, 0),
		*vec.New(0, 1, 0),
	)

	cam.Render(world, "cornel")
}

func getBox(a, b vec.Point, mat hit.Material) hit.HitList {
	sides := hit.NewHitList()

	min := *vec.New(math.Min(a.X(), b.X()), math.Min(a.Y(), b.Y()), math.Min(a.Z(), b.Z()))
	max := *vec.New(math.Max(a.X(), b.X()), math.Max(a.Y(), b.Y()), math.Max(a.Z(), b.Z()))

	dx := *vec.New(max.X()-min.X(), 0, 0)
	dy := *vec.New(0, max.Y()-min.Y(), 0)
	dz := *vec.New(0, 0, max.Z()-min.Z())

	sides.Add(hit.NewQuad(*vec.New(min.X(), min.Y(), max.Z()), dx, dy, mat))
	sides.Add(hit.NewQuad(*vec.New(max.X(), min.Y(), max.Z()), *vec.Negative(dz), dy, mat))
	sides.Add(hit.NewQuad(*vec.New(max.X(), min.Y(), min.Z()), *vec.Negative(dx), dy, mat))
	sides.Add(hit.NewQuad(*vec.New(min.X(), min.Y(), min.Z()), dz, dy, mat))
	sides.Add(hit.NewQuad(*vec.New(min.X(), max.Y(), max.Z()), dx, *vec.Negative(dz), mat))
	sides.Add(hit.NewQuad(*vec.New(min.X(), min.Y(), min.Z()), dx, dz, mat))

	return sides
}

func QuadsScene() {
	world := hit.NewHitList()
	left_red := hit.NewLambertianFromColor(*vec.New(1.0, 0.2, 0.2))
	back_green := hit.NewLambertianFromColor(*vec.New(0.2, 1.0, 0.2))
	right_blue := hit.NewLambertianFromColor(*vec.New(0.2, 0.2, 1.0))
	upper_orange := hit.NewLambertianFromColor(*vec.New(1.0, 0.5, 0.0))
	lower_teal := hit.NewLambertianFromColor(*vec.New(0.2, 0.8, 0.8))

	light := hit.NewLightFromColor(*vec.WHITE)

	world.Add(hit.NewQuad(*vec.New(-3, -2, 5), *vec.New(0, 0, -4), *vec.New(0, 4, 0), left_red))
	world.Add(hit.NewQuad(*vec.New(-2, -2, 0), *vec.New(4, 0, 0), *vec.New(0, 4, 0), back_green))
	world.Add(hit.NewQuad(*vec.New(3, -2, 1), *vec.New(0, 0, 4), *vec.New(0, 4, 0), right_blue))
	world.Add(hit.NewQuad(*vec.New(-2, 3, 1), *vec.New(4, 0, 0), *vec.New(0, 0, 4), upper_orange))
	world.Add(hit.NewQuad(*vec.New(-2, -3, 5), *vec.New(4, 0, 0), *vec.New(0, 0, -4), lower_teal))

	world.Add(hit.NewSphere(*vec.New(-2.5, 0, 0), .5, light))

	cam := camera.New(
		300,
		1,
		100,
		50,
	)
	cam.Background = *vec.BLACK
	cam.Adjust(
		80,
		*vec.New(0, 0, 9),
		*vec.New(0, 0, 0),
		*vec.New(0, 1, 0),
	)

	cam.Render(world, "quad")
}

func main() {
	arg := os.Args[1]

	switch arg {
	case "awesome":
		AwesomeScene()
	case "test":
		TestScene()
	case "checker":
		CheckerSphereScene()
	case "earth":
		EarthScene()
	case "light":
		SimpleLightScene()
	case "quad":
		QuadsScene()
	case "cornel":
		CornelScene()
	default:
		println("Please provide scene")
	}
}
