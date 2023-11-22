package main

import (
	"testing"
)

func sockMerchant(arr []int) int {
	total := 0
	pairs := map[int]bool{}
	for _, val := range arr {
		if ok := pairs[val]; ok {
			pairs[val] = false
			total++
			continue
		}
		pairs[val] = true
	}
	return total
}

func TestSockMerchant(t *testing.T) {
	matched := sockMerchant([]int{10, 20, 20, 10, 10, 30, 50, 10, 20})
	expected := 3
	if matched != expected {
		t.Errorf("Got unexpected total of matched socks: %v want %v", matched, expected)
	}
}

func TestSockMerchant2(t *testing.T) {
	matched := sockMerchant([]int{10, 20, 30, 40, 20})
	expected := 1
	if matched != expected {
		t.Errorf("Got unexpected total of matched socks: %v want %v", matched, expected)
	}
}
