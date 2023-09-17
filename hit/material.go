package hit

import (
	"ray-tracing/ray"
	"ray-tracing/vec"
)

type Material interface {
	Scatter(ray *ray.Ray, rec *HitData) (bool, vec.Color)
}
