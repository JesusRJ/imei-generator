// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gitlab.com/imei-br/imei-generator/pkg/imei"
)

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

	// Generate and print IMEI numbers using the imei package.
	devices, err := imei.GenerateIMEIs(start, count)
	if err != nil {
		fmt.Println("Error generating IMEIs:", err)
		return
	}

	// Print the generated IMEIs with their metadata.
	for _, device := range devices {
		fmt.Printf("IMEI: %s\nBrand: %s\nModel: %s\nColor: %s\nMemory: %s\nCPU: %s\n\n",
			device.IMEI, device.Brand, device.Model, device.Color, device.Memory, device.CPU)
	}
}
