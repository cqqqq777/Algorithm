package winter19

// 教练使用整数数组 actions 记录一系列核心肌群训练项目编号。
// 为增强训练趣味性，需要将所有奇数编号训练项目调整至偶数编号训练项目之前。请将调整后的训练项目编号以 数组 形式返回。
func trainingPlan(actions []int) []int {
	if len(actions) <= 1 {
		return actions
	}
	p1, p2 := -1, len(actions)
	var cur int
	for cur < p2 {
		if actions[cur]&1 == 1 {
			actions[cur], actions[p1+1] = actions[p1+1], actions[cur]
			p1++
			cur++
		} else {
			actions[cur], actions[p2-1] = actions[p2-1], actions[cur]
			p2--
		}
	}
	return actions
}
