package leetcode

// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/727/
func RemoveDuplicates(nums []int) int {
	const nonVal = -200
	var (
		toDeleteCount     int
	)
	//var latestSameValue int
	i := -1
	for {
		i++
		if i > len(nums)-1 {
			//log.Printf("greater or equal than nums len: %+v", nums)
			break
		}
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			toDeleteCount++
			//latestSameValue=nums[i]
			//log.Printf("i: %+v, toDeleteCount: %+v, same value: %+v \n", i, toDeleteCount, nums[i])
			continue
		}
		//log.Println("toDeleteCount", toDeleteCount)
		if toDeleteCount > 0 {
			for j := i; j < len(nums); j++ {

				jBefore := j-toDeleteCount
				//log.Printf("i	      : %+v", i)
				//log.Printf("latestSameValue: '%+v'", latestSameValue)
				//log.Printf("toDeleteCount: %+v", toDeleteCount)
				//log.Printf("jBefore: '%+v'", jBefore)
				//log.Printf("j      : '%+v'", j)
				//valueBefore := nums[jBefore]
				//log.Printf("valueBefore: %+v", valueBefore)
				//log.Printf("valueAfter : %+v", nums[j])

				nums[jBefore]=nums[j]
			}
			for l := len(nums)-toDeleteCount; l<len(nums); l++ {
				nums[l]=nonVal
			}
			//iBefore := i
			i = i - toDeleteCount
			//log.Println("lastGoodIndex changing", lastGoodIndex)
			//log.Println("lastGoodIndex changed ", lastGoodIndex)
			//log.Printf("changing i from %+v to %+v, toDeleteCount: %+v\n", iBefore, toDeleteCount, i)

			toDeleteCount = 0
		}

	}
	var lastGoodLen = len(nums)
	for k, v := range nums {

		//log.Println("lastGoodLen", lastGoodLen)
		//log.Println("k", k)
		if v == nonVal {
			break
		}
		lastGoodLen = k+1

	}
	//log.Println("lastGoodLen", lastGoodLen)
	return lastGoodLen
}

