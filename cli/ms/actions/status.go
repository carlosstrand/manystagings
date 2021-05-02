package actions

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func coloredStatus(status string) string {
	switch status {
	case "RUNNING":
		return color.GreenString(status)
	case "NOT RUNNING":
		return color.YellowString("PAUSED")
	case "PENDING", "UNKNOWN":
		fmt.Println(color.YellowString(status))
		return color.YellowString(status)
	case "FAILED":
		return color.RedString(status)
	}
	return status
}

func (a *Actions) Status() error {
	statuses, err := a.client.GetEnvironmentStatus(context.TODO(), a.config.EnvironmentID)
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Application", "Status", "Age", "Port", "Public URL"})
	for _, s := range statuses {
		age := ""
		if s.Application.StartedAt != nil {
			age = time.Since(*s.Application.StartedAt).Truncate(time.Millisecond).Round(time.Second).String()
		}
		table.Append([]string{s.Application.Name, coloredStatus(s.Status), age, fmt.Sprintf("%d", s.Application.Port), s.Application.PublicUrl})
	}
	table.Render()
	return nil
}
