package hammingweight

//Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).
func HammingWeight(num uint32) int {
	counter := 0
	i := 0
	for i < 32 {
		mask := num >> i

		if mask&1 == 1 {
			counter++
		}
		i++
	}
	return counter
}
