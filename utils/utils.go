package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func ReadNumber(low uint8, high uint8, msg string) (result uint8) {
	fmt.Print(msg, ": ")
	scanner := bufio.NewScanner(os.Stdin)
	for result == 0 {
		scanner.Scan()
		line := scanner.Text()
		if len(line) > 3 {
			fmt.Printf("Enter correct number from %d to %d: ", low, high)
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Enter correct number from %d to %d: ", low, high)
			continue
		}
		result = uint8(number)
		if result < low || result > high {
			fmt.Printf("Enter correct number from %d to %d: ", low, high)
			result = 0
			continue
		}
	}
	return
}

func RandomNumber(low uint8, high uint8) uint8 {
	return uint8(rand.Intn(int(high-low))) + low
}
