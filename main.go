package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Calculate the Luhn checksum for a number string.
func checksum(number string) int {
	sum := 0
	nDigits := len(number)
	parity := nDigits % 2

	for i := 0; i < nDigits; i++ {
		digit, _ := strconv.Atoi(string(number[i]))
		if i%2 == parity {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}
		sum += digit
	}

	return sum % 10
}

// Calculate the check digit using the Luhn algorithm.
func calcCheckDigit(number string) string {
	checksumValue := checksum(number + "0")
	if checksumValue == 0 {
		return "0"
	}
	return strconv.Itoa(10 - checksumValue)
}

// Ask for input and generate random IMEI numbers.
func main() {
	reader := bufio.NewReader(os.Stdin)

	// Loop until valid IMEI prefix (8-12 digits) is provided.
	var start string
	for {
		fmt.Print("Enter the first 8 - 12 digits: ")
		start, _ = reader.ReadString('\n')
		start = strings.TrimSpace(start)

		if _, err := strconv.Atoi(start); err == nil && len(start) >= 8 && len(start) <= 12 {
			break
		}

		fmt.Println("*** Invalid input: you must enter between 8 and 12 digits")
	}

	// Loop until valid count is provided.
	var count int
	for {
		fmt.Print("Enter the number of IMEI numbers to generate: ")
		countInput, _ := reader.ReadString('\n')
		countInput = strings.TrimSpace(countInput)

		if num, err := strconv.Atoi(countInput); err == nil && num > 0 {
			count = num
			break
		}

		fmt.Println("*** Invalid input: you must enter a number greater than zero")
	}

	// Generate random IMEI numbers.
	fmt.Println("")
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		imei := start

		// Randomly generate the remaining digits to make the length 14.
		for len(imei) < 14 {
			imei += strconv.Itoa(rand.Intn(10))
		}

		// Calculate and append the check digit.
		imei += calcCheckDigit(imei)
		fmt.Println(imei)
	}

	fmt.Println("")
}
