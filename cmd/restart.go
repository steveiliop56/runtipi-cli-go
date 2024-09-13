package cmd

import (
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/commands"
)

func init() {
	restartCmd.Flags().BoolVar(&noPermissions, "no-permissions", false, "Skip setting permissions.")
	restartCmd.Flags().StringVar(&envFile, "env-file", "", "Path to custom .env file, it has to be absolute")
	rootCmd.AddCommand(restartCmd)
}

var restartCmd = &cobra.Command{
	Use: "restart",
	Short: "Restart Runtipi",
	Long: "Use this command to restart the Runtipi docker stack",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Start(envFile, noPermissions)
	},
}
