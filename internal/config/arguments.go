// Author hoenig

package config

import "flag"

type Arguments struct {
	ConfigFile string
}

func ParseArguments() Arguments {
	var configFile string
	flag.StringVar(
		&configFile,
		"configfile",
		"/etc/ssh-key-sync.conf",
		"specify the configuration file location",
	)
	flag.Parse()
	return Arguments{
		ConfigFile: configFile,
	}
}
