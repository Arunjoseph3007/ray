package hit

import (
	"ray-tracing/ray"
	"ray-tracing/texture"
	"ray-tracing/vec"
)

type Light struct {
	Emit texture.Texture
}

func (l *Light) Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color) {
	return false, *vec.BLACK
}

func (l *Light) Emitted(u, v float64, p vec.Point) vec.Color {
	return l.Emit.Value(u, v, p)
}

func NewLightFromColor(c vec.Color) *Light {
	return &Light{texture.NewSolidTextureFromColor(c)}
}

func NewLightFromTexture(t texture.Texture) *Light {
	return &Light{t}
}
