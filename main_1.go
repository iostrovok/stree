package main

import (
	"flag"
	"fmt"
)

var dasedir string
var port int

type Node struct {
	Value       string
	Left, Right *Node
	Parent      *Node
}

var top, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13, n14 *Node

func init() {
	flag.StringVar(&dasedir, "dir", "./", "Root dir")
	flag.IntVar(&port, "port", 19720, "Port")
	flag.Parse()
}

func main() {
	InitTree()
	fmt.Println("WidthPrint: ")
	WidthPrint(top)
	fmt.Println("WidthRecursivePrint: ")
	WidthRecursivePrint(top)

	fmt.Println("\n-----------------------------\nInitTreeSymmetric: ")

	InitTreeSymmetric()
	fmt.Println("WidthPrint: ")
	WidthPrint(top)
	fmt.Println("WidthRecursivePrint: ")
	WidthRecursivePrint(top)
	fmt.Printf("CheckSymmetric: %t\n", CheckSymmetric(top))
}

func CheckSymmetric(top *Node) bool {
	if top == nil {
		return true
	}

	stackLeft := []*Node{}
	stackRight := []*Node{}

	if top.Left == nil && top.Right == nil {
		return true
	}

	if top.Left != nil && top.Right != nil {
		stackLeft = []*Node{top.Left, top.Right}
		stackRight = []*Node{top.Right, top.Left}
	} else {
		return false
	}

	var nextLeft, nextRight *Node

	for len(stackLeft) > 0 && len(stackRight) > 0 {
		nextLeft, stackLeft = _pop(stackLeft)
		nextRight, stackRight = _pop(stackRight)

		if nextLeft == nil && nextRight == nil {
			continue
		}

		if (nextLeft != nil && nextRight != nil) && (nextLeft.Value != nextRight.Value) {
			return false
		}

		stackLeft = append(stackLeft, nextLeft.Left)
		stackLeft = append(stackLeft, nextLeft.Right)

		stackRight = append(stackRight, nextRight.Right)
		stackRight = append(stackRight, nextRight.Left)
	}

	return true
}

func WidthPrint(n *Node) {
	if n == nil {
		return
	}

	stack := []*Node{n}
	var next *Node

	for len(stack) > 0 {
		next, stack = _pop(stack)

		fmt.Printf("-> %s\n", next.Value)
		if next.Left != nil {
			stack = append(stack, next.Left)
		}
		if next.Right != nil {
			stack = append(stack, next.Right)
		}
	}
}

func WidthRecursivePrint(n *Node) {
	if n == nil {
		return
	}

	fmt.Printf("-> %s ", n.Value)
	WidthRecursivePrint(n.Left)
	WidthRecursivePrint(n.Right)
}

func WidthRecursivePrintFromAnyPoint(n *Node, checkParent bool) {
	if n == nil {
		return
	}

	fmt.Printf("-> %s\n", n.Value)
	WidthRecursivePrintFromAnyPoint(n.Left, false)
	WidthRecursivePrintFromAnyPoint(n.Right, false)
	if checkParent {
		WidthRecursivePrintFromAnyPoint(n.Parent, true)
	}
}

func PrintAllPaths(n *Node, path string, checkParent bool) {
	if n == nil {
		return
	}

	path = path + " -> " + n.Value

	fmt.Printf("-> %s\n", path)
	PrintAllPaths(n.Left, path, false)
	PrintAllPaths(n.Right, path, false)
	if checkParent {
		PrintAllPaths(n.Parent, path, true)
	}
}

func _pop(stack []*Node) (*Node, []*Node) {

	if len(stack) == 0 {
		return nil, nil
	}

	if len(stack) == 1 {
		return stack[0], nil
	}

	out := stack[1:]
	return stack[0], out
}

func InitTree() {
	/*
			 				  top
				   1 					   2
			3 			4 			5 			6
		7 		8 	9		10	11		12	13 		14
	*/

	// 4 level
	n7 = &Node{
		Value: "n7",
	}
	n8 := &Node{
		Value: "n8",
	}
	n9 = &Node{
		Value: "n9",
	}
	n10 = &Node{
		Value: "n10",
	}
	n11 = &Node{
		Value: "n11",
	}
	n12 = &Node{
		Value: "n12",
	}
	n13 = &Node{
		Value: "n13",
	}
	n14 = &Node{
		Value: "n14",
	}
	// 3 level
	n3 = &Node{
		Value: "n3",
		Left:  n7,
		Right: n8,
	}

	n4 = &Node{
		Value: "n4",
		Left:  n9,
		Right: n10,
	}

	n5 = &Node{
		Value: "n5",
		Left:  n11,
		Right: n12,
	}

	n6 = &Node{
		Value: "n6",
		Left:  n13,
		Right: n14,
	}

	// 2 level
	n1 = &Node{
		Value: "n1",
		Left:  n3,
		Right: n4,
	}

	n2 = &Node{
		Value:  "n2",
		Left:   n5,
		Right:  n6,
		Parent: nil,
	}
	// t1 level
	top = &Node{
		Value:  "top",
		Left:   n1,
		Right:  n2,
		Parent: nil,
	}

	n1.Parent = top
	n2.Parent = top

	n3.Parent = n1
	n4.Parent = n1
	n5.Parent = n2
	n6.Parent = n2

	n7.Parent = n3
	n8.Parent = n3
	n9.Parent = n4
	n10.Parent = n4

	n11.Parent = n5
	n12.Parent = n5
	n13.Parent = n6
	n14.Parent = n6
}

func InitTreeSymmetric() {
	/*
			 				  top
				   1 					   2(1)
			3 			4 			5(4) 			6(3)
		7 		8 	9		10	11(10)		12(9) 13(8)  14(7)
	*/

	// 4 level
	n7 = &Node{
		Value: "n7",
	}
	n8 := &Node{
		Value: "n8",
	}
	n9 = &Node{
		Value: "n9",
	}
	n10 = &Node{
		Value: "n10",
	}
	n11 = &Node{
		Value: "n10", // n11
	}
	n12 = &Node{
		Value: "n9", // n12
	}
	n13 = &Node{
		Value: "n8", //
	}
	n14 = &Node{
		Value: "n7", //
	}
	// 3 level
	n3 = &Node{
		Value: "n3",
		Left:  n7,
		Right: n8,
	}

	n4 = &Node{
		Value: "n4",
		Left:  n9,
		Right: n10,
	}

	n5 = &Node{
		Value: "n4",
		Left:  n11,
		Right: n12,
	}

	n6 = &Node{
		Value: "n3",
		Left:  n13,
		Right: n14,
	}

	// 2 level
	n1 = &Node{
		Value: "n1",
		Left:  n3,
		Right: n4,
	}

	n2 = &Node{
		Value:  "n1",
		Left:   n5,
		Right:  n6,
		Parent: nil,
	}
	// t1 level
	top = &Node{
		Value:  "top",
		Left:   n1,
		Right:  n2,
		Parent: nil,
	}

	n1.Parent = top
	n2.Parent = top

	n3.Parent = n1
	n4.Parent = n1
	n5.Parent = n2
	n6.Parent = n2

	n7.Parent = n3
	n8.Parent = n3
	n9.Parent = n4
	n10.Parent = n4

	n11.Parent = n5
	n12.Parent = n5
	n13.Parent = n6
	n14.Parent = n6
}
