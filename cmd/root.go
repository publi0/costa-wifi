package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "costa-wifi",
    Short: "Manage your Costa Cruise Wifi Connection",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("use: costa-wifi help")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
