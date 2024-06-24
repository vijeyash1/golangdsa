package palindrome

import "fmt"

func CheckPalindrome(num int) {
	input_num := num
	var remainder int
	var res int

	for num > 0 {
		remainder = num % 10
		res = (res * 10) + remainder
		num = num / 10
	}

	if input_num == res {
		fmt.Println("its a palindrome")
	} else {
		fmt.Println("not a palindrome")
	}

}
