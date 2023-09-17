package hit

import (
	"ray-tracing/bvh"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type HitList struct {
	Objects []Hitable
	BBox    bvh.AABB
}

func NewHitList() HitList {
	origin := vec.New(0, 0, 0)
	return HitList{
		Objects: []Hitable{},
		BBox:    bvh.FromPoints(*origin, *origin),
	}
}

func (h *HitList) Add(obj Hitable) {
	h.Objects = append(h.Objects, obj)
	h.BBox = bvh.GroupAABB(h.BBox, obj.BoundingBox())
}

func (h HitList) Hit(r ray.Ray, limit utils.Interval, ret *HitData) bool {
	hit_some := false
	closest_so_far := limit.Max

	for _, obj := range h.Objects {
		if obj.Hit(r, utils.NewInterval(limit.Min, closest_so_far), ret) {
			hit_some = true
			closest_so_far = ret.T
		}
	}

	return hit_some
}

func (h HitList) BoundingBox() bvh.AABB {
	return h.BBox
}
