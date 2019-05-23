package main

import "github.com/spf13/cobra"

type getCmd struct{}

func registerGetCmds(c *Cli) {
	c.getCmd = &getCmd{}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get resources",
		RunE:  c.usageRunner(),
	}

	getVmCmd := &cobra.Command{
		Use:     "vm",
		Short:   "get vm",
		Example: `godc get vm`,
		PreRunE: c.preRunner(getVmPre),
		RunE:    c.runner(getVmRun),
	}

	getVolumeCmd := &cobra.Command{
		Use:     "volume",
		Short:   "get volume",
		Example: `godc get volume`,
		PreRunE: c.preRunner(getVolumePre),
		RunE:    c.runner(getVolumeRun),
	}

	getCmd.AddCommand(getVmCmd)
	getCmd.AddCommand(getVolumeCmd)
	c.rootCmd.AddCommand(getCmd)
}
