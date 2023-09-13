package camera

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

type Camera struct {
	center            vec.Point
	pixel00_loc       vec.Point
	pixel_delta_u     vec.Vec3
	pixel_delta_v     vec.Vec3
	AspectRatio       float64
	Width             int
	Height            int
	samples_per_pixel int
}

var WHITE = vec.New(1, 1, 1)
var SKY_BLUE = vec.New(0.5, 0.7, 1)

func (c *Camera) ray_color(r *ray.Ray, h hit.HitList) vec.Color {
	hit_data := hit.HitData{}

	if h.Hit(*r, utils.NewInterval(0, math.Inf(1)), &hit_data) {
		r.Direction = vec.RandOnHemishpere(hit_data.Normal)
		r.Origin = hit_data.Point
		return *vec.DivScalar(c.ray_color(r, h), 2)
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

func (c *Camera) get_ray(center vec.Point) ray.Ray {
	sample := vec.Add(center, c.pixel_sample_sq())
	direction := vec.Sub(*sample, c.center)

	return *ray.New(c.center, *direction)
}

func (c *Camera) Render(world hit.HitList) {
	out := fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height)

	for j := 0; j < c.Height; j++ {
		fmt.Printf("\rRemaining: %f /", float64(j)/float64(c.Height)*100)
		for i := 0; i < c.Width; i++ {
			color := vec.New(0, 0, 0)
			pixel_center := *vec.Add(
				c.pixel00_loc,
				*vec.MulScalar(c.pixel_delta_u, float64(i)),
				*vec.MulScalar(c.pixel_delta_v, float64(j)),
			)
			for i := 0; i < c.samples_per_pixel; i++ {
				r := c.get_ray(pixel_center)

				color.Add(c.ray_color(&r, world))
			}
			color.DivScalar(float64(c.samples_per_pixel))
			out += color.ToClrStr()
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
