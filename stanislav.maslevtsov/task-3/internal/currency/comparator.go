package currency

func Compare(right, left *Currency) int {
	if right.Value > left.Value {
		return -1
	} else if right.Value < left.Value {
		return 1
	}

	return 0
}
