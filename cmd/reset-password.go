package cmd

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/env"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
)

func init() {
	rootCmd.AddCommand(resetPasswordCmd)
}

var resetPasswordCmd = &cobra.Command{
	Use: "reset-password",
	Short: "Reset Runtipi's Password",
	Long: "Use this command to reset your Runtipi's password if you forget it",
	Run: func(cmd *cobra.Command, args []string) {
		spinner.SetMessage("Creating reset password request")
		spinner.Start()

		rootFolder, osErr := os.Getwd()
	
		if osErr != nil {
			spinner.Fail("Failed to get root folder")
			spinner.Stop()
			fmt.Printf("Error: %s\n", osErr)
			return;
		}

		time := time.Now().Unix()
		writeErr := os.WriteFile(path.Join(rootFolder, "state", "password-change-request"), []byte(string(time)), 0644)

		if writeErr != nil {
			spinner.Fail("Failed to create password change request")
			spinner.Stop()
			fmt.Printf("Error: %s\n", writeErr)
			return;
		}

		internalIp, _ := env.GetEnvValue("INTERNAL_IP")
		nginxPort, _ := env.GetEnvValue("NGINX_PORT")

		message := "Password request created. " + "Head back to http://" + internalIp + ":" + nginxPort + "/reset-password to reset your password."

		spinner.Succeed(message)
		spinner.Stop()
	},
}