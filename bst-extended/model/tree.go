package model

import "fmt"

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value Keykeeper) {
	if t.Root == nil {
		t.Root = &Node{Key: value.GetKey(), Value: value}
		return
	}
	t.Root.Insert(value)
}

func (t *Tree) Remove(key string) error {
	if t.Root != nil {
		return remove(&t.Root, key)
	}
	return fmt.Errorf("the tree is empty")
}

func (t *Tree) Find(key string) (Keykeeper, error) {
	if t.Root != nil {
		return t.Root.find(key)
	}
	return nil, fmt.Errorf("the tree is empty")
}

type Node struct {
	Key   string
	Value Keykeeper
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value Keykeeper) {
	if value.GetKey() < n.Key {
		// insert into the left tree
		if n.Left == nil {
			n.Left = &Node{Key: value.GetKey(), Value: value}
		} else {
			n.Left.Insert(value)
		}
		return
	}
	// insert into the right tree
	if n.Right == nil {
		n.Right = &Node{Key: value.GetKey(), Value: value}
	} else {
		n.Right.Insert(value)
	}
}

func remove(node **Node, key string) error {
	n := *node
	if n == nil {
		return fmt.Errorf("key value does not exist")
	}

	if key < n.Key {
		return remove(&n.Left, key)
	}
	if key > n.Key {
		return remove(&n.Right, key)
	}

	// case 1 - node to remove has 0 childs
	if n.Left == nil && n.Right == nil {
		*node = nil
		return nil
	}
	// case 2 - node to remove has 2 childs
	if n.Left != nil && n.Right != nil {
		minRightParent := n
		minRight := n.Right

		for minRight.Left != nil {
			minRightParent = minRight
			minRight = minRight.Left
		}

		tmp := *minRight
		if minRightParent != n {
			minRightParent.Left = minRight.Right
		} else {
			minRightParent.Right = minRight.Right
		}

		tmp.Left = n.Left
		tmp.Right = n.Right

		*node = &tmp
		return nil
	}
	// case 3 - node to remove has 1 child
	if n.Left != nil {
		*node = (*node).Left
	} else {
		*node = (*node).Right
	}
	return nil
}

func (n *Node) find(key string) (Keykeeper, error) {
	if n == nil {
		return nil, fmt.Errorf("key value does not exist")
	}

	if key < n.Key {
		return n.Left.find(key)
	}
	if key > n.Key {
		return n.Right.find(key)
	}
	return n.Value, nil
}
