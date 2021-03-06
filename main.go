package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func main() {

	argArray := os.Args[1:] // read comamnd line

	// program loop

	if argArray[0][len(argArray[0])-4:] != ".txt" {
		panic("Input to be a text file")
	}

	data, err := ioutil.ReadFile(argArray[0])

	if err != nil {
		fmt.Printf("%v", err)
	} else {
		// if a 2nd arguement was supplied then we are converting characters to BINARY
		// binary to characters is the default
		if len(argArray) < 2 {

			// * BINARY CONVERSION STARTS HERE
			bit_counter := -1 // counts expontent of 2 of each bit
			sum := 0          // running byte sum
			converted_string := ""
			for i := len(data) - 1; i >= 0; i-- {
				digit := string(rune(data[i]))
				bit_counter++

				curr_bit := math.Exp2(float64(bit_counter))

				if digit == "1" {
					sum += int(curr_bit)
				}

				if bit_counter == 7 {

					bit_counter = -1 // starts at -1 so is 0 at beginnign of loop

					converted_string = string(sum) + converted_string
					sum = 0

				}

			}
			// * Prints to stdout here, in future add file write
			fmt.Printf("%v \n", converted_string)

		} else {

			fmt.Printf("NOT IMPLEMENTED! \n\n")

		}
	}

}
