package texture

import "ray-tracing/vec"

type SolidTexture struct {
	Color vec.Color
}

func NewSolidTextureFromColor(color vec.Color) SolidTexture {
	return SolidTexture{
		Color: color,
	}
}

func NewSolidTextureFromRGB(r, g, b float64) SolidTexture {
	return SolidTexture{
		Color: *vec.New(r, g, b),
	}
}

func (st SolidTexture) Value(u float64, v float64, p vec.Vec3) vec.Color {
	return st.Color
}
