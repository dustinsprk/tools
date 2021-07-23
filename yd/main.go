package main

import (
	"fmt"
	"time"
)

func main() {
	yd := time.Now().UTC().YearDay()
	fmt.Print(yd)
}
