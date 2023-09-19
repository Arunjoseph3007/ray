package utils

import "math"

type Interval struct {
	Min float64
	Max float64
}

var UnitInterval Interval = NewInterval(0, 1)

func GroupInterval(a, b Interval) Interval {
	return Interval{
		Min: math.Min(a.Min, b.Min),
		Max: math.Max(a.Max, b.Max),
	}
}

func (i *Interval) Contains(x float64) bool {
	return i.Min <= x && i.Max >= x
}

func (i *Interval) Surrounds(x float64) bool {
	return i.Min < x && i.Max > x
}

func (i *Interval) Clamp(x float64) float64 {
	if x < i.Min {
		return i.Min
	}
	if x > i.Max {
		return i.Max
	}
	return x
}

func (i *Interval) Size() float64 {
	return i.Max - i.Min
}

func (i *Interval) Expand(delta float64) Interval {
	padding := delta / 2
	return Interval{
		Min: i.Min - padding,
		Max: i.Max + padding,
	}
}

func NewInterval(min, max float64) Interval {
	return Interval{
		Min: min,
		Max: max,
	}
}

var Empty = Interval{Min: math.Inf(1), Max: math.Inf(-1)}
var Universe = Interval{Min: math.Inf(-1), Max: math.Inf(1)}
