//CHASE
//Concurrent Harvester for Adversary Studies & Examination

package chase

import (
	"chase/internal/harvest"
	"fmt"
)

func main() {

	var status int
	var report harvest.Report
	report.Url = "https://bergcybersec.club"
	status = harvest.Ping(report.Url)
	if status == 200 {
		fmt.Println("Valid Report")
	}

}
