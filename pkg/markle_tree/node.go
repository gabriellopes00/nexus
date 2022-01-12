package markletree

import "crypto/sha512"

type MarkleNode struct {
	Left  *MarkleNode
	Right *MarkleNode
	Data  []byte
}

func NewMarkleNode(left, right *MarkleNode, data []byte) *MarkleNode {
	node := MarkleNode{}

	if left == nil && right == nil {
		hash := sha512.Sum512(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha512.Sum512(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}
