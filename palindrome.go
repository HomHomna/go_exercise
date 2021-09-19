package main

import (
	"fmt"
)

func main() {
	var inputPalindrome string
	var statusPalindrome bool

	for {
		fmt.Println("Please enter your palindrome(exit to end program): ")
		fmt.Scanln(&inputPalindrome)

		if inputPalindrome == "exit" {
			fmt.Println("-----End the program-----")
			break
		} else if len(inputPalindrome) > 0 {
			for i := 0; i < len(inputPalindrome)/2; i++ {
				if inputPalindrome[i] == inputPalindrome[len(inputPalindrome)-(i+1)] {
					statusPalindrome = true

				} else {
					statusPalindrome = false
					break
				}
			}
			if statusPalindrome == true {
				fmt.Printf("%s is a palindrome.\n", inputPalindrome)

			} else {
				fmt.Printf("%s is not a palindrome.\n", inputPalindrome)

			}

		} else if len(inputPalindrome) == 0 {
			fmt.Println("Please enter number or text again.")

		} else {
			fmt.Println("Something wrong, please try again.")

		}

	}
}
