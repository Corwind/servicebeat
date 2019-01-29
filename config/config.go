// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
        Command string `config:"command"`
        Services []string `config:"services"`
        Regex string `config:"regex"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
        Command: "systemctl check",
        Services: []string {"sshd"},
        Regex: "active",
}
