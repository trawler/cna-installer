package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/asset"
)

// createCmd represents the create command
var assetCmd = &cobra.Command{
	Use:   "asset",
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
	fmt.Println("create deployment")
	k8sClient, err := asset.Client(logDir)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("k8sClient: %+v", k8sClient)
}

func assetDestroy() {
	fmt.Println("destroy deployment")
}

func init() {
	rootCmd.AddCommand(assetCmd)
	assetCmd.AddCommand(assetDeployCmd)
	assetCmd.AddCommand(assetDestroyCmd)
}
