package schedulecli

import (
 "fmt"
 "os"
 "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:  "schedulecli",
    Short: "schedulecli - this will help you find the next cron run for a cron string",
    Long: `schedulecli - this will help you find the next cron run for a cron string`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error with the CLI '%s'", err)
        os.Exit(1)
    }
}
