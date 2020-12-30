package go_rope

import "math"

type RopeConcat struct {
	depth byte
	len int
	left Rope
	right Rope
}

func NewRopeConcat(left, right Rope) *RopeConcat {
	return &RopeConcat{
		depth: byte(math.Max(float64(left.Depth()), float64(right.Depth()))) + 1,
		len:    left.Len() + right.Len(),
		left:   left,
		right:  right,
	}
}

func (rc *RopeConcat) rightChild() Rope {
	return rc.right
}

func (rc *RopeConcat) leftChild() Rope {
	return rc.left
}

func (rc *RopeConcat) Depth() byte {
	return rc.depth
}

func (rc *RopeConcat) Len() int {
	return rc.len
}

func (rc *RopeConcat) String() string {
	return rc.left.String() + rc.right.String()
}

func (rc *RopeConcat) Index(idx int) byte {
	if idx < rc.left.Len() { return rc.left.Index(idx) }
	return rc.right.Index(idx - rc.left.Len())
}

func (rc *RopeConcat) Concat(r Rope) Rope {
	rope1, ok1 := rc.right.(*RopeLeaf)
	rope2, ok2 := r.(*RopeLeaf)
	if ok1 && ok2 && rope1.Len() + rope2.Len() <= MAX_LEAF_LENGTH {
		right := NewRopeLeaf(rope1.String() + rope2.String())
		left := rc.left
		return NewRopeConcat(left, right)
	}
	return NewRopeConcat(rc, r)
}