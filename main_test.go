package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	result := hello("Srikanth")
	
	if (result != "Hello Srikanth") {
		t.Errorf("hello failed, expected %v, got  %v", "Hello Srikanth", result)
	} else {
		t.Logf("hello success, expected %v, got %v", "Hello Srikanth", result)
	}
	
	
	result1 := hello2("Srikanth")
	
	if (result1 != "Sorry Srikanth") {
		t.Errorf("hello2 failed, expected %v, got %v", "Hello Srikanth", result1)
	} else {
		t.Logf("hello2 success, expected %v, got %v", "Sorry Srikanth", result1)
	}
	
	
}