package vec

func Negative(a Vec3) *Vec3 {
	return &Vec3{
		e: [3]float64{-a.e[0], -a.e[1], -a.e[2]},
	}
}

func Add(vectors ...Vec3) *Vec3 {
	x := 0.0
	y := 0.0
	z := 0.0

	for _, v := range vectors {
		x += v.e[0]
		y += v.e[1]
		z += v.e[2]
	}
	return &Vec3{
		e: [3]float64{x, y, z},
	}
}

func Sub(a, b Vec3) *Vec3 {
	return &Vec3{
		e: [3]float64{
			a.e[0] - b.e[0],
			a.e[1] - b.e[1],
			a.e[2] - b.e[2],
		},
	}
}

func MulScalar(a Vec3, t float64) *Vec3 {
	return &Vec3{
		e: [3]float64{
			a.e[0] * t,
			a.e[1] * t,
			a.e[2] * t,
		},
	}
}

func DivScalar(a Vec3, t float64) *Vec3 {
	return MulScalar(a, 1/t)
}

func Cross(a, b Vec3) *Vec3 {
	return &Vec3{
		e: [3]float64{
			a.e[1]*b.e[2] - a.e[2]*b.e[1],
			a.e[2]*b.e[0] - a.e[0]*b.e[2],
			a.e[0]*b.e[1] - a.e[1]*b.e[0],
		},
	}
}

func Dot(a, b Vec3) float64 {
	return a.e[0]*b.e[0] + a.e[1]*b.e[1] + a.e[2]*b.e[2]
}

func UnitVec(v Vec3) *Vec3 {
	return MulScalar(v, 1/v.Length())
}
