// variadic (variable length argument) by affixing ... before parameter's type

func avg(nums ...float64) float64 {
	n := len(nums)
	t := 0.0
	for _, v := range nums {
		t +=v 
	}

	return t / float64(n)
}

points := []float64{1,2,3,4}
fmt.Printf("sum(%v) = %f\n", points, sum(points...))
