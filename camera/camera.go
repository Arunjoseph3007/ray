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
	vfov              float64

	u, v, w vec.Vec3
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

func (c *Camera) pixel_sample_sq() vec.Vec3 {
	return *vec.Add(
		*vec.MulScalar(c.pixel_delta_u, utils.Rand(-0.5, .5)),
		*vec.MulScalar(c.pixel_delta_v, utils.Rand(-0.5, .5)),
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

				color.Add(c.ray_color(&r, world, c.max_depth))
			}
			color.DivScalar(float64(c.samples_per_pixel))
			fmt.Print(color.ToClrStr())
		}
	}
}

func New(width int, aspectRatio float64, samplePerPixel int, maxDepth int) Camera {
	c := Camera{}

	c.AspectRatio = aspectRatio
	c.Width = width
	c.Height = int(float64(c.Width) / c.AspectRatio)
	c.samples_per_pixel = samplePerPixel
	c.max_depth = maxDepth

	return c
}

func (c *Camera) Adjust(angle float64, from, at, viewup vec.Vec3) {
	c.vfov = angle

	c.center = from
	focal_length := vec.Sub(at, from).Length()
	rad := c.vfov * math.Pi / 180
	viewport_height := math.Tan(rad/2) * 2 * focal_length
	viewport_width := viewport_height * float64(c.Width) / float64(c.Height)

	c.w = *vec.UnitVec(*vec.Sub(from, at))
	c.u = *vec.UnitVec(*vec.Cross(c.w, viewup))
	c.v = *vec.Cross(c.u, c.w)

	viewport_u := vec.MulScalar(c.u, viewport_width)
	viewport_v := vec.MulScalar(*vec.Negative(c.v), viewport_height)

	c.pixel_delta_u = *vec.DivScalar(*viewport_u, float64(c.Width))
	c.pixel_delta_v = *vec.DivScalar(*viewport_v, float64(c.Height))

	viewport_up_left := vec.Sub(
		c.center,
		*vec.Add(
			*vec.MulScalar(c.w, focal_length),
			*vec.MulScalar(*viewport_u, .5),
			*vec.MulScalar(*viewport_v, .5),
		),
	)
	c.pixel00_loc = *vec.Add(
		*viewport_up_left,
		*vec.DivScalar(c.pixel_delta_u, 2),
		*vec.DivScalar(c.pixel_delta_v, 2),
	)
}
