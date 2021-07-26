package main

// qq

import (
	"fmt"
	"math"
	"time"
)

func main() {
	now := time.Now()
	then := time.Date(int(2023), time.April, int(30), int(16), int(0), int(0), int(0), time.Local)
	r := then.Sub(now).Hours() / float64(24)
	rem := math.Ceil(r)
	fmt.Println(int(rem))
}
