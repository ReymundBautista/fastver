package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fastly/go-fastly/fastly"
	aurora "github.com/logrusorgru/aurora"
	"github.com/reymundbautista/fastver/getenv"
	"github.com/spf13/cobra"
)

var id string
var limit int

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get service versions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if getenv.Exists(apiTokenEnvName) {
			client := newFastlyClient(apiTokenEnvName)
			service, err := getService(client)
			if err != nil {
				log.Fatalf("Listing services failed: %v", err)
				os.Exit(1)
			}

			versions := service.Versions

			// Reverse the array
			for i, j := 0, len(versions)-1; i < j; i, j = i+1, j-1 {
				versions[i], versions[j] = versions[j], versions[i]
			}

			versions = filterVersions(versions, limit)

			for _, v := range versions {
				fmt.Println("Version: " + strconv.Itoa(v.Number))
				if v.Active {
					fmt.Println(aurora.Sprintf("%s", aurora.Green("Active: "+strconv.FormatBool(v.Active))))
				} else {
					fmt.Println("Active: " + strconv.FormatBool(v.Active))
				}
				fmt.Println("")
			}

		} else {
			fmt.Println("FASTLY_API_TOKEN environment variable must be set!")
		}
	},
}

func init() {
	servicesCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&id, "id", "i", "", "Fastly service ID")
	getCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Number of previous versions (default = 10")
	getCmd.MarkFlagRequired("id")
}

// Interface that has the same GetServices() signature from fastly.Client
type fastlyClientGet interface {
	GetService(i *fastly.GetServiceInput) (*fastly.Service, error)
}

func getService(client fastlyClientGet) (*fastly.Service, error) {
	var input = fastly.GetServiceInput{ID: id}
	service, err := client.GetService(&input)
	if err != nil {
		log.Fatalf("Getting service failed: %v", err)
		os.Exit(1)
	}

	return service, err
}

func filterVersions(versions []*fastly.Version, limit int) []*fastly.Version {
	var filteredVersions []*fastly.Version

	for index, v := range versions {
		if index >= limit {
			break
		}
		filteredVersions = append(filteredVersions, v)
	}

	return filteredVersions
}
