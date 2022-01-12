package merkletree

type MerkleTree struct {
	RootNode *MerkleNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		// If the amount of data is odd, then is needed to duplicate the last data, to complete the tree
		data = append(data, data[len(data)-1])
	}

	// starts adding adding the values to the nodes and pushing it to the array
	for _, value := range data {
		node := NewMerkleNode(nil, nil, value)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var levels []MerkleNode // sub-levels of nodes

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			levels = append(levels, *node)
		}

		nodes = levels
	}

	// creates the merkle-tree
	return &MerkleTree{&nodes[0]}
}
