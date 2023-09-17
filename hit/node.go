package hit

import (
	"math/rand"
	"ray-tracing/bvh"
	"ray-tracing/ray"
	"ray-tracing/utils"
	"sort"
)

type Node struct {
	Left  Hitable
	Right Hitable
	BBox  bvh.AABB
}

func (n Node) BoundingBox() bvh.AABB { return n.BBox }

func (n Node) Hit(r ray.Ray, limit utils.Interval, ret *HitData) bool {
	if !n.BBox.Hit(r, limit) {
		return false
	}

	hit_left := n.Left.Hit(r, limit, ret)

	itv := utils.NewInterval(limit.Min, limit.Max)

	if hit_left {
		itv.Max = ret.T
	}

	hit_right := n.Right.Hit(r, itv, ret)

	return hit_left || hit_right
}

func NodeFromHitables(objects []Hitable) Node {
	axis := rand.Intn(3)

	comp := func(i, j int) bool {
		return objects[i].BoundingBox().GetAxis(axis).Min < objects[j].BoundingBox().GetAxis(axis).Min
	}

	span := len(objects)

	node := Node{}

	if span == 1 {
		node.Left = objects[0]
		node.Right = objects[0]
	} else if span == 2 {
		if comp(0, 1) {
			node.Left = objects[0]
			node.Right = objects[1]
		} else {
			node.Right = objects[0]
			node.Left = objects[1]
		}
	} else {
		sort.SliceStable(objects, comp)
		mid := span / 2
		node.Left = NodeFromHitables(objects[:mid])
		node.Right = NodeFromHitables(objects[mid:])
	}

	node.BBox = bvh.GroupAABB(node.Left.BoundingBox(), node.Right.BoundingBox())

	return node
}
