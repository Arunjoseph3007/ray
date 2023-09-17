package hit

import (
	"ray-tracing/ray"
	"ray-tracing/vec"
)

type Metal struct {
	Albedo vec.Color
	Fuzz   float64
}

func (m *Metal) Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color) {
	scatter_dir := vec.Reflect(*vec.UnitVec(ray.Direction), rec.Normal)
	ray.Direction = *vec.Add(
		scatter_dir,
		*vec.MulScalar(vec.RandomUnitVec(), m.Fuzz),
	)
	ray.Origin = rec.Point
	return true, m.Albedo
}

func (m *Metal) Emitted(u, v float64, p vec.Point) vec.Color {
	return *vec.BLACK
}
