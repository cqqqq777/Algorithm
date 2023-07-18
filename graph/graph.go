package graph

type Edge struct {
	Weight int
	From   *Node // edge's start node
	To     *Node // edge's end node
}

type Node struct {
	Value int
	In    int     // how many edges end at this node
	Out   int     // how many edges start at this node
	Nexts []*Node // adjacent nodes
	Edges []*Edge // edges which associated with this node
}

type Graph struct {
	Nodes map[int]*Node      // key is the node number
	Edges map[*Edge]struct{} // the set of edges
}

// GenerateGraph 提供一个表示边的矩阵来生成一个图
// 每条边由一个长度为3的一维数组表示，分别为权重，起始点，终点
func GenerateGraph(matrix [][]int) *Graph {
	graph := new(Graph)
	graph.Nodes = make(map[int]*Node)
	graph.Edges = make(map[*Edge]struct{})
	for i := 0; i < len(matrix); i++ {
		weight, from, to := matrix[i][0], matrix[i][1], matrix[i][2]
		if _, ok := graph.Nodes[from]; !ok {
			graph.Nodes[from] = &Node{Value: from}
		}
		if _, ok := graph.Nodes[to]; !ok {
			graph.Nodes[to] = &Node{Value: to}
		}
		fromNode, toNode := graph.Nodes[from], graph.Nodes[to]
		edge := &Edge{Weight: weight, From: fromNode, To: toNode}
		graph.Edges[edge] = struct{}{}
		fromNode.Out++
		fromNode.Nexts = append(fromNode.Nexts, toNode)
		fromNode.Edges = append(fromNode.Edges, edge)
		toNode.In++
		toNode.Edges = append(toNode.Edges, edge)
	}
	return graph
}

type Queue struct{}

func (q *Queue) Add(_ *Node) {}

func (q *Queue) Poll() *Node { return nil }

func (q *Queue) IsEmpty() bool { return true }

type Stack struct{}

func (s *Stack) Push(_ *Node) {}

func (s *Stack) Pop() *Node { return nil }

func (s *Stack) IsEmpty() bool { return true }
