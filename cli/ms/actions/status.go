package actions

import (
	"context"
	"fmt"
	"os"

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
	table.SetHeader([]string{"Application", "Status"})
	for _, s := range statuses {
		table.Append([]string{s.Application.Name, coloredStatus(s.Status)})
	}
	table.Render()
	return nil
}
