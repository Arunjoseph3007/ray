package texture

import "ray-tracing/vec"

type Texture interface {
	Value(u float64, v float64, p vec.Vec3) vec.Color
}
