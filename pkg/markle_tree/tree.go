package markletree

type MarkleTree struct {
	RootNode *MarkleNode
}

func NewMarkleTree(data [][]byte) *MarkleTree {
	var nodes []MarkleNode

	if len(data)%2 != 0 {
		// If the amount of data is odd, then is needed to duplicate the last data, to complete the tree
		data = append(data, data[len(data)-1])
	}

	// starts adding adding the values to the nodes and pushing it to the array
	for _, value := range data {
		node := NewMarkleNode(nil, nil, value)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var levels []MarkleNode // sub-levels of nodes

		for j := 0; j < len(nodes); j += 2 {
			node := NewMarkleNode(&nodes[j], &nodes[j+1], nil)
			levels = append(levels, *node)
		}

		nodes = levels
	}

	// creates the markle-tree
	return &MarkleTree{&nodes[0]}
}
