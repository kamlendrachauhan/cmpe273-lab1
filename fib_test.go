package main

import "testing"

func TestFibonacci(t *testing.T) {
    var v int
    v = calculateFib(6)
    if v !=  8{
        t.Error("Expected 8 but got ", v)
    }
    
    v = calculateFib(-5)
    if v !=  -1{
        t.Error("Expected -1 but got ", v)
    }
}