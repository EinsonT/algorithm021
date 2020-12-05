func trap(height []int) int {
	stack := []int{}
	sum := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			//jisuanshuiliang
			h := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := i - stack[len(stack)-1] - 1
			sum += (min(height[i], height[stack[len(stack)-1]]) - h) * distance
		}
		stack = append(stack, i)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}