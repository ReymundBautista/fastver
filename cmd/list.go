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
	Short: "List Fastly services",
	Long: `Lists all Fastly services along with their ID that is needed for
	other commands`,
	Run: func(cmd *cobra.Command, args []string) {
		if getenv.Exists(apiTokenEnvName) {
			client := newFastlyClient(apiTokenEnvName)
			services, _ := listServices(client)
			for _, service := range services {
				fmt.Println(service.Name + " ID: " + service.ID)
			}

		} else {
			fmt.Println("FASTLY_API_TOKEN environment variable must be set!")
		}
	},
}

func init() {
	servicesCmd.AddCommand(listCmd)
}

// Interface that has the same ListServices() signature from fastly.Client
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
