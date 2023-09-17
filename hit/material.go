package hit

import (
	"math"
	"math/rand"
	"ray-tracing/ray"
	"ray-tracing/texture"
	"ray-tracing/vec"
)

type Material interface {
	Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color)
}

type Lambertian struct {
	Albedo texture.Texture
}

func (l *Lambertian) Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color) {
	scatter_dir := vec.Add(rec.Normal, vec.RandomUnitVec())
	ray.Direction = *scatter_dir
	ray.Origin = rec.Point
	return true, l.Albedo.Value(rec.U, rec.V, rec.Point)
}

func NewLambertianFromColor(color vec.Color) *Lambertian {
	return &Lambertian{
		Albedo: texture.NewSolidTextureFromColor(color),
	}
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

type Dielectric struct {
	RefractIndex float64
}

func (d *Dielectric) Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color) {
	refraction_ratio := d.RefractIndex
	if rec.front_face {
		refraction_ratio = 1.0 / refraction_ratio
	}

	unit_direction := *vec.UnitVec(ray.Direction)

	cos_theta := math.Min(vec.Dot(rec.Normal, *vec.Negative(unit_direction)), 1.0)
	sin_theta := math.Sqrt(1 - cos_theta*cos_theta)

	cannot_refract := sin_theta*refraction_ratio > 1

	if cannot_refract || d.reflectance(cos_theta, refraction_ratio) > rand.Float64() {
		reflected := vec.Reflect(unit_direction, rec.Normal)
		ray.Direction = reflected
	} else {
		refracted := vec.Refract(unit_direction, rec.Normal, refraction_ratio)
		ray.Direction = refracted
	}

	ray.Origin = rec.Point
	return true, *vec.New(1, 1, 1)
}

func (d *Dielectric) reflectance(cosine, ref_idx float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - ref_idx) / (1 + ref_idx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
