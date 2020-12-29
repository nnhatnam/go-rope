package go_rope

type RopeLeaf struct {
	_sequence string
}

func NewRopeLeaf(s string) *RopeLeaf {
	return &RopeLeaf{_sequence: s}
}

func (leaf *RopeLeaf) Depth() int {
	return 0
}

func (leaf *RopeLeaf) Index(idx int) byte {
	return leaf._sequence[idx]
}

func (leaf *RopeLeaf) Len() int {
	return len(leaf._sequence)
}

func (leaf *RopeLeaf) String() string {
	return leaf._sequence
}

func (leaf *RopeLeaf) Concat(r Rope) Rope {
	if leaf2, isLeaf := r.(RopeLeaf); isLeaf && (leaf.Len() + leaf2.Len()) < MAX_LEAF_LENGTH {
		return NewRopeLeaf(leaf._sequence + leaf2._sequence )
	}
	//TO DO: Handle concat rope case
	return nil
}

