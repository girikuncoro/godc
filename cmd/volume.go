package main

import "github.com/spf13/cobra"

type volumeCmd struct{}

func registerVolumeCmds(c *Cli) {
	c.volumeCmd = &volumeCmd{}

	volumeCmd := &cobra.Command{
		Use:   "volume",
		Short: "Volume related commands",
		RunE:  c.usageRunner(),
	}

	listVolumeCmd := &cobra.Command{
		Use:     "list",
		Short:   "list volume",
		Example: `godc volume list --host-endpoint HOST_ENDPOINT1 --host-endpoint HOST_ENDPOINT2`,
		PreRunE: c.preRunner(listVolumePre),
		RunE:    c.runner(listVolumeRun),
	}

	volumeCmd.AddCommand(listVolumeCmd)
	c.rootCmd.AddCommand(volumeCmd)
}
