package hit

import (
	"ray-tracing/ray"
	"ray-tracing/vec"
)

type HitData struct {
	Point      vec.Point
	Normal     vec.Vec3
	T          float64
	front_face bool
}

func (h *HitData) SetFrontFace(r ray.Ray, out_normal vec.Vec3) {
	h.front_face = vec.Dot(h.Normal, out_normal) < 0
}

type Hitable interface {
	Hit(r ray.Ray, min float64, max float64, ret *HitData) bool
}
