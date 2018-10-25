package main

import (
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var godcConfigPath = ""

// NewCLI creates cobra object for top-level kapten server
func NewCLI(out io.Writer) *cobra.Command {
	log.SetOutput(out)
	cmd := &cobra.Command{
		Use:   "godc",
		Short: "GoDC datacenter and cluster management system",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initConfig(cmd, defaultConfig)
		},
	}
	cmd.SetOutput(out)

	configGlobalFlags(cmd.PersistentFlags())

	cmd.AddCommand(NewCmdKopral(out, defaultConfig))

	return cmd
}

func initConfig(cmd *cobra.Command, targetConfig *godcConfig) {
	v := viper.New()
	configPath := os.Getenv("GODC_CONFIG")

	if godcConfigPath != "" {
		configPath = godcConfigPath
	}

	if configPath != "" {
		v.SetConfigFile(godcConfigPath)
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("Unable to read the config file: %s", err)
		}
	}

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		v.BindPFlag(f.Name, f)
		v.BindEnv(f.Name, "GODC_"+strings.ToUpper(strings.Replace(f.Name, "-", "_", -1)))
	})
	err := v.Unmarshal(targetConfig)
	if err != nil {
		log.Fatalf("Unable to create configuration: %s", err)
	}
}

func main() {
	cli := NewCLI(os.Stdout)
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
