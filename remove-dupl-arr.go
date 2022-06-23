package leetcode

import "log"

// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/727/
func RemoveDuplicates(nums []int) int {
	var (
		toDeleteCount     int
		totalDeletedCount int
	)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			totalDeletedCount++
			toDeleteCount++
			continue
		}
		if toDeleteCount > 0 {
			// 0 1 1 1 2
			//   1     4
			//     2 3
			for j := i; j < len(nums); j++ {
				iBefore := j-toDeleteCount
				valueBefore := nums[iBefore]
				nums[iBefore]=nums[j]
				log.Printf("changing nums[%+v] from %+v to nums[%+v] %+v", iBefore, valueBefore, i, nums[i])
			}
			toDeleteCount = 0
		}
	}
	return len(nums)-totalDeletedCount
}