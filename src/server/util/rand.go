package util

import (
	"math/rand"
	"time"
)

var myRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandNum ...
func RandNum(num int32) int32 {
	return myRand.Int31n(num)
}

// RandomBetween ...
func RandomBetween(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return myRand.Intn(max-min+1) + min
}

// RandomBetweenZero ...
func RandomBetweenZero(min, max int) int {
	if min >= max {
		return max
	}
	return myRand.Intn(max-min+1) + min
}

// RandomBetween31n ...
func RandomBetween31n(min, max int32) int32 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return myRand.Int31n(max-min+1) + min
}

// IsProbIntn ...
func IsProbIntn(percent int) bool {
	return myRand.Intn(100) < percent
}

// IsProbInt31n ...
func IsProbInt31n(percent int32) bool {
	return myRand.Int31n(100) < percent
}

// GetRandomN ...
func GetRandomN(total, num int) []int {
	if total < num {
		num = total
	}
	idx := myRand.Perm(total)
	res := make([]int, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, idx[i])
	}
	return res
}
