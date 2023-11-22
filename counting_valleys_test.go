package main

import "testing"

func countingValleys(arr []string) int {
	var (
		countUp     = 0
		countDown   = 0
		countValley = 0
	)

	for _, val := range arr {
		if val == "U" {
			countUp++

			if (countUp - countDown) == 0 {
				countValley++
			}
		} else {
			countDown++
		}
	}

	return countValley
}

func TestCountingValleys(t *testing.T) {
	valleys := countingValleys([]string{"U", "D", "D", "D", "U", "D", "U", "U"})
	expected := 1
	if valleys != expected {
		t.Errorf("Got unexpected total of valley: %v want %v", valleys, expected)
	}
}

func TestCountingValleys2(t *testing.T) {
	valleys := countingValleys([]string{"D", "D", "U", "U", "D", "D", "U", "U"})
	expected := 2
	if valleys != expected {
		t.Errorf("Got unexpected total of valley: %v want %v", valleys, expected)
	}
}
