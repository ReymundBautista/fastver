package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fastly/go-fastly/fastly"
	"github.com/reymundbautista/fastver/getenv"
	"github.com/spf13/cobra"
)

// activeCmd represents the active command
var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "Display the active version for a service",
	Run: func(cmd *cobra.Command, args []string) {
		if getenv.Exists(apiTokenEnvName) {
			client := newFastlyClient(apiTokenEnvName)
			service, err := getService(client)
			if err != nil {
				log.Fatalf("Listing services failed: %v", err)
				os.Exit(1)
			}

			versions := service.Versions
			active := getActiveVersion(versions)

			if active != 0 {
				fmt.Printf("%v active version is %d \n", service.Name, active)
			} else {
				fmt.Printf("No active version for service: %v", service.Name)
			}

		} else {
			fmt.Println("FASTLY_API_TOKEN environment variable must be set!")
		}
	},
}

func init() {
	servicesCmd.AddCommand(activeCmd)
	activeCmd.Flags().StringVarP(&id, "id", "i", "", "Fastly service ID")
	activeCmd.MarkFlagRequired("id")
}

func getActiveVersion(versions []*fastly.Version) int {
	for _, v := range versions {
		if v.Active {
			return v.Number
		}
	}
	return 0
}
