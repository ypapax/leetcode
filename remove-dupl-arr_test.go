package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		inp     []int
		expArr  []int
		expLeft int
	}
	cases := []testCase{{inp: []int{1,1,2}, expArr: []int{1,2}, expLeft: 2}}
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
