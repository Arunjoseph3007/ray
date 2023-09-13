package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"ray-tracing/hit"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

var WHITE = vec.New(1, 1, 1)
var BLUE = vec.New(0.5, 0.7, 1)

func ray_color(r ray.Ray, h hit.HitList) vec.Color {
	hit_data := hit.HitData{}

	if h.Hit(r, utils.NewInterval(0, math.Inf(1)), &hit_data) {
		return *vec.MulScalar(
			*vec.Add(
				hit_data.Normal,
				*vec.New(1, 1, 1),
			),
			0.5,
		)
	}

	dir := vec.UnitVec(r.Direction)
	a := 0.5 * (dir.Y() + 1)

	return *vec.Add(
		*vec.MulScalar(*WHITE, 1-a),
		*vec.MulScalar(*BLUE, a),
	)
}

func main() {
	ratio := 16.0 / 9.0
	width := 200
	height := int(float64(width) / ratio)

	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * float64(width) / float64(height)
	cam_center := vec.New(0, 0, 0)

	viewport_u := vec.New(viewport_width, 0, 0)
	viewport_v := vec.New(0, -viewport_height, 0)

	pixel_delta_u := vec.DivScalar(*viewport_u, float64(width))
	pixel_delta_v := vec.DivScalar(*viewport_v, float64(height))

	viewport_up_left := vec.Sub(
		*cam_center,
		*vec.Add(
			*vec.New(0, 0, focal_length),
			*vec.DivScalar(*viewport_u, 2),
			*vec.DivScalar(*viewport_v, 2),
		),
	)
	pixel_00 := vec.Add(
		*viewport_up_left,
		*vec.DivScalar(*pixel_delta_u, 2),
		*vec.DivScalar(*pixel_delta_v, 2),
	)

	world := hit.HitList{}
	world.Add(hit.Shpere{
		Center: *vec.New(0, 0, -1),
		Radius: .5,
	})

	world.Add(hit.Shpere{
		Center: *vec.New(0, -100.5, -1),
		Radius: 100,
	})

	out := fmt.Sprintf("P3\n%d %d\n255\n", width, height)

	for j := 0; j < height; j++ {
		fmt.Printf("\rRemaining: %f /", float64(j)/float64(height)*100)
		for i := 0; i < width; i++ {
			pixel_center := vec.Add(
				*pixel_00,
				*vec.MulScalar(*pixel_delta_u, float64(i)),
				*vec.MulScalar(*pixel_delta_v, float64(j)),
			)
			ray_direction := vec.Sub(
				*pixel_center,
				*cam_center,
			)
			r := ray.New(*cam_center, *ray_direction)

			pixel_color := ray_color(*r, world)
			out += pixel_color.ToClrStr()
		}
	}

	file, err := os.Create("image.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err2 := file.WriteString(out)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Done!")
}
