package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputSize string
	for {
		fmt.Print("Please enter your size between 1 and 20 (0 for exit): ")
		fmt.Scanln(&inputSize)
		if inputSize == "0" {
			break
		}
		size, err := strconv.ParseInt(inputSize, 0, 64)
		if err != nil {
			fmt.Println("Please enter a number We can't get string.")
		} else if size >= 1 && size <= 20 {
			{
				for row := 0; row < int(size); row++ {
					if row == 0 || row == (int(size)-1) {
						fmt.Println(strings.Repeat("* ", int(size)))
					} else {
						fmt.Print("*")
						fmt.Print(strings.Repeat(" ", 2*int(size)-3))
						fmt.Println("*")
					}
				}
			}
		} else {
			fmt.Println("Please enter a number greater than 0.")
		}

	}
}
