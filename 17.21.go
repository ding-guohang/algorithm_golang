package main

/*
给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

*/

// 从左往右，每一个非0的为起点i，找到右边第一个比他大的j，中间所有都可以储存 val(左)-val(k) 的水量, 然后用j作为i继续
// 从右往左，反向来一次
func trap(height []int) int {
	// 找到左边第一个非0的
	start := 0
	ret := 0
	for start < len(height) {
		if height[start] == 0 {
			start++
			continue
		}

		next := findNext(height, start, true)
		if next < 0 {
			break
		}
		ret += sum(height, start, next, start)
		start = next
	}

	start = len(height) - 1
	for start >= 0 {
		if height[start] == 0 {
			start--
			continue
		}

		next := findNext(height, start, false)
		if next < 0 {
			break
		}
		ret += sum(height, next, start, start)
		start = next
	}
	return ret
}

func findNext(height []int, start int, order bool) int {
	if order {
		for i := start + 1; i < len(height); i++ {
			if height[i] >= height[start] {
				return i
			}
		}
		return -1
	}

	for i := start - 1; i >= 0; i-- {
		// 正向时包含相等，负向时忽略
		if height[i] > height[start] {
			return i
		}
	}
	return -1
}

func sum(height []int, left, right, target int) int {
	targetVal := height[target]
	ret := 0
	for i := left; i < right; i++ {
		val := targetVal - height[i]
		if val > 0 {
			ret += val
		}
	}
	return ret
}

