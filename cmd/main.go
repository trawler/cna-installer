package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
)

var cfgFile string
var cluster *terraform.Cluster
var err error
var logDir string

var stateFileName string
var tf *terraform.Executor

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cna-installer",
	Short: "Creates a CNA cluster",
	Long: `cna-installer is a binary that installs, sets-up and configures a
Kubernetes cluster with the CNA stack applications.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cna-installer.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// If no config file was provided, try to use the default config file.
	if cfgFile == "" {
		home, err := getHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		cfgFile = path.Join(home, ".cna-installer.yaml")
	}

	// Parse the config file into a Cluster struct
	cluster, err = terraform.ParseConfigFile(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
}

func getStateFilePath(tfName string) (string, error) {
	dir, err := getLogDir()
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	statefileName := filepath.Join(dir, fmt.Sprintf("%s_terraform.tfstate", tfName))
	return statefileName, nil
}

func getLogDir() (string, error) {
	logDir, err := filepath.Abs("../logs")
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return logDir, nil
}

func getHomeDir() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("cannot get home dir. is $HOME set in your environment? \n%v", err)
	}
	return home, nil
}

// Agent Pool Name must start with a lowercase letter, have max length of 12, and only have characters a-z0-9.
// This function takes the cluster's name and formats a valid agent pool name.
func sanitizeAgentPoolName(cName string, cOwner string) (string, error) {
	rawName := fmt.Sprintf("%s-%s", cOwner, cName)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	splitArray := reg.Split(rawName, -1)
	maxLength := 12
	tmpName := ""
	addon := ""

	for _, element := range splitArray {
		// don't capitalize first letter
		addon = strings.ToLower(element)

		// limit output to maxLength
		if len(tmpName+string(element)) <= maxLength {
			tmpName = tmpName + addon
		} else {
			break
		}

	}
	return tmpName, nil
}
