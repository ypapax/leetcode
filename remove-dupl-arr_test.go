package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	type testCase struct {
		inp     []int
		expArr  []int
		expLeft int
	}
	cases := []testCase{
		{inp: []int{1,1,2}, expArr: []int{1,2}, expLeft: 2},
		{inp: []int{0,0,1,1}, expArr: []int{0,1}, expLeft: 2},
		{inp: []int{0,0,1,1,1}, expArr: []int{0,1}, expLeft: 2},
		{inp: []int{0,0,1,1,1,2,2}, expArr: []int{0,1,2}, expLeft: 3},
		{inp: []int{0,0,1,1,1,2,2,3,3,4}, expArr: []int{0,1,2,3,4}, expLeft: 5},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			t.Logf("c.inp 1: %+v", c.inp)
			defer func() {
				t.Logf("c.inp 2: %+v", c.inp)
			}()
			as := assert.New(t)
			actLeft := RemoveDuplicates(c.inp)
			if !as.True(len(c.inp) >= len(c.expArr)) {
				return
			}
			if !as.Equal(c.expArr, c.inp[:len(c.expArr)]) {
				return
			}
			if !as.Equal(c.expLeft, actLeft) {
				return
			}
		})
	}

}
