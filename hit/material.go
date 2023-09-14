package hit

import (
	"ray-tracing/ray"
	"ray-tracing/vec"
)

type Material interface {
	Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color)
}

type Lambertian struct {
	Albedo vec.Color
}

func (l *Lambertian) Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color) {
	scatter_dir := vec.Add(rec.Normal, vec.RandomUnitVec())
	ray.Direction = *scatter_dir
	ray.Origin = rec.Point
	return true, l.Albedo
}

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
