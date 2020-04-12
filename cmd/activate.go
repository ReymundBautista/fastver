package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fastly/go-fastly/fastly"
	aurora "github.com/logrusorgru/aurora"
	"github.com/reymundbautista/fastver/getenv"
	"github.com/spf13/cobra"
)

var version int

// activateCmd represents the activate command
var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate a version for a Fastly service",
	Run: func(cmd *cobra.Command, args []string) {
		if getenv.Exists(apiTokenEnvName) {
			client := newFastlyClient(apiTokenEnvName)
			service, err := getService(client)
			if err != nil {
				log.Fatalf("Failed to bind to service with error reported: %v", err)
			}

			if !verifyVersion(version, service) {
				log.Fatalf("Version %d not found for service %v", version, aurora.Bold(service.Name))
			}

			invalidEntry := true
			response := "n"

			for invalidEntry {
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("Are you sure you want to activate version %d for service %v?\n", version, service.Name)
				fmt.Print("Enter y or n:  ")
				response, _ = reader.ReadString('\n')

				switch strings.TrimSuffix(response, "\n") {
				case "y", "n":
					invalidEntry = false
				default:
					fmt.Printf("%v is not a valid entry. Please enter y or n\n", response)
				}
			}

			if strings.TrimSuffix(response, "\n") == "y" {
				i := &fastly.ActivateVersionInput{
					Service: id,
					Version: version,
				}
				client.ActivateVersion(i)
				fmt.Printf("Activated version %d for service %v\n", version, service.Name)
			}
		} else {
			fmt.Println("FASTLY_API_TOKEN environment variable must be set!")
		}
	},
}

func init() {
	servicesCmd.AddCommand(activateCmd)
	activateCmd.Flags().StringVarP(&id, "id", "i", "", "Fastly service ID")
	activateCmd.MarkFlagRequired("id")
	activateCmd.Flags().IntVarP(&version, "version", "v", -1, "Limit number of previous versions (default = 10")
	activateCmd.MarkFlagRequired("version")
}

func verifyVersion(version int, service *fastly.Service) bool {
	for _, v := range service.Versions {
		if version == v.Number {
			return true
		}
	}
	return false
}
