package winter11

// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
func majorityElement(nums []int) int {
	// 推论一： 若记 众数 的票数为 +1 ，非众数 的票数为 −1 ，则一定有所有数字的 票数和 >0 。
	// 推论二： 若数组的前 a 个数字的 票数和 = 0，则 数组剩余 (n−a) 个数字的 票数和一定仍 >0，即后 (n−a) 个数字的 众数仍为 x 。
	count, x := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == x {
			count++
		} else {
			count--
		}
		if count == 0 {
			x = nums[i]
		}
	}
	return x
}
