package hit

import (
	"math"
	"ray-tracing/ray"
	"ray-tracing/vec"
)

type Shpere struct {
	Center vec.Point
	Radius float64
}

func (s Shpere) Hit(r ray.Ray, min float64, max float64, ret *HitData) bool {
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
	if root <= min || max <= root {
		root = (-half_b + discSqrt) / a
		if root <= min || max <= root {
			return false
		}
	}

	p := r.At(root)
	out_norm := vec.DivScalar(*vec.Sub(p, s.Center), s.Radius)

	ret.T = root
	ret.Point = p
	ret.Normal = *out_norm
	ret.SetFrontFace(r, *out_norm)

	return true
}
