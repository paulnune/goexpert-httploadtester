package cmd

import (
	"fmt"
	"https://github.com/paulnune/goexpert-httploadtester/internal/stresstest"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var st = &stresstest.StressTest{}
var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Stress test tool",
	Long:  `Stress test for make HTTP requests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if st.Headers != nil && !stresstest.ValidateHeaders(st.Headers) {
			return fmt.Errorf("invalid headers format")
		}

		report, err := st.Run()
		if err != nil {
			return err
		}

		report.Print(os.Stdout)

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&st.Url, "url", "u", "", "Target URL")
	rootCmd.Flags().IntVarP(&st.Requests, "requests", "r", 10, "Number of requests")
	rootCmd.Flags().IntVarP(&st.Concurrency, "concurrency", "c", 100, "Number of concurrent requests")
	rootCmd.Flags().StringVarP(&st.Method, "method", "m", "GET", "Request method")
	rootCmd.Flags().StringSliceVarP(&st.Headers, "header", "H", []string{}, "Header in format of NAME:VALUE")
	rootCmd.Flags().DurationVarP(&st.Timeout, "timeout", "t", time.Second*5, "Each request timeout, example: 5s")
	rootCmd.Flags().StringVarP(&st.BodyEncoded, "body", "b", "", "Request body in base64 format")
}
