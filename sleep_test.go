package main

import (
    "fmt"
    "testing"
    "time"
    )
func TestSleepFunction(t *testing.T) {
    var timeTaken int
    start := time.Now()
    sleep(3)
    elapsed := time.Since(start).Seconds()
    timeTaken = int(elapsed)
    fmt.Printf("Operation took ", timeTaken, " Seconds")
    if timeTaken != 3{
        t.Error("Expected Sleep time is 3 but actual Sleep time is ",timeTaken)    
    }
}