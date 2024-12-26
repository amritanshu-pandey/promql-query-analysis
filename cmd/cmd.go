package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"promql-query-analysis/dashboardQueries"
)

var rootCmd = &cobra.Command{
	Use:   "promql-query-analysis",
	Short: "promql-query-analysis tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: promql-query-analysis [command]")
	},
}

var grafanaDashboard = &cobra.Command{
	Use:   "grafana-dashboard",
	Short: "Extract queries from a Grafana dashboard",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		grafanaURL, _ := cmd.Flags().GetString("grafanaURL")
		grafanaST, _ := cmd.Flags().GetString("grafanaServiceToken")

		gc, err := dashboardQueries.CreateGrafanaClient(grafanaURL, grafanaST, false, true)
		if err != nil {
			log.Fatal(err)
		}
		dashboard, err2 := dashboardQueries.NewDashboardQueries(gc, args[0])
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println(dashboard)
	},
}

func Execute() {
	grafanaDashboard.Flags().StringP("grafanaURL", "u", "http://localhost:3000", "Grafana URL")
	grafanaDashboard.Flags().StringP("grafanaServiceToken", "t", "", "Grafana service token")
	rootCmd.AddCommand(grafanaDashboard)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
