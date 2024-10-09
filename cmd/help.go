package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var helperCmd = &cobra.Command{
	Use:   "help",
	Short: "Display help information for all commands",
	Run: func(cmd *cobra.Command, args []string) {
		pterm.DefaultHeader.WithFullWidth().Printfln("Costa WiFi CLI Helper")
		pterm.Println()

		commands := []*cobra.Command{
			loginCmd,
			connectCmd,
			disconnectCmd,
			sessionsCmd,
			versionCmd,
		}

		data := [][]string{
			{"Command", "Description", "Usage"},
		}

		for _, command := range commands {
			usage := command.Use
			if command.Flags().HasFlags() {
				usage += " [flags]"
			}
			data = append(data, []string{command.Name(), command.Short, usage})
		}

		err := pterm.DefaultTable.WithHasHeader().WithData(data).Render()
		if err != nil {
			pterm.Error.Printf("Failed to render table: %v\n", err)
			return
		}

		pterm.Println()
		pterm.Info.Println("For more details on each command, use: costa-wifi [command] --help")
	},
}

func init() {
	rootCmd.AddCommand(helperCmd)
}
