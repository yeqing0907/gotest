package main

import "fmt"

func main() {

	max := max(2, 10)
	fmt.Println(max)

	y, x := swap("hello", "world")
	fmt.Println(y, x)

	fmt.Printf("变量max的地址: %x\n", &max)

}

func max(num1, num2 int) int {
	var result int

	if (num1 > num2) {
		result = num1
	}else {
		result = num2
	}

	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
