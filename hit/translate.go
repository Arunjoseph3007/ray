package hit

import (
	"ray-tracing/bvh"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type Translate struct {
	Object Hitable
	Offset vec.Vec3
	BBox   bvh.AABB
}

func NewTranslate(obj Hitable, off vec.Vec3) Translate {
	t := Translate{
		Object: obj,
		Offset: off,
	}

	t.BBox = obj.BoundingBox().Translate(off)

	return t
}

func (t Translate) Hit(r ray.Ray, limit utils.Interval, ret *HitData) bool {
	offset_ray := ray.New(
		*vec.Sub(r.Origin, t.Offset),
		r.Direction,
	)

	if !t.Object.Hit(*offset_ray, limit, ret) {
		return false
	}

	ret.Point.Add(t.Offset)

	return true
}

func (t Translate) BoundingBox() bvh.AABB {
	return t.Object.BoundingBox()
}
