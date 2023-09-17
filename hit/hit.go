package hit

import (
	"ray-tracing/bvh"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type HitData struct {
	Point      vec.Point
	Normal     vec.Vec3
	T          float64
	Material   Material
	front_face bool
}

func (h *HitData) SetFrontFace(r ray.Ray, out_normal vec.Vec3) {
	h.front_face = vec.Dot(r.Direction, out_normal) < 0
	if h.front_face {
		h.Normal = out_normal
	} else {
		h.Normal = *vec.Negative(out_normal)
	}
}

type Hitable interface {
	Hit(r ray.Ray, limit utils.Interval, ret *HitData) bool

	BoundingBox() bvh.AABB
}
