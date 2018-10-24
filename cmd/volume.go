package main

import "github.com/spf13/cobra"

type volumeCmd struct {
	name string
}

func registerVolumeCmds(c *Cli) {
	c.volumeCmd = &volumeCmd{}

	volumeCmd := &cobra.Command{
		Use:   "volume",
		Short: "volume related commands",
		RunE:  c.usageRunner(),
	}

	listVolumeCmd := &cobra.Command{
		Use:     "list",
		Short:   "list volume",
		Example: `godc volume list --host-endpoint HOST_ENDPOINT1 --host-endpoint HOST_ENDPOINT2`,
		PreRunE: c.preRunner(listVolumePre),
		RunE:    c.runner(listVolumeRun),
	}

	createVolumeCmd := &cobra.Command{
		Use:     "create",
		Short:   "create volume",
		Example: `godc volume create --host-endpoint HOST_ENDPOINT --name VOLUME_NAME`,
		PreRunE: c.preRunner(createVolumePre),
		RunE:    c.runner(createVolumeRun),
	}

	volumeCmd.AddCommand(listVolumeCmd)
	volumeCmd.AddCommand(createVolumeCmd)
	c.rootCmd.AddCommand(volumeCmd)

	createVolumeCmd.Flags().StringVarP(&c.volumeCmd.name, "name", "n", "", "volume name to be created")
}
