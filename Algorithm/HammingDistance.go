package Algorithm

func HanmingDistance(x int, y int) int {
	var max, min int
	if x > y {
		max = x
		min = y
	} else if x == y {
		return 0
	} else {
		max = y
		min = x
	}
	var i uint = 0
	num := 0
	for max>>i > 0 {
		if (max>>i)&1 != (min>>i)&1 {
			num++
		}
		i++
	}
	return num
}
