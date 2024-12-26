package main

import (
	"fmt"
	"os"
	"promql-query-analysis/dashboardQueries"
)

func main() {
	grafana, err := dashboardQueries.CreateGrafanaClient("https://grafana.xps.lan", "123", false, true)
	if err != nil {
		fmt.Println("Unable to create Grafana Client")
		fmt.Println(err)
		os.Exit(1)
	}

	dashboard, err := grafana.Dashboard("rYdddlPWk")
	fmt.Println(dashboard, err)
}
