package hit

import (
	"math"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type Sphere struct {
	Center   vec.Point
	Radius   float64
	Material Material
}

func (s Sphere) Hit(r ray.Ray, limit utils.Interval, ret *HitData) bool {
	oc := vec.Sub(r.Origin, s.Center)
	a := r.Direction.LengthSquare()
	half_b := vec.Dot(r.Direction, *oc)
	c := oc.LengthSquare() - s.Radius*s.Radius
	dicriminant := half_b*half_b - a*c

	if dicriminant < 0 {
		return false
	}

	discSqrt := math.Sqrt(dicriminant)

	root := (-half_b - discSqrt) / a
	if root <= limit.Min || limit.Max <= root {
		root = (-half_b + discSqrt) / a
		if root <= limit.Min || limit.Max <= root {
			return false
		}
	}

	p := r.At(root)
	out_norm := vec.DivScalar(*vec.Sub(p, s.Center), s.Radius)

	ret.T = root
	ret.Point = p
	ret.Normal = *out_norm
	ret.Material = s.Material
	ret.SetFrontFace(r, *out_norm)

	return true
}
