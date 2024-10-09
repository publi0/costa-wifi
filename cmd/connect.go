package cmd

import (
	"costa-wifi/internal/service"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var ipAddress string

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to WiFi",
	Run: func(cmd *cobra.Command, args []string) {
		sessions, err := service.GetPlanSessions()
		if err != nil {
			pterm.Error.Printf("Failed to get sessions: %v\n", err)
			return
		}

		if len(sessions.Data) == 0 || len(sessions.Data[0].InternetPackagesCategories) == 0 ||
			len(sessions.Data[0].InternetPackagesCategories[0].InternetPackages) == 0 {
			pterm.Error.Println("No available internet packages found.")
			return
		}

		bookingID := sessions.Data[0].InternetPackagesCategories[0].InternetPackages[0].PackageDetails.BookingID

		if ipAddress == "" {
			var err error
			ipAddress, err = pterm.DefaultInteractiveTextInput.
				WithMultiLine(false).
				WithDefaultText("Enter your IP address:").
				Show()

			if err != nil {
				pterm.Error.Printf("Failed to get IP address: %v\n", err)
				return
			}
		}

		spinner, _ := pterm.DefaultSpinner.Start("Connecting to WiFi...")

		connectResp, err := service.ConnectSession(bookingID, ipAddress)
		if err != nil {
			spinner.Fail(pterm.Sprintf("Failed to connect: %v", err))
			return
		}
		spinner.Success(pterm.Sprintf("Successfully connected. Session details:"))
		pterm.Info.Printf("Session ID: %s\n", connectResp.Data.SessionID)
		pterm.Info.Printf("Start Time: %s\n", connectResp.Data.StartTime)
		pterm.Info.Printf("IP Address: %s\n", connectResp.Data.IPAddress)
		pterm.Info.Printf("Mac Address: %s\n", connectResp.Data.MacAddress)
		pterm.Info.Printf("Status: %s\n", connectResp.Data.Status)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&ipAddress, "ip", "i", "", "IP address to connect with")
}
