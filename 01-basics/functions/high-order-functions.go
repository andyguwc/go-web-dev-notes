// high order functions
// go function can take another function as a parameter and also return a function as a result value

func apply(nums []int, f func(int) int) func() {
	for i, v := range nums {
		nums[i] = f(v)
	}
	return func() {
		fmt.Println(nums)
	}
}