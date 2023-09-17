package vec

import (
	"fmt"
	"math"
)

type Vec3 struct {
	e [3]float64
}

type Point = Vec3

type Color = Vec3

func New(a, b, c float64) *Vec3 {
	return &Vec3{[3]float64{a, b, c}}
}

func FromArray(p [3]float64) *Vec3 {
	return &Vec3{p}
}

func (v *Vec3) X() float64 { return v.e[0] }

func (v *Vec3) Y() float64 { return v.e[1] }

func (v *Vec3) Z() float64 { return v.e[2] }

func (v *Vec3) Get(n int) float64 {
	return v.e[n]
}

func (v Vec3) Negative() *Vec3 {
	return &Vec3{
		e: [3]float64{-v.e[0], -v.e[1], -v.e[2]},
	}
}

func (v *Vec3) Add(n Vec3) {
	v.e[0] += n.e[0]
	v.e[1] += n.e[1]
	v.e[2] += n.e[2]
}

func (v *Vec3) Sub(n Vec3) {
	v.e[0] -= n.e[0]
	v.e[1] -= n.e[1]
	v.e[2] -= n.e[2]
}

func (v *Vec3) MulScalar(t float64) {
	v.e[0] *= t
	v.e[1] *= t
	v.e[2] *= t
}

func (v *Vec3) DivScalar(t float64) {
	v.MulScalar(1 / t)
}

func (v *Vec3) LengthSquare() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquare())
}

func (v *Vec3) ToStr() string {
	return fmt.Sprintf("%f %f %f ", v.e[0], v.e[1], v.e[2])
}

func (v *Vec3) ToClrStr() string {
	x := int(math.Sqrt(v.e[0]) * 255.99)
	y := int(math.Sqrt(v.e[1]) * 255.99)
	z := int(math.Sqrt(v.e[2]) * 255.99)
	return fmt.Sprintf("%d %d %d\n", x, y, z)
}

func (v *Vec3) Print() {
	fmt.Print(v.ToStr())
}
