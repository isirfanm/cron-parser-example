package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	s := "* * * * *"
	sch, err := cron.ParseStandard(s)
	if err != nil {
		fmt.Println("Error parsing: ", err)
		os.Exit(1)
	}
	n := time.Now()
	fmt.Println("Now: ", n)
	t := sch.Next(n)
	fmt.Println("Next: ", t)
}
