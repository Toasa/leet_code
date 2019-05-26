func maxSatisfied(customers []int, grumpy []int, X int) int {
	l := len(customers)

	tmp_max := 0

	no_grumpy := 0

	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			tmp_max += customers[i]
		} else {
			no_grumpy += customers[i]
		}
	}
	head := 1
	tail := X
	tmp := tmp_max

	for i := 0; i < l-X; i++ {
		if grumpy[head-1] == 1 {
			tmp -= customers[head-1]
		}
		if grumpy[tail] == 1 {
			tmp += customers[tail]
		}
		head++
		tail++
		if tmp_max < tmp {
			tmp_max = tmp
		}

		if grumpy[i+X] == 0 {
			no_grumpy += customers[i+X]
		}

	}

	return tmp_max + no_grumpy
}