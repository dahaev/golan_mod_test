package main

import (
	"fmt"
	"time"

	"github.com/dahaev/golan_mod_test/internal/app/connect"
)

func main() {
	start := time.Now()
	connect.AnoData()
	end := time.Now()
	fmt.Printf("time: %v", end.Sub(start))

}
