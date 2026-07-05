package random

import "math/rand/v2"

func RandomNumber(num int) int {
	if num <= 0 {
		return 0
	}
	return rand.IntN(num)
}
