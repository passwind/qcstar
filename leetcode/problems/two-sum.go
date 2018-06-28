// 两数之和问题 https://leetcode-cn.com/problems/two-sum/description/
// https://leetcode-cn.com/submissions/detail/3589865/
package problems

func twoSum(nums []int, target int) []int {
	numidxs := []int{}
	for i := 0; i < len(nums); i++ {
		numidxs = append(numidxs, i)
	}

	for i := 0; i < len(numidxs)-1; i++ {
		for j := 1; j < len(numidxs); j++ {
			if nums[numidxs[i]] > nums[numidxs[j]] {
				tn := numidxs[j]
				numidxs[j] = numidxs[i]
				numidxs[i] = tn
			}
		}
	}

	for i := 0; i < len(numidxs)-1; i++ {
		for j := len(numidxs) - 1; j > i; j-- {
			if nums[numidxs[i]]+nums[numidxs[j]] == target {
				return []int{numidxs[i], numidxs[j]}
			}
		}
	}

	return []int{}
}
