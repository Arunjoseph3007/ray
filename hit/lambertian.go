package hit

import (
	"ray-tracing/ray"
	"ray-tracing/texture"
	"ray-tracing/vec"
)

type Lambertian struct {
	Albedo texture.Texture
}

func (l *Lambertian) Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color) {
	scatter_dir := vec.Add(rec.Normal, vec.RandomUnitVec())
	ray.Direction = *scatter_dir
	ray.Origin = rec.Point
	return true, l.Albedo.Value(rec.U, rec.V, rec.Point)
}

func (l *Lambertian) Emitted(u, v float64, p vec.Point) vec.Color {
	return *vec.BLACK
}

func NewLambertianFromColor(color vec.Color) *Lambertian {
	return &Lambertian{
		Albedo: texture.NewSolidTextureFromColor(color),
	}
}

func NewLambertianFromtexture(tex texture.Texture) *Lambertian {
	return &Lambertian{
		Albedo: tex,
	}
}
