package hit

import (
	"math"
	"ray-tracing/bvh"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type Quad struct {
	Q        vec.Point
	U, V     vec.Vec3
	Material Material
	BBox     bvh.AABB

	normal, w vec.Vec3
	d         float64
}

func (q Quad) Hit(r ray.Ray, limit utils.Interval, ret *HitData) bool {
	denom := vec.Dot(q.normal, r.Direction)
	if math.Abs(denom) < 1e-8 {
		return false
	}

	t := (q.d - vec.Dot(q.normal, r.Origin)) / denom

	if !limit.Contains(t) {
		return false
	}

	intersection := r.At(t)
	planar_vec := vec.Sub(intersection, q.Q)

	alpha := vec.Dot(q.w, *vec.Cross(*planar_vec, q.V))
	beta := vec.Dot(q.w, *vec.Cross(q.U, *planar_vec))

	if alpha < 0 || alpha > 1 || beta < 0 || beta > 1 {
		return false
	}

	ret.U = alpha
	ret.V = beta

	ret.T = t
	ret.Point = intersection
	ret.Material = q.Material
	ret.SetFrontFace(r, q.normal)

	return true
}

func (q *Quad) set_bounding_box() {
	q.BBox = bvh.FromPoints(q.Q, *vec.Add(q.Q, q.U, q.V)).Pad()
}

func (q Quad) BoundingBox() bvh.AABB {
	return q.BBox
}

func NewQuad(q, u, v vec.Vec3, mat Material) Quad {
	quad := Quad{
		Q:        q,
		U:        u,
		V:        v,
		Material: mat,
	}

	n := *vec.Cross(u, v)
	quad.normal = *vec.UnitVec(n)
	quad.d = vec.Dot(quad.normal, q)
	quad.w = *vec.DivScalar(n, n.LengthSquare())

	quad.set_bounding_box()

	return quad
}
