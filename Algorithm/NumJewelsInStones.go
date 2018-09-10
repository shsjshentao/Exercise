package Algorithm

func NumJewelsInStones(J string, S string) int {
	var num int
	for _,v1:=range []byte(J) {
		for _,v2 := range []byte(S) {
			if v1 == v2 {
				num++
			}
		}
	}
	return num
}