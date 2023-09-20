package hit

import (
	"math"
	"ray-tracing/bvh"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type RotateY struct {
	Object   Hitable
	sin, cos float64
	BBox     bvh.AABB
}

func NewRotateY(obj Hitable, ang float64) RotateY {
	sin := math.Sin(ang)
	cos := math.Cos(ang)
	r := RotateY{
		Object: obj,
		sin:    sin,
		cos:    cos,
	}

	r.BBox = obj.BoundingBox()

	min := *vec.New(math.Inf(1), math.Inf(1), math.Inf(1))
	max := *vec.New(math.Inf(-1), math.Inf(-1), math.Inf(-1))

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				x := float64(i)*r.BBox.X.Max + float64(1-i)*r.BBox.X.Min
				y := float64(j)*r.BBox.Y.Max + float64(1-j)*r.BBox.Y.Min
				z := float64(k)*r.BBox.Z.Max + float64(1-k)*r.BBox.Z.Min
				newx := r.cos*x + r.sin*z
				newz := -r.sin*x + r.cos*z

				tester := *vec.New(newx, y, newz)

				for c := 0; c < 3; c++ {
					min.Set(c, math.Min(min.Get(c), tester.Get(c)))
					max.Set(c, math.Max(max.Get(c), tester.Get(c)))
				}
			}
		}
	}

	r.BBox = bvh.FromPoints(min, max)

	return r
}

func (ry RotateY) Hit(r ray.Ray, limit utils.Interval, ret *HitData) bool {
	ox := ry.cos*r.Origin.Get(0) - ry.sin*r.Origin.Get(2)
	oz := ry.sin*r.Origin.Get(0) + ry.cos*r.Origin.Get(2)

	dx := ry.cos*r.Direction.Get(0) - ry.sin*r.Direction.Get(2)
	dz := ry.sin*r.Direction.Get(0) + ry.cos*r.Direction.Get(2)

	ray_rotated := *ray.New(
		*vec.New(ox, r.Origin.Y(), oz),
		*vec.New(dx, r.Direction.Y(), dz),
	)

	if !ry.Object.Hit(ray_rotated, limit, ret) {
		return false
	}

	p0 := ry.cos*ret.Point.X() + ry.sin*ret.Point.Z()
	p2 := -ry.sin*ret.Point.X() + ry.cos*ret.Point.Z()

	// Change the normal from object space to world space
	n0 := ry.cos*ret.Normal.X() + ry.sin*ret.Normal.Z()
	n2 := -ry.sin*ret.Normal.X() + ry.cos*ret.Normal.Z()

	ret.Point = *vec.New(p0, ret.Point.Y(), p2)
	ret.Normal = *vec.New(n0, ret.Normal.Y(), n2)

	return true
}

func (ry RotateY) BoundingBox() bvh.AABB {
	return ry.Object.BoundingBox()
}
