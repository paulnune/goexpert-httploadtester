package stresstest

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"time"
)

type (
	Report struct {
		TimeSpent          time.Duration
		TotalRequests      int
		SuccessfulRequests int
		FailedRequests     map[int]int
	}
)

func (report *Report) Print(w io.Writer) {
	t := table.NewWriter()
	t.SetOutputMirror(w)

	columns := table.Row{"Total Requests", "Time Spent", "Successful"}
	row := table.Row{report.TotalRequests, report.TimeSpent, report.SuccessfulRequests}
	for k, v := range report.FailedRequests {
		columns = append(columns, fmt.Sprintf("HTTP %d", k))
		row = append(row, v)
	}
	t.AppendHeader(columns)

	t.AppendRows([]table.Row{
		row,
	})
	t.Render()
}
