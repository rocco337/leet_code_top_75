package countbits

//Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.
//It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
//Space complexity should be O(n).
//Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.
func CountBits(num int) []int {
	bits := make([]int, num+1)
	for n := 1; n <= num; n++ {
		bits[n] = bits[n&(n-1)] + 1
	}

	return bits
}
