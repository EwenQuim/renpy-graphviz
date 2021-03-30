package parser

import (
	"fmt"
	"time"
)

// Runningtime computes running time
func RunningTime(s string) (string, time.Time) {
	fmt.Println(s)
	return s, time.Now()
}

// Track is this
func Track(s string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println("  - ", endTime.Sub(startTime))
}
