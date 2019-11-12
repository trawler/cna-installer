package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/assets"
)

// createCmd represents the create command
var assetCmd = &cobra.Command{
	Use:   "assets",
	Short: "Manage Assets",
	Long:  `Deploy or Dessroy Kubernetes addons on to the cluster`,
}

var assetDeployCmd = &cobra.Command{
	Use:   "create",
	Short: "Deploy workload assets.",
	Run: func(cmd *cobra.Command, args []string) {
		assetCreate()
	},
}

var assetDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Undeploy workload assets.",
	Run: func(cmd *cobra.Command, args []string) {
		assetDestroy()
	},
}

func assetCreate() {
	k8sClient, err := assets.NewClient(logDir)
	if err != nil {
		fmt.Printf("\nERROR:\nfailed to initialize Kubernetes client:\n")
		log.Fatal(err)
	}

	err = assets.Install(k8sClient)
	if err != nil {
		fmt.Printf("\nERROR:\nfailed to install Kubernetes assets:\n")
		log.Fatal(err)
	}
}

func assetDestroy() {
	fmt.Println("destroy deployment")
}

func init() {
	rootCmd.AddCommand(assetCmd)
	assetCmd.AddCommand(assetDeployCmd)
	assetCmd.AddCommand(assetDestroyCmd)
}
