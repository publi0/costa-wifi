package cmd

import (
	"costa-wifi/internal/service"
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "Get all active sessions",
	Run: func(cmd *cobra.Command, args []string) {
		sessions, err := service.GetPlanSessions()
		if err != nil {
			pterm.Error.Printf("Failed to get sessions: %v\n", err)
			return
		}

		usrSessions, err := service.GetUserSession(sessions)
		if err != nil {
			pterm.Error.Printf("Failed to get sessions: %v\n", err)
			return
		}

		if len(usrSessions) == 0 {
			pterm.Info.Println("No active sessions found.")
			return
		}

		tableData := [][]string{
			{"Session ID", "Start Time", "IP Address", "MAC Address", "Status"},
		}

		for _, session := range usrSessions {
			startTime, _ := time.Parse(time.RFC3339, session.StartTime)
			formattedTime := startTime.Format("2006-01-02 15:04:05")
			
			device := session.UserAgent
			if len(device) > 20 {
				device = device[:17] + "..."
			}

			tableData = append(tableData, []string{
				session.SessionID,
				formattedTime,
				session.IPAddress,
				session.MacAddress,
				session.Status,
			})
		}

		pterm.DefaultHeader.WithFullWidth().Printfln("Active WiFi Sessions")
		pterm.Println()

		err = pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
		if err != nil {
			pterm.Error.Printf("Failed to render table: %v\n", err)
			return
		}

		pterm.Println() 
		pterm.Info.Printf("Total active sessions: %d\n", len(usrSessions))
	},
}

func init() {
	rootCmd.AddCommand(sessionsCmd)
}
