package camera

import (
	"fmt"
	"math"
	"ray-tracing/hit"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type Camera struct {
	center            vec.Point
	pixel00_loc       vec.Point
	pixel_delta_u     vec.Vec3
	pixel_delta_v     vec.Vec3
	AspectRatio       float64
	Width             int
	Height            int
	samples_per_pixel int
	max_depth         int
}

var WHITE = vec.New(1, 1, 1)
var BLACK = vec.New(0, 0, 0)
var SKY_BLUE = vec.New(0.5, 0.7, 1)

func (c *Camera) ray_color(r *ray.Ray, world hit.HitList, depth int) vec.Color {
	hit_data := hit.HitData{}

	if depth == 0 {
		return *BLACK
	}

	if world.Hit(*r, utils.NewInterval(0.01, math.Inf(1)), &hit_data) {
		did_scatter, atten := hit_data.Material.Scatter(r, &hit_data)
		if did_scatter {
			return *vec.Mul(atten, c.ray_color(r, world, depth-1))
		}
		return *BLACK
	}

	dir := vec.UnitVec(r.Direction)
	a := 0.5 * (dir.Y() + 1)

	return *vec.Add(
		*vec.MulScalar(*WHITE, 1-a),
		*vec.MulScalar(*SKY_BLUE, a),
	)
}

func (c *Camera) Initialize() {
	c.AspectRatio = 16.0 / 9.0
	c.Width = 400
	c.Height = int(float64(c.Width) / c.AspectRatio)
	c.samples_per_pixel = 4
	c.max_depth = 20

	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * float64(c.Width) / float64(c.Height)
	c.center = *vec.New(0, 0, 0)

	viewport_u := vec.New(viewport_width, 0, 0)
	viewport_v := vec.New(0, -viewport_height, 0)

	c.pixel_delta_u = *vec.DivScalar(*viewport_u, float64(c.Width))
	c.pixel_delta_v = *vec.DivScalar(*viewport_v, float64(c.Height))

	viewport_up_left := vec.Sub(
		c.center,
		*vec.Add(
			*vec.New(0, 0, focal_length),
			*vec.DivScalar(*viewport_u, 2),
			*vec.DivScalar(*viewport_v, 2),
		),
	)
	c.pixel00_loc = *vec.Add(
		*viewport_up_left,
		*vec.DivScalar(c.pixel_delta_u, 2),
		*vec.DivScalar(c.pixel_delta_v, 2),
	)
}

func (c *Camera) pixel_sample_sq() vec.Vec3 {
	return *vec.Add(
		*vec.MulScalar(c.pixel_delta_u, utils.Rand(-0.3, .3)),
		*vec.MulScalar(c.pixel_delta_v, utils.Rand(-0.3, .3)),
	)
}

func (c *Camera) Render(world hit.HitList) {
	fmt.Println("P3")
	fmt.Println(c.Width, c.Height)
	fmt.Println("255")

	for j := 0; j < c.Height; j++ {
		for i := 0; i < c.Width; i++ {
			color := vec.New(0, 0, 0)
			pixel_center := *vec.Add(
				c.pixel00_loc,
				*vec.MulScalar(c.pixel_delta_u, float64(i)),
				*vec.MulScalar(c.pixel_delta_v, float64(j)),
			)
			for i := 0; i < c.samples_per_pixel; i++ {
				sample := vec.Add(pixel_center, c.pixel_sample_sq())
				direction := vec.Sub(*sample, c.center)
				r := *ray.New(c.center, *direction)
				// r := c.get_ray(pixel_center)

				color.Add(c.ray_color(&r, world, c.max_depth))
			}
			color.DivScalar(float64(c.samples_per_pixel))
			fmt.Print(color.ToClrStr())
		}
	}
}
