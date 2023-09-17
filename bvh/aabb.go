package bvh

import (
	"math"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type AABB struct {
	X, Y, Z utils.Interval
}

func New(x, y, z utils.Interval) AABB {
	return AABB{
		X: x,
		Y: y,
		Z: z,
	}
}

func FromPoints(a, b vec.Point) AABB {
	return AABB{
		X: utils.Interval{Min: math.Min(a.X(), b.X()), Max: math.Max(a.X(), b.X())},
		Y: utils.Interval{Min: math.Min(a.Y(), b.Y()), Max: math.Max(a.Y(), b.Y())},
		Z: utils.Interval{Min: math.Min(a.Z(), b.Z()), Max: math.Max(a.Z(), b.Z())},
	}
}

func GroupAABB(a, b AABB) AABB {
	return AABB{
		X: utils.GroupInterval(a.X, b.X),
		Y: utils.GroupInterval(a.Y, b.Y),
		Z: utils.GroupInterval(a.Z, b.Z),
	}
}

func (a AABB) GetAxis(n int) utils.Interval {
	if n == 0 {
		return a.X
	}
	if n == 1 {
		return a.Y
	}
	return a.Z
}

func (ab *AABB) Hit(r ray.Ray, t utils.Interval) bool {
	for a := 0; a < 3; a++ {
		invD := 1 / r.Direction.Get(a)
		orig := r.Origin.Get(a)

		t0 := (ab.GetAxis(a).Min - orig) * invD
		t1 := (ab.GetAxis(a).Max - orig) * invD

		if invD < 0 {
			temp := t0
			t0 = t1
			t1 = temp
		}

		if t0 > t.Min {
			t.Min = t0
		}
		if t1 < t.Max {
			t.Max = t1
		}

		if t.Max <= t.Min {
			return false
		}
	}
	return true
}
