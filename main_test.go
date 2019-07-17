package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func TestMain(t *Testing.T) {
	result := hello(Srikanth)
	
	if result != "Hello Srikanth" {
		t.Error("hello failed, expected %v, %v, "Hello Srikanth", result)
	} else {
		t.Logf(("hello success, expected %v, %v, "Hello Srikanth", result)
	}
	
	
	result1 := hello2(Srikanth)
	
	if result1 != "Hello Srikanth" {
		t.Error("hello2 failed, expected %v, actual %v, "Hello Srikanth", result1)
	} else {
		t.Logf(("hello2 success, expected %v, %v, "Hello Srikanth", result)
	}
	
	
}