package main

import "github.com/spf13/cobra"

type vmCmd struct {
	name string
}

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

	createVmCmd := &cobra.Command{
		Use:     "create",
		Short:   "create vm",
		Example: `godc vm create --host-endpoint HOST_ENDPOINT1 --name VM_NAME`,
		PreRunE: c.preRunner(createVmPre),
		RunE:    c.runner(createVmRun),
	}

	vmCmd.AddCommand(listVmCmd)
	vmCmd.AddCommand(createVmCmd)
	c.rootCmd.AddCommand(vmCmd)

	createVmCmd.Flags().StringVarP(&c.vmCmd.name, "name", "n", "", "VM name to be created")
}
