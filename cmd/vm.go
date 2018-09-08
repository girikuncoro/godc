package main

import "github.com/spf13/cobra"

type vmCmd struct{}

func registerVMCmds(c *Cli) {
	c.vmCmd = &vmCmd{}

	vmCmd := &cobra.Command{
		Use:   "vm",
		Short: "VM related commands",
		RunE:  c.usageRunner(),
	}

	listVmCmd := &cobra.Command{
		Use:     "list",
		Short:   "list vm",
		Example: `godc vm list --host-endpoint HOST_ENDPOINT1 --host-endpoint HOST_ENDPOINT2`,
		PreRunE: c.preRunner(listVmPre),
		RunE:    c.runner(listVmRun),
	}

	vmCmd.AddCommand(listVmCmd)
	c.rootCmd.AddCommand(vmCmd)
}
