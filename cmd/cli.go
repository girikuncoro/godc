package main

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cli struct for GoDC Cli
type Cli struct {
	// current command
	Cmd *cobra.Command
	// arguments for current command
	Args []string

	v *viper.Viper
	// command hierarchy
	rootCmd   *cobra.Command
	vmCmd     *vmCmd
	volumeCmd *volumeCmd

	getCmd     *getCmd
	installCmd *installCmd
}

const (
	// CliProgram CLI program name
	CliProgram = "godc"
	// configName is name of the config file
	configName = "config"
	// keys from config file
	configKeyHosts      = "hosts"
	configKeyDNSServer  = "dnsServer"
	configKeyDHCPServer = "dnsServer"
)

// NewCli configures new CLI for GoDC.
// Loads environment configuraiton and registers sub commands.
func NewCli() *Cli {
	cli := &Cli{
		v: viper.New(),
	}

	cli.rootCmd = &cobra.Command{
		Use:   CliProgram,
		Short: "GOJEK Datacenter CLI",
		RunE:  cli.usageRunner(),
	}

	cli.readConfig()

	registerGetCmds(cli)
	registerInstallCmds(cli)

	return cli
}

// Run runs the CLI
func (c *Cli) Run() {
	c.v.ReadInConfig()
	if err := c.rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// usage print usage of current command
func (c *Cli) usage() error {
	if c.Cmd != nil {
		return c.Cmd.Usage()
	}
	return c.rootCmd.Usage()
}

// runner execute current command with args
func (c *Cli) runner(runner func(*Cli) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c.rootCmd.SilenceUsage = true
		c.Cmd = cmd
		c.Args = args
		return runner(c)
	}
}

// preRunner run a preRunner function for current command and arguments
func (c *Cli) preRunner(preRunner func(*Cli) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c.Cmd = cmd
		c.Args = args
		return preRunner(c)
	}
}

// usageRunner call usage on current command
func (c *Cli) usageRunner() func(*cobra.Command, []string) error {
	return c.runner(func(c *Cli) error {
		return c.usage()
	})
}

// readConfig reads godc config from file located in home directory
func (c *Cli) readConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c.v.AddConfigPath(path.Join(home, ".godc"))
	c.v.SetConfigName(configName)
	err = c.v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
