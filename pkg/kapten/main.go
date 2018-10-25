package kapten

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var kaptenConfigPath = ""

// NewCLI creates cobra object for top-level kapten server
func NewCLI(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kapten",
		Short: "Kapten is server for GoDC datacenter and cluster management system",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initConfig(cmd, defaultConfig)
		},
	}
	cmd.SetOutput(out)

	configGlobalFlags(cmd.PersistentFlags())

	cmd.AddCommand(NewCmdLocal(out, defaultConfig))
	cmd.AddCommand(NewCmdFunctions(out, defaultConfig))
	cmd.AddCommand(NewCmdImages(out, defaultConfig))
	cmd.AddCommand(NewCmdSecrets(out, defaultConfig))
	cmd.AddCommand(NewCmdEvents(out, defaultConfig))
	cmd.AddCommand(NewCmdAPIs(out, defaultConfig))
	cmd.AddCommand(NewCmdIdentity(out, defaultConfig))
	cmd.AddCommand(NewCmdServices(out, defaultConfig))

	return cmd
}

func initConfig(cmd *cobra.Command, targetConfig *kaptenConfig) {
	v := viper.New()
	configPath := os.Getenv("GODC_CONFIG")

	if kaptenConfigPath != "" {
		configPath = kaptenConfigPath
	}

	if configPath != "" {
		v.SetConfigFile(kaptenConfigPath)
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
