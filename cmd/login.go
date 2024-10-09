package cmd

import (
	"costa-wifi/internal/config"
	"costa-wifi/internal/service"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with your Costa card info",
	Run: func(cmd *cobra.Command, args []string) {
		cardID, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Costa card ID")
		if err != nil {
			pterm.Error.Println("Failed to get Costa card ID:", err)
			return
		}

		birthday, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your birthday (YYYY-MM-DD)")
		if err != nil {
			pterm.Error.Println("Failed to get birthday:", err)
			return
		}

		_, err = service.Login(
			cardID,
			birthday,
		)
		if err != nil {
			pterm.Error.Println("Login attempt failed!")
			panic(err)
		}
		pterm.Success.Printf("Login attempt with Card ID: %s and Birthday: %s\n", cardID, birthday)

		config.WriteConfigValue(config.KeyCard, cardID)
		config.WriteConfigValue(config.KeyBithday, birthday)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
