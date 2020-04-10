package cmd

import (
	"fmt"
	"log"
	"os"

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
		verifyAuth(apiTokenEnvName)
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

func verifyAuth(envName string) {
	if getenv.Exists(envName) {
		client := newFastlyClient(envName)
		user, _ := getCurrentUser(client)
		fmt.Println("Authorization succeeded for user: " + user.Login)
	} else {
		fmt.Println("FASTLY_API_TOKEN environment variable must be set!")
	}
}

func newFastlyClient(envVarName string) *fastly.Client {
	token := getenv.ReadEnvVar(envVarName)
	client, err := fastly.NewClient(token)
	if err != nil {
		log.Fatalf("Fastly client creation failed: %v", err)
		os.Exit(1)
	}
	return client
}

// Interface that has the same GetCurrentUser() signature from fastly.Client
type fastlyClient interface {
	GetCurrentUser() (*fastly.User, error)
}

// Accepts the client parameter using the fastlyClient interface type
func getCurrentUser(client fastlyClient) (*fastly.User, error) {
	user, err := client.GetCurrentUser()
	if err != nil {
		log.Fatalf("Authorization failed. Please verify your Fastly API Key.: %v", err)
		os.Exit(1)
	}

	return user, err
}
