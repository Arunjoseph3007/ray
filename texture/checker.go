package texture

import (
	"math"
	"ray-tracing/vec"
)

type CheckerTexture struct {
	inv_scale float64
	odd, even Texture
}

func (ct CheckerTexture) Value(u float64, v float64, p vec.Vec3) vec.Color {
	sum := math.Floor(ct.inv_scale*p.X()) +
		math.Floor(ct.inv_scale*p.Y()) +
		math.Floor(ct.inv_scale*p.Z())

	isEven := int(sum)%2 == 0

	if isEven {
		return ct.even.Value(u, v, p)
	} else {
		return ct.odd.Value(u, v, p)
	}

}

func NewCheckerFromTextures(scale float64, t1, t2 Texture) CheckerTexture {
	return CheckerTexture{
		inv_scale: 1 / scale,
		odd:       t1,
		even:      t2,
	}
}

func NewCheckerFromColors(scale float64, c1, c2 vec.Color) CheckerTexture {
	return CheckerTexture{
		inv_scale: 1 / scale,
		odd:       NewSolidTextureFromColor(c1),
		even:      NewSolidTextureFromColor(c2),
	}
}
