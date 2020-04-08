package sumoftwointegers

//Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -
func GetSum(a int, b int) int {

	for b != 0 {
		carry := a & b

		a = a ^ b

		b = carry << 1
	}

	return a
}
