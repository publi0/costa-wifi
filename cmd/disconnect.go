package cmd

import (
	"costa-wifi/internal/service"
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "List and disconnect a WiFi session",
	Run: func(cmd *cobra.Command, args []string) {
		sessions, err := service.GetPlanSessions()
		if err != nil {
			pterm.Error.Printf("Failed to get sessions: %v\n", err)
			return
		}

		userSessions, err := service.GetUserSession(sessions)
		if err != nil {
			pterm.Error.Printf("Failed to get user sessions: %v\n", err)
			return
		}

		if len(userSessions) == 0 {
			pterm.Info.Println("No active sessions found.")
			return
		}

		options := make([]string, len(userSessions))
		for i, session := range userSessions {
			options[i] = fmt.Sprintf("Session ID: %s - Start Time: %s - IP: %s", session.SessionID, session.StartTime, session.IPAddress)
		}

		selectedOption, err := pterm.DefaultInteractiveSelect.
			WithOptions(options).
			WithDefaultText("Select a session to disconnect:").
			Show()

		if err != nil {
			pterm.Error.Printf("Failed to select session: %v\n", err)
			return
		}

		var selectedSessionID string
		for i, option := range options {
			if option == selectedOption {
				selectedSessionID = userSessions[i].SessionID
				break
			}
		}

		spinner, _ := pterm.DefaultSpinner.Start("Disconnecting session...")

		err = service.DisconnectSession(selectedSessionID)
		if err != nil {
			spinner.Fail(fmt.Sprintf("Failed to disconnect session: %v", err))
			return
		}

		spinner.Success(fmt.Sprintf("Successfully disconnected session: %s", selectedSessionID))
	},
}

func init() {
	rootCmd.AddCommand(disconnectCmd)
}
