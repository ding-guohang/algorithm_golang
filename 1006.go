package main

/*
通常，正整数 n 的阶乘是所有小于或等于 n 的正整数的乘积。例如，factorial(10) = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1。

相反，我们设计了一个笨阶乘 clumsy：在整数的递减序列中，我们以一个固定顺序的操作符序列来依次替换原有的乘法操作符：乘法(*)，除法(/)，加法(+)和减法(-)。

例如，clumsy(10) = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1。然而，这些运算仍然使用通常的算术运算顺序：我们在任何加、减步骤之前执行所有的乘法和除法步骤，并且按从左到右处理乘法和除法步骤。

另外，我们使用的除法是地板除法（floor division），所以10 * 9 / 8。这保证结果是一个整数。

实现上面定义的笨函数：给定一个整数 N，它返回 N 的笨阶乘。

示例 1：

输入：4
输出：7
解释：7 = 4 * 3 / 2 + 1
示例 2：

输入：10
输出：12
解释：12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1

提示：

1 <= N <= 10000
-2^31 <= answer <= 2^31 - 1 （答案保证符合 32 位整数。）

*/

func calPrefix(a, b, c int) int {
	ret := a
	if b > 0 {
		ret = ret * b
	}
	if c > 0 {
		ret = ret / c
	}
	return ret
}

func clumsy(N int) int {
	// clumsy(10) = (10 * 9 / 8) + 7 - (6 * 5 / 4) + 3 - 2 * 1
	// (* /) + - (* /) + - (* /)
	// 3个数一组
	if N <= 4 {
		ret := calPrefix(N, N-1, N-2)
		if N-3 > 0 {
			ret += N - 3
		}
		return ret
	}

	ret := calPrefix(N, N-1, N-2) + N - 3 // N=10, ret=10*9/8+7=18
	i := N - 4                            // i=6
	for i > 0 {
		ret = ret - calPrefix(i, i-1, i-2) // ret=18-6*5/4=11
		if i-3 > 0 {                       //i-3=3
			ret = ret + i - 3 //ret=11+3=14
		}
		i = i - 4 //i=2
	}
	// calPrefix(N,N-1,N-2) + (N-3) - calPrefix(N-4,N-5,N-6) + (N-7) - calPrefix(N-8,N-9,N-10) + (N-11)
	return ret
}

