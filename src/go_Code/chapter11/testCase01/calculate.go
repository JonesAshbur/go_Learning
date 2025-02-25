package main

func addUpper(num1 int) int {
	res := 0
	for i := 0; i <= num1; i++ {
		res += i
	}
	return res
}
