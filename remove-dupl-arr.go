package leetcode

import (
	"fmt"
	"log"
	"strings"
)

// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/727/
func RemoveDuplicates(nums []int) int {
	var (
		toDeleteCount     int
	)
	var latestSameValue int
	i := -1
	lastGoodIndex := 0
	for {
		i++
		log.Println("i", i)
		log.Println("toDeleteCount", toDeleteCount)
		//time.Sleep(time.Millisecond * 10)
		if i >= len(nums)-1 {
			log.Printf("greater or equal than nums len: %+v", nums)
			break
		}
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			toDeleteCount++
			latestSameValue=nums[i]
			log.Printf("i: %+v, toDeleteCount: %+v, same value: %+v \n", i, toDeleteCount, nums[i])
			continue
		}
		log.Println("toDeleteCount", toDeleteCount)
		if toDeleteCount > 0 {
			// 0 1 1 1 2
			//   1     4
			//     2 3
			for j := i; j < len(nums); j++ {

				jBefore := j-toDeleteCount
				log.Printf("i	      : %+v", i)
				log.Printf("latestSameValue: '%+v'", latestSameValue)
				log.Printf("toDeleteCount: %+v", toDeleteCount)
				log.Printf("jBefore: '%+v'", jBefore)
				log.Printf("j      : '%+v'", j)
				valueBefore := nums[jBefore]
				log.Printf("valueBefore: %+v", valueBefore)
				log.Printf("valueAfter : %+v", nums[j])

				log.Printf("nums before: %+v\n", strArr(nums))
				nums[jBefore]=nums[j]
				log.Printf("nums after : %+v\n\n\n", strArr(nums))
			}
			iBefore := i
			i = i - toDeleteCount
			log.Println("lastGoodIndex changing", lastGoodIndex)
			lastGoodIndex = i
			log.Println("lastGoodIndex changed ", lastGoodIndex)
			log.Printf("changing i from %+v to %+v, toDeleteCount: %+v\n", iBefore, toDeleteCount, i)

			toDeleteCount = 0
			latestSameValue=0
		}

	}
	log.Println("len(nums)", len(nums))
	log.Println("lastGoodIndex", lastGoodIndex)
	lastGoodLen := lastGoodIndex + 1
	lastGoodLenPlusSkippedOne := lastGoodLen + 1
	return lastGoodLenPlusSkippedOne
}

func strArr(nums []int) string {
	var vv []string
	for i, v := range nums {
		vv = append(vv, fmt.Sprintf("%+v:%+v", i, v))
	}
	return fmt.Sprintf("%+v   %+v", nums, strings.Join(vv, ", "))
}