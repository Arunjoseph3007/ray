package hit

import "ray-tracing/ray"

type HitList struct {
	Objects []Hitable
}

func (h *HitList) Add(obj Hitable) {
	h.Objects = append(h.Objects, obj)
}

func (h *HitList) Hit(r ray.Ray, min float64, max float64, ret *HitData) bool {
	hit_some := false
	closest_so_far := max

	for _, obj := range h.Objects {
		if obj.Hit(r, min, closest_so_far, ret) {
			hit_some = true
			closest_so_far = ret.T
		}
	}

	return hit_some
}
