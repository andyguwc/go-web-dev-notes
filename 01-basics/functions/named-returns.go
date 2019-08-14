
// when using named returns they are passed to function as regular declared variables and can be modified 
// upon return, the last assigned value is returned 
func div(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
	q++
	r = r - dvsr
	return
}


