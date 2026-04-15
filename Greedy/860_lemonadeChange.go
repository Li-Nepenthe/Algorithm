package Greedy

func lemonadeChange(bills []int) bool {
	var five, ten int
	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			five++
		case 10:
			if five < 1 {
				return false
			}
			five--
			ten++
		case 20:
			//优先花掉十块的
			if five >= 1 && ten >= 1 {
				ten--
				five--
				continue
			}
			if five >= 3 {
				five -= 3
				continue
			}
			return false
		}
	}
	return true
}
