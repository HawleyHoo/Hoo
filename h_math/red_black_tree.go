package main

import (
	"errors"
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"math/rand"
)

type (
	RBTree struct {
		root *node
	}

	node struct {
		color      bool
		leftNode   *node
		rightNode  *node
		fatherNode *node
		value      int
	}
)

const (
	RBTRed   = false
	RBTBlack = true
)

var tree *RBTree

func main() {
	test()
	return
	tree = &RBTree{}
	for i := 1; i <= 1000; i++ {
		x := rand.Intn(1000)
		tree.insert(x)
		preperties()
	}

	fmt.Println("tree:", tree.root.value)
	fmt.Println("find:", *(tree.find(425)))
}

func test() {
	tree := redblacktree.NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	fmt.Println(tree.Left())  // 1
	fmt.Println(tree.Right()) // 7

	fmt.Println(tree)
	// RedBlackTree
	//│   ┌── 7
	//└── 6
	//    │   ┌── 5
	//    └── 4
	//        │   ┌── 3
	//        └── 2
	//            └── 1
	fmt.Println(tree.Ceiling(8)) // <nil> false
	fmt.Println(tree.Floor(8))   // 7 true
}

func preperties() {
	//the numbers 2,4,5 need to test
	//2
	if !isBlack(tree.root) {
		fmt.Println("tree's root is not black")
	}
	//4
	err := colorOfChildren(tree.root)
	if err != nil {
		fmt.Println("red nodes have red children")
	}
	//5
	_, err = theNumOfBlack(tree.root)
	if err != nil {
		fmt.Println("tree's num of black nodes are different")
	}
}
func colorOfChildren(n *node) (err error) {
	if n == nil {
		return
	}
	err = colorOfChildren(n.leftNode)
	if err != nil {
		return errors.New("the forth property is destroyed")
	}
	err = colorOfChildren(n.rightNode)
	if err != nil {
		return errors.New("the forth property is destroyed")
	}
	if n.color == RBTRed {
		if isBlack(n.leftNode) && isBlack(n.rightNode) {
			return
		} else {
			return errors.New("the forth property is destroyed")
		}
	}
	return
}

func theNumOfBlack(n *node) (num int, err error) {
	if n == nil {
		return 0, nil
	}
	leftNum, err := theNumOfBlack(n.leftNode)
	if err != nil {
		return 0, errors.New("the fifth property is destroyed")
	}
	rightNum, err := theNumOfBlack(n.rightNode)
	if err != nil {
		return 0, errors.New("the fifth property is destroyed")
	}
	if leftNum != rightNum {
		return 0, errors.New("the fifth property is destroyed")
	}
	if n.color == RBTBlack {
		return leftNum + 1, nil
	} else {
		return leftNum, nil
	}
}

func (t *RBTree) leftRotate(n *node) {
	rn := n.rightNode
	// first give n's father to rn's father
	rn.fatherNode = n.fatherNode
	if n.fatherNode != nil {
		if n.fatherNode.leftNode == n {
			n.fatherNode.leftNode = rn
		} else {
			n.fatherNode.rightNode = rn
		}
	} else {
		t.root = rn
	}

	n.fatherNode = rn
	n.rightNode = rn.leftNode
	if n.rightNode != nil {
		n.rightNode.fatherNode = n
	}

	rn.leftNode = n
}

func (t *RBTree) rightRotate(n *node) {
	ln := n.leftNode
	ln.fatherNode = n.fatherNode
	if n.fatherNode != nil {
		if n.fatherNode.leftNode == n {
			n.fatherNode.leftNode = ln
		} else {
			n.fatherNode.rightNode = ln
		}
	} else {
		t.root = ln
	}

	n.fatherNode = ln
	n.leftNode = ln.rightNode
	if n.leftNode != nil {
		n.leftNode.fatherNode = n
	}
	ln.rightNode = n
}

func (t *RBTree) insert(v int) {
	if t.root == nil {
		t.root = &node{value: v, color: RBTBlack}
		return
	}
	n := t.root

	insertNode := &node{value: v, color: RBTRed}
	var nf *node
	for n != nil {
		nf = n
		if v < n.value {
			n = n.leftNode
		} else if v > n.value {
			n = n.rightNode
		} else {
			// TODO
			fmt.Println("todo: fix the condition that replace value.")
			return
		}
	}

	//设置新插入节点的父节点
	insertNode.fatherNode = nf
	//将新的节点挂到父节点上
	if v < nf.value {
		nf.leftNode = insertNode
	} else {
		nf.rightNode = insertNode
	}
	t.insertFixUp(insertNode)
}

func (t *RBTree) insertFixUp(n *node) {
	for !isBlack(n.fatherNode) {
		uncleNode := findBroNode(n.fatherNode)
		if !isBlack(uncleNode) {
			n.fatherNode.color = RBTBlack
			uncleNode.color = RBTBlack
			uncleNode.fatherNode.color = RBTRed
			n = n.fatherNode.fatherNode
		} else if n.fatherNode == n.fatherNode.fatherNode.leftNode {
			if n == n.fatherNode.leftNode {
				n.fatherNode.fatherNode.color = RBTRed
				n.fatherNode.color = RBTBlack
				n = n.fatherNode.fatherNode
				t.rightRotate(n)
			} else {
				n = n.fatherNode
				t.leftRotate(n)
			}
		} else {
			if n == n.fatherNode.rightNode {
				n.fatherNode.fatherNode.color = RBTRed
				n.fatherNode.color = RBTBlack
				n = n.fatherNode.fatherNode
				t.leftRotate(n)
			} else {
				n = n.fatherNode
				t.rightRotate(n)
			}
		}
		t.root.color = RBTBlack
	}
}

func isBlack(n *node) bool {
	if n == nil {
		return true
	} else {
		return n.color == RBTBlack
	}
}

func setColor(n *node, color bool) {
	if n == nil {
		return
	}
	n.color = color
}

func findBroNode(n *node) (bro *node) {
	if n.fatherNode == nil {
		return nil
	}
	if n.fatherNode.leftNode == n {
		bro = n.fatherNode.rightNode
	} else if n.fatherNode.rightNode == n {
		bro = n.fatherNode.leftNode
	} else {
		if n.fatherNode.leftNode == nil {
			bro = n.fatherNode.rightNode
		} else {
			bro = n.fatherNode.leftNode
		}
	}
	return bro
}

func (t *RBTree) delete(v int) {
	n := t.find(v)
	if n == nil {
		return
	}

	//copy color of fixNode
	var fixColor = n.color
	//if fixNode == nil copy node of start fix node
	//set it's father and set color black
	var fixNode = &node{fatherNode: n.fatherNode, color: RBTBlack}
	if n.leftNode == nil {
		t.transplant(n, n.rightNode)
		if n.rightNode != nil {
			fixNode = n.rightNode
		}
	} else if n.rightNode == nil {
		t.transplant(n, n.leftNode)
		if n.leftNode != nil {
			fixNode = n.leftNode
		}
	} else {
		succNode := t.miniNum(n.rightNode)
		fixColor = succNode.color
		if succNode.rightNode == nil {
			if succNode.fatherNode != n {
				fixNode = &node{fatherNode: succNode.fatherNode, color: RBTBlack}
			} else {
				fixNode = &node{fatherNode: succNode, color: RBTBlack}
			}
		} else {
			fixNode = succNode.rightNode
		}

		if succNode.fatherNode != n {
			t.transplant(succNode, succNode.rightNode)
			succNode.rightNode = n.rightNode
			succNode.rightNode.fatherNode = succNode
		}

		t.transplant(n, succNode)
		succNode.leftNode = n.leftNode
		succNode.leftNode.fatherNode = succNode
		succNode.color = n.color
	}

	if fixColor == RBTBlack {
		t.deleteFixUp(fixNode)
	}

}

func (t *RBTree) deleteFixUp(n *node) {
	if t.root == nil {
		return
	}
	for n != t.root && isBlack(n) {
		bro := findBroNode(n)
		if bro != n.fatherNode.leftNode {
			if !isBlack(bro) {
				n.fatherNode.color = RBTRed
				bro.color = RBTBlack
				t.leftRotate(n.fatherNode)
				bro = findBroNode(n)
			}

			//if bro is black its children perhaps be nil
			//if bro's children are black
			// n up
			if isBlack(bro.leftNode) && isBlack(bro.rightNode) {
				setColor(bro, RBTRed)
				n = n.fatherNode
			} else {
				if !isBlack(bro.rightNode) {
					bro.color = n.fatherNode.color
					bro.rightNode.color = RBTBlack
					n.fatherNode.color = RBTBlack
					t.leftRotate(n.fatherNode)
					n = t.root
				} else {
					bro.color = RBTRed
					bro.leftNode.color = RBTBlack
					t.rightRotate(bro)
					bro = findBroNode(n)
				}
			}

		} else {
			if !isBlack(bro) {
				n.fatherNode.color = RBTRed
				bro.color = RBTBlack
				t.rightRotate(n.fatherNode)
				bro = findBroNode(n)
			}

			if isBlack(bro.leftNode) && isBlack(bro.rightNode) {
				setColor(bro, RBTRed)
				n = n.fatherNode
			} else {
				if !isBlack(bro.leftNode) {
					bro.color = n.fatherNode.color
					bro.leftNode.color = RBTBlack
					n.fatherNode.color = RBTBlack
					t.rightRotate(n.fatherNode)
					break
				} else {
					bro.color = RBTRed
					bro.rightNode.color = RBTBlack
					t.leftRotate(bro)
				}

			}
		}

	}
	n.color = RBTBlack

}

func (t *RBTree) miniNum(n *node) *node {
	for n.leftNode != nil {
		n = n.leftNode
	}
	return n
}

func (t *RBTree) transplant(u, v *node) {
	if u.fatherNode == nil {
		t.root = v
		if v != nil {
			v.fatherNode = nil
		}
	} else if u == u.fatherNode.leftNode {
		u.fatherNode.leftNode = v
	} else {
		u.fatherNode.rightNode = v
	}
	if v != nil {
		v.fatherNode = u.fatherNode
	}
}

func (t *RBTree) find(v int) *node {
	n := t.root
	count := 0
	for n != nil {
		count++
		fmt.Println("times:", count)
		if v < n.value {
			//小于当前节点的话，往左节点找
			n = n.leftNode
		} else if v > n.value {
			//大于当前节点的话，往右节点找
			n = n.rightNode
		} else {
			//等于的话表示找到，返回
			return n
		}
	}
	//循环结束没找到，返回
	return nil
}
