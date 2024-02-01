package winter15

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 记录每个节点的入度
	inMap := make(map[int]int)
	// 记录每个节点的邻接节点
	nexts := make(map[int][]int)

	for _, edge := range prerequisites {
		inMap[edge[0]]++
		next, ok := nexts[edge[1]]
		if !ok {
			next = make([]int, 0)
		}
		next = append(next, edge[0])
		nexts[edge[1]] = next
	}

	res := make([]int, 0, numCourses)
	for node := 0; node < numCourses; node++ {
		if inMap[node] == 0 {
			res = append(res, node)
		}
	}

	for cur := 0; cur < numCourses; cur++ {
		if cur == len(res) {
			return false
		}
		curNode := res[cur]
		next := nexts[curNode]
		for _, v := range next {
			inMap[v]--
			if inMap[v] == 0 {
				res = append(res, v)
			}
		}
	}

	return true
}
