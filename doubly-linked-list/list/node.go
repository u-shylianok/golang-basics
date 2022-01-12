package list

type Node struct {
	next  *Node
	prev  *Node
	value int
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Value() int {
	return n.value
}
