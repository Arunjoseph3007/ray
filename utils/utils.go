package utils

import (
	"math"
	"math/rand"
)

func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func Rand(min float64, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
