//CHASE
//Concurrent Harvester for Adversary Studies & Examination

package main

import (
	"fmt"
)

func main() {

	var status int
	var report Report
	report.url = "https://bergcybersec.club"
	status = ping(report.url)
	if status == 200 {
		fmt.Println("Valid Report")
	}

}
