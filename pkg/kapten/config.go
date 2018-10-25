package kapten

import "github.com/spf13/pflag"

type kaptenConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

var defaultConfig = &kaptenConfig{}

func configGlobalFlags(flags *pflag.FlagSet) {
	flags.StringVar(&kaptenConfigPath, "config", "", "config file to use")

	flags.String("host", "127.0.0.1", "Host/IP to listen on")
	flags.Int("port", 8080, "HTTP port to listen on")
}
