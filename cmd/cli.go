package main

import (
	"os"

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

	// list of <IP>:<port> of libvirt host
	hostEndpoints []string
}

const (
	// CliProgram CLI program name
	CliProgram = "godc"
	// Environment Keys
	hostEndpointKey = "HOST_ENDPOINT"
)

// NewCli configures new CLI for GoDC.
// Loads environment configuraiton and registers sub commands.
func NewCli() *Cli {
	cli := &Cli{
		v: viper.New(),
	}

	cli.rootCmd = &cobra.Command{
		Use:   CliProgram,
		Short: "GO-JEK Datacenter CLI",
		RunE:  cli.usageRunner(),
	}

	cli.rootCmd.PersistentFlags().
		StringArrayVar(&cli.hostEndpoints, "host-endpoint", cli.hostEndpoints, "Host Endpoint, [$HOST_ENDPOINT]")

	cli.v.BindPFlag(hostEndpointKey, cli.rootCmd.PersistentFlags().Lookup("host-endpoint"))

	registerVMCmds(cli)
	registerVolumeCmds(cli)

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
