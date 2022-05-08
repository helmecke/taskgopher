package config

import (
	"fmt"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

type config struct {
	DataDir string
}

// Config hold configuration
var Config config

// Init creates config with give file
func Init(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".taskgopher" (without extension).
		viper.AddConfigPath(xdg.ConfigHome + "/taskgopher")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.SetDefault("DataDir", xdg.DataHome+"/taskgopher")

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Printf("Unable to decode into config struct, %v", err)
	}
}
