package main

import "testing"

// From the details, we can assume after 100 trips
// the turned-on switches are multiple of 100
// since it mentioned that the second trip is multiplying two,
// the third is multiplying of three, and so on
func turnedOnLamps(tripNumber int) int {
	onLamp := 0
	repetition := 100
	if tripNumber > repetition/2 {
		return 1
	}
	for i := tripNumber; i <= repetition; i++ {
		if i%tripNumber == 0 {
			onLamp++
		}
	}

	return onLamp
}

func TestTurnedOnLamps(t *testing.T) {
	turnedOn := turnedOnLamps(100)
	expected := 1
	if turnedOn != expected {
		t.Errorf("Got unexpected total of turned on lamp: %v want %v", turnedOn, expected)
	}
}

func TestTurnedOnLamps2(t *testing.T) {
	turnedOn := turnedOnLamps(4)
	expected := 25
	if turnedOn != expected {
		t.Errorf("Got unexpected total of turned on lamp: %v want %v", turnedOn, expected)
	}
}
