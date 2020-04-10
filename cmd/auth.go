package cmd

import (
	"fmt"

	"github.com/fastly/go-fastly/fastly"
	"github.com/reymundbautista/fastver/getenv"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		verifyAuth()
	},
}

func init() {
	verifyCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func verifyAuth() {
	var envName = "FASTLY_API_TOKEN"

	if getenv.Exists(envName) {
		token := getenv.ReadEnvVar(envName)
		client, err := fastly.NewClient(token)
		if err != nil {
			fmt.Println("Client creation failed.")
			return
		}
		// Get information about the current user
		user, err := client.GetCurrentUser()
		if err != nil {
			fmt.Println("Authorization failed. Please verify your Fastly API Key.")
			return
		}

		fmt.Println("Authorization succeeded for user: " + user.Login)
	} else {
		fmt.Println("FASTLY_API_TOKEN environment variable must be set!")
	}

}
