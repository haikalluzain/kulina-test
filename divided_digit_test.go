package main

import (
	"bytes"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func DividedDigit(number int) {
	digits := len(strconv.Itoa(number)) - 1
	divider := int(math.Pow(10, float64(digits)))

	for divider > 0 {
		// Calculate the quotient and remainder
		quotient := number / divider
		remainder := number % divider
		// Print the result
		log.Println(quotient * divider)
		// Update the number for the next loop
		number = remainder
		// divider divided by 10 for the next loop
		divider /= 10
	}
}

func TestDividedDigit(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	DividedDigit(1345679)
	output := buf.String()
	expectedLines := []string{
		"1000000",
		"300000",
		"40000",
		"5000",
		"600",
		"70",
		"9",
	}
	actualLines := strings.Split(output, "\n")
	for i, expectedLine := range expectedLines {
		if i < len(actualLines) && !strings.Contains(actualLines[i], expectedLine) {
			t.Errorf("Line %d: Expected '%s' not found in '%s'", i+1, expectedLine, actualLines[i])
		}
	}
}

func TestDividedDigit2(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	DividedDigit(145698)
	output := buf.String()
	expectedLines := []string{
		"100000",
		"40000",
		"5000",
		"600",
		"90",
		"8",
	}
	actualLines := strings.Split(output, "\n")
	for i, expectedLine := range expectedLines {
		if i < len(actualLines) && !strings.Contains(actualLines[i], expectedLine) {
			t.Errorf("Line %d: Expected '%s' not found in '%s'", i+1, expectedLine, actualLines[i])
		}
	}
}
