package main

import "github.com/spf13/cobra"

type createCmd struct {
	*createVmCmd
	*createVolumeCmd
}

type createVmCmd struct {
	// name of VM to be created
	name string
}

type createVolumeCmd struct {
	// name of volume to be created
	name string
	// source is URL where base image from
	source string
}

func registerCreateCmds(c *Cli) {
	c.createCmd = &createCmd{
		createVmCmd:     &createVmCmd{},
		createVolumeCmd: &createVolumeCmd{},
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create resources",
		RunE:  c.usageRunner(),
	}

	createVmCmd := &cobra.Command{
		Use:     "vm",
		Short:   "create vm",
		Example: `godc create vm`,
		PreRunE: c.preRunner(createVmPre),
		RunE:    c.runner(createVmRun),
	}

	createVolumeCmd := &cobra.Command{
		Use:     "volume",
		Short:   "create volume",
		Example: `godc create volume`,
		PreRunE: c.preRunner(createVolumePre),
		RunE:    c.runner(createVolumeRun),
	}

	createCmd.AddCommand(createVmCmd)
	createCmd.AddCommand(createVolumeCmd)
	c.rootCmd.AddCommand(createCmd)

	createVmCmd.Flags().StringVarP(&c.createCmd.createVmCmd.name, "name", "n", "", "VM name to be created")

	createVolumeCmd.Flags().StringVarP(&c.createCmd.createVolumeCmd.name, "name", "n", "", "Volume name to be created")
	createVolumeCmd.Flags().StringVarP(&c.createCmd.createVolumeCmd.name, "source", "s", "", "Source is URL where base image from")
}
