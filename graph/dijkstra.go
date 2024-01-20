package graph

import (
	"math"
)

// Dijkstra 求连通图中某一节点到其他所有节点的最短路径
func Dijkstra(head *Node) map[*Node]int {
	distanceMap := make(map[*Node]int)
	distanceMap[head] = 0
	unlockedSet := make(map[*Node]struct{})
	minNode := GetMinDistanceNode(distanceMap, unlockedSet)
	for minNode != nil {
		for _, edge := range minNode.Edges {
			distance := distanceMap[edge.To]
			if _, exist := distanceMap[edge.To]; !exist {
				distanceMap[edge.To] = distance + edge.Weight
			} else {
				distanceMap[edge.To] = int(math.Min(float64(distanceMap[edge.To]), float64(distance+edge.Weight)))
			}
		}
		unlockedSet[minNode] = struct{}{}
		minNode = GetMinDistanceNode(distanceMap, unlockedSet)
	}
	return distanceMap
}

func GetMinDistanceNode(distanceMap map[*Node]int, unlocked map[*Node]struct{}) *Node {
	var minNode *Node
	minDistance := math.MaxInt
	for k, distance := range distanceMap {
		if _, exist := unlocked[k]; !exist {
			if distance < minDistance {
				minNode = k
			}
		}
	}
	return minNode
}
