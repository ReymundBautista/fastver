package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fastly/go-fastly/fastly"
	"github.com/reymundbautista/fastver/getenv"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if getenv.Exists(apiTokenEnvName) {
			client := newFastlyClient(apiTokenEnvName)
			services, _ := listServices(client)
			for _, service := range services {
				fmt.Println(service.Name)
			}

		} else {
			fmt.Println("FASTLY_API_TOKEN environment variable must be set!")
		}
	},
}

func init() {
	servicesCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Interface that has the same GetCurrentUser() signature from fastly.Client
type fastlyClientSerivces interface {
	ListServices(i *fastly.ListServicesInput) ([]*fastly.Service, error)
}

// Accepts the client parameter using the fastlyClient interface type
func listServices(client fastlyClientSerivces) ([]*fastly.Service, error) {
	var i *fastly.ListServicesInput
	services, err := client.ListServices(i)
	if err != nil {
		log.Fatalf("Listing services failed: %v", err)
		os.Exit(1)
	}

	return services, err
}
