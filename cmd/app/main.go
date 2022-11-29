package main

import (
	"fmt"
	"time"

	"github.com/dahaev/golan_mod_test/internal/app/connect"
)

func main() {
	start := time.Now()
	connect.AnoData()
	s := connect.Data

	for _, value := range s {
		fmt.Println(value.Odata_nextlink)
	}

	end := time.Now()
	fmt.Printf("time: %v", end.Sub(start))

}
