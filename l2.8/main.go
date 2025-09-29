package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting NTP time: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ntpTime.Format("2006-01-02 15:04:05 -07:00"))
}
