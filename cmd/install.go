package main

import (
	"fmt"
	"os"

	"github.com/girikuncoro/godc/pkg/cli/configfile"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type installCmd struct {
	// path to config file to install cmd
	configFile string
	configDest string
	configDir  string
}

func registerInstallCmds(c *Cli) {
	c.installCmd = &installCmd{}

	installCmd := &cobra.Command{
		Use:     "install",
		Short:   "Install godc cli",
		Example: `godc install --file config.yml`,
		PreRunE: c.preRunner(installPre),
		RunE:    c.runner(installRun),
	}

	c.rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringVarP(&c.installCmd.configFile, "file", "f", "", "path to YAML configuration file for the cluster")
	installCmd.Flags().StringVarP(&c.installCmd.configDest, "destination", "d", "~/.godc", "Destination of CLI configuration")
}

func installPre(c *Cli) error {
	if c.installCmd.configFile == "" {
		return fmt.Errorf("Install config file is not provided")
	}

	configDir, err := homedir.Expand(c.installCmd.configDest)
	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			return fmt.Errorf("Error creating config destination directory")
		}
	}
	c.installCmd.configDir = configDir

	return nil
}

func installRun(c *Cli) error {
	datacenterConfig, err := configfile.ParseDatacenterConfig(c.installCmd.configFile)

	// TODO(giri): validate spec
	err = configfile.WriteDatacenterConfig(c.installCmd.configDir, datacenterConfig)
	return err
}
