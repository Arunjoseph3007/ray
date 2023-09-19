package camera

import (
	"fmt"
	"math"
	"ray-tracing/hit"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
	"image"
    "image/png"
	"os"
)

type Camera struct {
	Width       int
	Height      int
	AspectRatio float64

	samples_per_pixel int
	max_depth         int

	center  vec.Point
	vfov    float64
	u, v, w vec.Vec3

	pixel00_loc   vec.Point
	pixel_delta_u vec.Vec3
	pixel_delta_v vec.Vec3

	DefocusAngle   float64
	FocusDist      float64
	defocus_disk_u vec.Vec3
	defocus_disk_v vec.Vec3

	Background vec.Color
}

func (c *Camera) ray_color(r *ray.Ray, world hit.HitList, depth int) vec.Color {
	hit_data := hit.HitData{}

	if depth == 0 {
		return *vec.BLACK
	}

	if !world.Hit(*r, utils.NewInterval(0.01, math.Inf(1)), &hit_data) {
		return c.Background
	}

	emitted_color := hit_data.Material.Emitted(hit_data.U, hit_data.V, hit_data.Point)

	did_scatter, atten := hit_data.Material.Scatter(r, &hit_data)
	if !did_scatter {
		return emitted_color
	}
	scatter_color := *vec.Mul(atten, c.ray_color(r, world, depth-1))

	return *vec.Add(emitted_color, scatter_color)
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

	upLeft := image.Point{0, 0}
	lowRight := image.Point{c.Width, c.Height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for j := 0; j < c.Height; j++ {
		print("\rRemaining ", j*100/c.Height, "%")
		for i := 0; i < c.Width; i++ {
			color := vec.New(0, 0, 0)
			pixel_center := *vec.Add(
				c.pixel00_loc,
				*vec.MulScalar(c.pixel_delta_u, float64(i)),
				*vec.MulScalar(c.pixel_delta_v, float64(j)),
			)
			for i := 0; i < c.samples_per_pixel; i++ {
				sample := vec.Add(pixel_center, c.pixel_sample_sq())
				ray_origin := vec.Add(
					c.center,
					*vec.MulScalar(c.defocus_disk_u, utils.Rand(-.5, .5)),
					*vec.MulScalar(c.defocus_disk_v, utils.Rand(-.5, .5)),
				)
				direction := vec.Sub(*sample, *ray_origin)
				r := *ray.New(*ray_origin, *direction)

				color.Add(c.ray_color(&r, world, c.max_depth))
			}

			color.DivScalar(float64(c.samples_per_pixel))
			img.Set(i,j,color.ToIntArr())
			// fmt.Print(color.ToClrStr())
		}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
	defer f.Close()
	print("\rRemaining 100%")
	println("\nDone")
}

func New(width int, aspectRatio float64, samplePerPixel int, maxDepth int) Camera {
	c := Camera{}

	c.AspectRatio = aspectRatio
	c.Width = width
	c.Height = int(float64(c.Width) / c.AspectRatio)
	c.samples_per_pixel = samplePerPixel
	c.max_depth = maxDepth

	c.DefocusAngle = 0
	c.FocusDist = 10

	c.Background = *vec.SKY_BLUE

	return c
}

func (c *Camera) Adjust(angle float64, from, at, viewup vec.Vec3) {
	c.vfov = angle

	c.center = from
	rad := c.vfov * math.Pi / 180
	viewport_height := math.Tan(rad/2) * 2 * c.FocusDist
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
			*vec.MulScalar(c.w, c.FocusDist),
			*vec.MulScalar(*viewport_u, .5),
			*vec.MulScalar(*viewport_v, .5),
		),
	)
	c.pixel00_loc = *vec.Add(
		*viewport_up_left,
		*vec.DivScalar(c.pixel_delta_u, 2),
		*vec.DivScalar(c.pixel_delta_v, 2),
	)

	// Calculate the camera defocus disk basis vectors.
	defocus_rad := c.DefocusAngle * math.Pi / 180
	defocus_radius := c.FocusDist * math.Tan(defocus_rad/2)
	c.defocus_disk_u = *vec.MulScalar(c.u, defocus_radius)
	c.defocus_disk_v = *vec.MulScalar(c.v, defocus_radius)
}
