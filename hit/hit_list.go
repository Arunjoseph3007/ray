package hit

import (
	"ray-tracing/ray"
	"ray-tracing/utils"
)

type HitList struct {
	Objects []Hitable
}

func (h *HitList) Add(obj Hitable) {
	h.Objects = append(h.Objects, obj)
}

func (h *HitList) Hit(r ray.Ray, limit utils.Interval, ret *HitData) bool {
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
