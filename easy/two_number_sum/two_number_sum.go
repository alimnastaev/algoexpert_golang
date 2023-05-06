package two_number_sum

func Run(array []int, target int) (result []int) {
	m := map[int]struct{}{}

	for _, el := range array {
		x := target - el

		if _, found := m[x]; found {
			return append(result, el, x)
			
		}

		m[el] = struct{}{}
	}

	return
}
