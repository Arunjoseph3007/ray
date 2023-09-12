package ray

import "ray-tracing/vec"

type Ray struct {
	Origin    vec.Point
	Direction vec.Vec3
}

func New(origin vec.Point, direction vec.Vec3) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r *Ray) At(t float64) vec.Point {
	return *vec.Add(r.Origin, *vec.MulScalar(r.Direction, t))
}
