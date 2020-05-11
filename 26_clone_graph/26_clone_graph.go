package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	cloned := cloneGraphRecursive(node, make(map[int]*Node, 0))
	return cloned
}

func cloneGraphRecursive(root *Node, m map[int]*Node) *Node {
	if root == nil {
		return nil
	}

	if val, ok := m[root.Val]; ok {
		return val
	}

	newNode := &Node{}
	newNode.Val = root.Val

	m[root.Val] = newNode

	for _, child := range root.Neighbors {
		node := cloneGraphRecursive(child, m)
		newNode.Neighbors = append(newNode.Neighbors, node)
	}

	return newNode
}
