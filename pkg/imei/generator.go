// imei.go
package imei

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// DeviceMetadata contains information about a specific phone model.
type DeviceMetadata struct {
	Brand  string
	Model  string
	Color  string
	Memory string
	CPU    string
	IMEI   string
}

// List of available devices with their metadata.
var devices = []DeviceMetadata{
	{Brand: "Apple", Model: "iPhone 13", Color: "Black", Memory: "128GB", CPU: "A15 Bionic"},
	{Brand: "Apple", Model: "iPhone 14", Color: "Blue", Memory: "256GB", CPU: "A16 Bionic"},
	{Brand: "Samsung", Model: "Galaxy S22", Color: "White", Memory: "128GB", CPU: "Exynos 2200"},
	{Brand: "Samsung", Model: "Galaxy Z Flip", Color: "Purple", Memory: "256GB", CPU: "Snapdragon 888"},
	{Brand: "Xiaomi", Model: "Mi 11", Color: "Gray", Memory: "128GB", CPU: "Snapdragon 888"},
	{Brand: "OnePlus", Model: "9 Pro", Color: "Green", Memory: "256GB", CPU: "Snapdragon 888"},
	{Brand: "Google", Model: "Pixel 6", Color: "Black", Memory: "128GB", CPU: "Google Tensor"},
	{Brand: "Google", Model: "Pixel 7", Color: "White", Memory: "256GB", CPU: "Google Tensor G2"},
}

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

// GenerateIMEI generates a random IMEI and returns the device metadata associated with it.
func GenerateIMEI(prefix string) (DeviceMetadata, error) {
	if len(prefix) < 8 || len(prefix) > 12 {
		return DeviceMetadata{}, ErrInvalidPrefix
	}

	imei := prefix
	rand.Seed(time.Now().UnixNano())

	// Randomly generate the remaining digits to make the length 14.
	for len(imei) < 14 {
		imei += strconv.Itoa(rand.Intn(10))
	}

	// Calculate and append the check digit.
	imei += calcCheckDigit(imei)

	// Randomly select a device from the list and assign the IMEI to it.
	device := devices[rand.Intn(len(devices))]
	device.IMEI = imei

	return device, nil
}

// GenerateIMEIs generates a list of IMEIs and their associated device metadata.
func GenerateIMEIs(prefix string, count int) ([]DeviceMetadata, error) {
	if count <= 0 {
		return nil, ErrInvalidCount
	}

	var imeis []DeviceMetadata
	for i := 0; i < count; i++ {
		device, err := GenerateIMEI(prefix)
		if err != nil {
			return nil, err
		}
		imeis = append(imeis, device)
	}

	return imeis, nil
}

// Custom error messages
var (
	ErrInvalidPrefix = errors.New("invalid prefix: must be between 8 and 12 digits")
	ErrInvalidCount  = errors.New("invalid count: must be greater than zero")
)
