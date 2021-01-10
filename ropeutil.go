package go_rope

type traverFunc func (r Rope)


func Concatenate(left Rope, right Rope) Rope {
	//if left rope is empty, return right rope
	if left.Len() == 0 {
		return right
	}

	//if right rope is empty, return left rope
	if right.Len() == 0 {
		return left
	}

	return concatenate(left, right)
}

func concatenate(left Rope, right Rope) Rope {

	//TODO: Need to detect overflow here
	//if the total length of left and right rope < MAX_LEAF_LENGTH, return a new RopeLeaf
	if left.Len() + right.Len() < MAX_LEAF_LENGTH {
		return NewRopeLeaf(left.String() + right.String())
	}

	cLeft, ok1 := left.(*RopeConcat)
	cRight, ok2 := right.(*RopeConcat)

	if !ok1 && ok2 {
		//If the right argument is a concatenation node whose left son is a short leaf, and the
		//left argument is also a short leaf, then we concatenate the two leaves, and then
		//concatenate the result to the right son of the right argument
		if  cRight.left.Len() + left.Len() < MAX_LEAF_LENGTH {
			return autoRebalance(NewRopeConcat( NewRopeLeaf(left.String() + cRight.left.String()), cRight.right ))
		}
	}
	
	
	if !ok2 && ok1 {
		//If the left argument is a concatenation node whose right son is a short leaf, and the
		//right argument is also a short leaf, then we concatenate the two leaves, and then
		//concatenate the result to the left son of the left argument
		if cLeft.right.Len()  + right.Len() < MAX_LEAF_LENGTH{
			return autoRebalance(NewRopeConcat( cLeft.left ,  NewRopeLeaf(cLeft.right.String() + right.String()) ))
		}
	}

	return autoRebalance(NewRopeConcat(left, right))
}

func autoRebalance(rope Rope) Rope {
	if rope.Depth() > MAX_ROPE_DEPTH {
		return rebalance(rope)
	}
	return rope
}


func build() Rope {
	return nil
}

func isBalanced(r Rope) bool {
	depth := r.Depth()
	if depth > (len(FIBONACCI) - 2) {
		return false
	}
	return FIBONACCI[depth + 2] <= r.Len()
}

func inOrderTraverse(r Rope, f traverFunc)  {
	switch r.(type) {
	case *RopeConcat:
		concat := r.(*RopeConcat)
		inOrderTraverse(concat.leftChild(), f)
		inOrderTraverse(concat.rightChild(), f)
	case *RopeLeaf:
		f(r)
	}
}

func rebalance(rope Rope) Rope {
	var leafNodes = []Rope{}
	inOrderTraverse(rope, func(r1 Rope) {
		leafNodes = append(leafNodes, r1)
	})

	return merge(leafNodes, 0, len(leafNodes))
}

func merge(leafNodes []Rope, start, end int) Rope {
	distance := end - start
	switch distance {
	case 1:
		return leafNodes[start]
	case 2:
		return NewRopeConcat(leafNodes[start], leafNodes[start + 1])
	default:
		mid := start + ( distance >> 1 )
		return NewRopeConcat(merge(leafNodes, start, mid), merge(leafNodes, mid, end))

	}
}