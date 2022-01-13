package main

import (
	"fmt"

	"github.com/u-shylianok/golang-basics/bst-extended/model"
)

func main() {
	var t model.Tree

	t.Insert(model.User{
		Username: "user_10",
		Name:     "name10",
		Password: "password10",
	})
	t.Insert(model.User{
		Username: "user_01",
		Name:     "name1",
		Password: "password1",
	})
	t.Insert(model.User{
		Username: "user_15",
		Name:     "name15",
		Password: "password15",
	})
	t.Insert(model.User{
		Username: "user_16",
		Name:     "name16",
		Password: "password16",
	})
	t.Insert(model.User{
		Username: "user_08",
		Name:     "name8",
		Password: "password8",
	})
	t.Insert(model.User{
		Username: "user_10",
		Name:     "name10",
		Password: "password10",
	})
	t.Insert(model.User{
		Username: "user_03",
		Name:     "name3",
		Password: "password3",
	})

	fmt.Println("INFIX_TRAVERSE example:")
	printInOrder(t.Root)
	fmt.Println("======")

	fmt.Println("PREFIX_TRAVERSE example:")
	printPreOrder(t.Root)
	fmt.Println("======")

	fmt.Println("POSTFIX_TRAVERSE example:")
	printPostOrder(t.Root)
	fmt.Println("======")

	key := "user_15"
	fmt.Printf("\nremove key = %v example\n", key)
	printInOrder(t.Root)
	fmt.Println("====== before\n")

	if err := t.Remove(key); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	printInOrder(t.Root)
	fmt.Println("====== after")

	key = "user_08"
	fmt.Printf("\nfind node with key = %v example\n", key)
	node, err := t.Find(key)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("node: %#v\n", node)

	fmt.Println()
}

func printPreOrder(n *model.Node) {
	if n == nil {
		return
	}

	fmt.Printf("%s->%v\n", n.Key, n.Value)
	printPreOrder(n.Left)
	printPreOrder(n.Right)
}

func printPostOrder(n *model.Node) {
	if n == nil {
		return
	}

	printPostOrder(n.Left)
	printPostOrder(n.Right)
	fmt.Printf("%s->%v\n", n.Key, n.Value)
}

func printInOrder(n *model.Node) {
	if n == nil {
		return
	}

	printInOrder(n.Left)
	fmt.Printf("%s->%v\n", n.Key, n.Value)
	printInOrder(n.Right)
}
