package Algorithm

func MyAtoi(str string) int {
	runes := []rune(str)
	positive := true
	var number int
	for _, v := range runes {
		if v == 32 {
			runes = runes[1:]
			continue
		}
		break
	}
	if len(runes) < 1 {
		return 0
	}
	if runes[0] == 45 {
		runes = runes[1:]
		positive = false
	} else if runes[0] == 43 {
		runes = runes[1:]
	}
	for i := len(runes) - 1; i >= 0; i-- {
		if int(runes[i])-48 > 9 || int(runes[i])-48 < 0 {
			runes = runes[:i]
		}
	}
	if len(runes) < 1 {
		return 0
	}
	var length int = len(runes)
	for _, v := range runes {
		if int(v)-48 > 9 || int(v)-48 < 0 {
			break
		}
		var times int = 1
		for i := length - 1; i >= 1; i-- {
			times = times * 10
		}
		if positive == true {
			if int(v) > 50 && times >= 1000000000 {
				return 2147483647
			}
			if 2147483647-number-(int(v)-48)*times < 0 {
				return 2147483647
			}
		} else {
			if int(v) > 50 && times >= 1000000000 {
				return -2147483648
			}
			if 2147483648-number-(int(v)-48)*times < 0 {
				return -2147483648
			}
		}

		number = number + (int(v)-48)*times
		length = length - 1
	}
	if positive == false {
		number = 0 - number
	}
	return number
}
