package binary_tree

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stack struct{}

func (s *Stack) Push(_ *Node) {}

func (s *Stack) Pop() *Node { return nil }

func (s *Stack) IsEmpty() bool { return true }

func (s *Stack) Peek() *Node { return nil }

type Queue struct{}

func (q *Queue) Add(_ *Node) {}

func (q *Queue) Poll() *Node { return nil }

func (q *Queue) IsEmpty() bool { return true }

// PreOrderTraversal1 递归方式先序遍历二叉树
func PreOrderTraversal1(head *Node) {
	if head == nil {
		return
	}

	fmt.Println(head.Data)
	PreOrderTraversal1(head.Left)
	PreOrderTraversal1(head.Right)
}

// PreOrderTraversal2 使用栈的方式先序遍历二叉树
func PreOrderTraversal2(head *Node) {
	if head == nil {
		return
	}

	stack := new(Stack)
	stack.Push(head)
	for !stack.IsEmpty() {
		head = stack.Pop()
		fmt.Println(head.Data)
		if head.Right != nil {
			stack.Push(head.Right)
		}
		if head.Left != nil {
			stack.Push(head.Left)
		}
	}
}

// InOrderTraversal1 递归方式中序遍历二叉树
func InOrderTraversal1(head *Node) {
	if head != nil {
		return
	}

	InOrderTraversal1(head.Left)
	fmt.Println(head.Data)
	InOrderTraversal1(head.Right)
}

// InOrderTraversal2 使用栈的方式中序遍历二叉树
func InOrderTraversal2(head *Node) {
	if head == nil {
		return
	}

	stack := new(Stack)
	for !stack.IsEmpty() || head != nil {
		if head != nil {
			stack.Push(head)
			head = head.Left
		} else {
			head = stack.Pop()
			fmt.Println(head.Data)
			head = head.Right
		}
	}
}

// PosOrderTraversal1 递归方式后序遍历二叉树
func PosOrderTraversal1(head *Node) {
	if head == nil {
		return
	}

	PosOrderTraversal1(head.Left)
	PosOrderTraversal1(head.Right)
	fmt.Println(head.Data)
}

// PosOrderTraversal2 使用栈的方式后序遍历二叉树
func PosOrderTraversal2(head *Node) {
	if head == nil {
		return
	}

	stack1 := new(Stack)
	stack2 := new(Stack)
	stack1.Push(head)
	for !stack1.IsEmpty() {
		head = stack1.Pop()
		stack2.Push(head)
		if head.Left != nil {
			stack1.Push(head.Left)
		}
		if head.Right != nil {
			stack2.Push(head.Right)
		}
	}
	for !stack2.IsEmpty() {
		fmt.Println(stack2.Pop().Data)
	}
}

// BSTraversal 宽度优先遍历
func BSTraversal(head *Node) {
	if head == nil {
		return
	}

	queue := new(Queue)
	queue.Add(head)
	for !queue.IsEmpty() {
		cur := queue.Poll()
		fmt.Println(cur.Data)
		if cur.Left != nil {
			queue.Add(cur.Left)
		}
		if cur.Right != nil {
			queue.Add(cur.Right)
		}
	}
}

// GetMaxNodesAndLevel1 获取二叉树节点最多的层数和该层的节点数
func GetMaxNodesAndLevel1(head *Node) (int, int) {
	if head == nil {
		return 0, 0
	}

	queue := new(Queue)
	queue.Add(head)
	levelMap := make(map[*Node]int)
	levelMap[head] = 1

	curLevel, curLevelNodes, max, maxLevel := 1, 0, -1, 1
	for !queue.IsEmpty() {
		cur := queue.Poll()
		if curLevel == levelMap[cur] {
			curLevelNodes++
		} else {
			if max < curLevelNodes {
				max = curLevelNodes
				maxLevel = curLevel
			}
			curLevel++
			curLevelNodes = 1
		}

		if cur.Left != nil {
			queue.Add(cur.Left)
			levelMap[cur.Left] = curLevel + 1
		}
		if cur.Right != nil {
			queue.Add(cur.Right)
			levelMap[cur.Right] = curLevel + 1
		}
	}

	return max, maxLevel
}

// GetMaxNodesAndLevel2 获取二叉树的节点最多的层数和该层的节点数
func GetMaxNodesAndLevel2(head *Node) (int, int) {
	if head == nil {
		return 0, 0
	}

	queue := new(Queue)
	queue.Add(head)

	// cueEnd and nextEnd used to storage the last node in the current level and next level
	var curEnd, nextEnd *Node = head, nil
	max, maxLevel, curLevel, curLevelNodes := 0, 0, 1, 0

	for !queue.IsEmpty() {
		cur := queue.Poll()
		if cur.Left != nil {
			queue.Add(cur.Left)
			nextEnd = cur.Left
		}
		if cur.Right != nil {
			queue.Add(cur.Right)
			nextEnd = cur.Right
		}

		// cur == curEnd means coming to the last node in the current level
		if cur == curEnd {
			// update curEnd and nextEnd
			curEnd = nextEnd
			nextEnd = nil
			curLevelNodes++
			if max < curLevelNodes {
				// update max and maxLevel
				max = curLevelNodes
				maxLevel = curLevel
			}
			// reset the curLevelNodes to storage the next level's nodes num
			curLevelNodes = 0
			// update curLevel
			curLevel++
		} else {
			curLevelNodes++
		}
	}
	return max, maxLevel
}

type ReturnData struct {
	IsBST    bool
	Min, Max int
}

// IsBST 递归判断一颗二叉树是不是二叉搜索树
func IsBST(head *Node) bool {
	return IsBSTProcess(head).IsBST
}

func IsBSTProcess(head *Node) *ReturnData {
	if head == nil {
		return nil
	}

	result := &ReturnData{IsBST: true, Max: head.Data, Min: head.Data}

	leftData := IsBSTProcess(head.Left)
	rightData := IsBSTProcess(head.Right)

	// get current tree's max and min
	if leftData != nil {
		result.Min = int(math.Min(float64(leftData.Min), float64(head.Data)))
		result.Max = int(math.Max(float64(leftData.Max), float64(head.Data)))
	}
	if rightData != nil {
		result.Min = int(math.Min(float64(rightData.Min), float64(head.Data)))
		result.Max = int(math.Max(float64(rightData.Max), float64(head.Data)))
	}

	// check whether current tree is a BST
	if leftData != nil && (!leftData.IsBST || leftData.Max >= head.Data) {
		result.IsBST = false
	}
	if rightData != nil && (!rightData.IsBST || rightData.Min <= head.Data) {
		result.IsBST = false
	}

	return result
}

// IsBST2 判断一颗树是不是二叉搜索树
func IsBST2(head *Node) bool {
	nodes := make([]*Node, 0)

	process(head, nodes)
	for i := 1; i < len(nodes); i++ {
		if nodes[i].Data < nodes[i-1].Data {
			return false
		}
	}

	return true
}

func process(head *Node, nodes []*Node) {
	if head == nil {
		return
	}

	process(head.Left, nodes)
	nodes = append(nodes, head)
	process(head.Right, nodes)
}

// IsBST3 使用 Morris 算法
// 注意不要中途返回，让二叉树还原
func IsBST3(head *Node) bool {
	if head == nil {
		return true
	}
	cur := head
	preValue := math.MinInt
	result := true
	for cur != nil {
		if cur.Left == nil {
			if preValue >= cur.Data {
				result = false
			}
			preValue = cur.Data
			cur = cur.Right // 如果左节点为空就往右走
		} else {
			mostRight := cur.Left // 找到左子树上的最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur // 用最右节点的右指针标记 cur 用于下次跳回 cur
				cur = cur.Left        // cur 往左走
			} else {
				mostRight.Right = nil // 这是第二次来到 cur
				if preValue >= cur.Data {
					result = false
				}
				preValue = cur.Data
				cur = cur.Right // cur 往右走
			}
		}
	}
	return result
}

// IsCBT 判断是否是一颗完全二叉树
func IsCBT(head *Node) bool {
	if head == nil {
		return true
	}

	queue := new(Queue)
	queue.Add(head)

	// leaf is set to true when the right child node is nil the first time it is encountered
	leaf := false
	for !queue.IsEmpty() {
		cur := queue.Poll()
		if (cur.Right != nil && cur.Left == nil) || (leaf && (cur.Left != nil || cur.Right != nil)) {
			return false
		}
		if cur.Left != nil {
			queue.Add(cur.Left)
		}
		if cur.Right != nil {
			queue.Add(cur.Right)
		} else {
			leaf = true
		}
	}
	return true
}

// IsFBT1 判断一颗二叉树是不是满二叉树
func IsFBT1(head *Node) bool {
	if head == nil {
		return false
	}

	queue := new(Queue)
	queue.Add(head)
	nodes, level := 0, 1
	for !queue.IsEmpty() {
		cur := queue.Poll()
		nodes++
		if cur.Left != nil {
			queue.Add(cur.Left)
		}
		if cur.Right != nil {
			queue.Add(cur.Right)
		}
		if cur.Left != nil || cur.Right != nil {
			level++
		}
	}

	if nodes != int(math.Pow(2, float64(level))) {
		return false
	}
	return true
}

// IsFBT2 递归判断一颗树是不是满二叉树
func IsFBT2(head *Node) bool {
	ok, _ := FBTProcess(head)
	return ok
}

func FBTProcess(head *Node) (bool, int) {
	if head == nil {
		return true, 0
	}
	leftOk, leftNum := FBTProcess(head.Left)
	rightOk, rightNum := FBTProcess(head.Right)
	if !leftOk || !rightOk || leftNum != rightNum {
		return false, leftNum + rightNum + 1
	}
	return true, leftNum + rightNum + 1
}

// IsBalanced 判断一颗二叉树是否是平衡二叉树
func IsBalanced(head *Node) bool {
	ok, _ := IsBalancedProcess(head)
	return ok
}

func IsBalancedProcess(head *Node) (bool, int) {
	if head == nil {
		return true, 0
	}

	leftOk, leftHeight := IsBalancedProcess(head.Left)
	rightOk, RightHeight := IsBalancedProcess(head.Right)

	if !leftOk || !rightOk || math.Abs(float64(leftHeight-RightHeight)) > 1 {
		return false, int(math.Max(float64(leftHeight), float64(RightHeight))) + 1
	}

	return true, int(math.Max(float64(leftHeight), float64(RightHeight))) + 1
}

// LowestCommonAncestor1 找到最低公共祖先
func LowestCommonAncestor1(head, node1, node2 *Node) *Node {
	if head == nil || head == node1 || head == node2 {
		return head
	}

	leftNode := LowestCommonAncestor1(head.Left, node1, node2)
	rightNode := LowestCommonAncestor1(head.Right, node1, node2)

	if leftNode != nil && rightNode != nil {
		return head
	}

	if leftNode == nil {
		return rightNode
	}
	return leftNode
}

// LowestCommonAncestor2 使用 map，找到公共最低祖先
func LowestCommonAncestor2(head, node1, node2 *Node) *Node {
	fatherMap := make(map[*Node]*Node)
	fatherSet := make(map[*Node]struct{})

	LCAProcess(head, fatherMap)

	cur := node1
	fatherSet[cur] = struct{}{}
	for cur != head {
		fatherSet[fatherMap[cur]] = struct{}{}
		cur = fatherMap[cur]
	}

	cur = node2
	for cur != head {
		_, ok := fatherSet[cur]
		if ok {
			return cur
		}
		cur = fatherMap[cur]
	}
	return head
}

func LCAProcess(head *Node, fatherMap map[*Node]*Node) {
	if head.Left != nil {
		fatherMap[head.Left] = head
		LCAProcess(head.Left, fatherMap)
	}
	if head.Right != nil {
		fatherMap[head.Right] = head
		LCAProcess(head.Right, fatherMap)
	}
}

type NewNode struct {
	Data                int
	Parent, Left, Right *NewNode
}

// GetSuccessorNode 在一个含有父节点指针的二叉树中获取 node 的后继节点
func GetSuccessorNode(node *NewNode) *NewNode {
	if node == nil {
		return node
	}

	if node.Right != nil {
		// get the leftest node
		head := node.Right
		for head.Left != nil {
			head = head.Left
		}
		return head
	} else {
		cur := node.Parent
		for cur.Parent != nil && cur.Left != node {
			cur = cur.Parent
			node = node.Parent
		}
		return cur
	}
}

// MarshalBTToString 将二叉树按先序序列化为一个字符串
func MarshalBTToString(head *Node) string {
	return MProcess(head)
}

func MProcess(head *Node) string {
	if head == nil {
		return "*!"
	}

	leftRes := MProcess(head.Left)
	rightRes := MProcess(head.Right)

	return fmt.Sprintf("%d!%s%s", head.Data, leftRes, rightRes)
}

type StringQueue struct {
}

func (s *StringQueue) Add(_ string) {

}

func (s *StringQueue) Poll() string {
	return ""
}

func (s *StringQueue) IsEmpty() bool {
	return true
}

// UnMarshalStringToBT 将字符串反序列化到内存中的二叉树
func UnMarshalStringToBT(s string) *Node {
	data := strings.Split(s, "!")
	queue := new(StringQueue)
	for i := 0; i < len(data)-1; i++ {
		queue.Add(data[i])
	}
	return UProcess(queue)
}

func UProcess(queue *StringQueue) *Node {
	value := queue.Poll()
	if value == "!" {
		return nil
	}

	data, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}
	node := &Node{Data: int(data)}
	node.Left = UProcess(queue)
	node.Right = UProcess(queue)
	return node
}

// PaperFolds 解决折纸问题
func PaperFolds(n int) {
	// true used to express 'down'
	PaperFoldsProcess(n, true)
}

func PaperFoldsProcess(n int, b bool) {
	if n < 1 {
		return
	}
	PaperFoldsProcess(n-1, true)
	if b {
		fmt.Println("down")
	} else {
		fmt.Println("up")
	}
	PaperFoldsProcess(n-1, false)
}
