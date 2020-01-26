package tn

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Tn tinet config
type Tn struct {
	PreCmd      []PreCmd     `yaml:"precmd"`
	PreInit     []PreInit    `yaml:"preinit"`
	PostInit    []PostInit   `yaml:"postinit"`
	PostFini    []PostFini   `yaml:"postfini"`
	Nodes       []Node       `yaml:"nodes" mapstructure:"nodes"`
	Switches    []Switch     `yaml:"switches" mapstructure:"switches"`
	NodeConfigs []NodeConfig `yaml:"node_configs" mapstructure:"node_configs"`
	Test        []Test       `yaml:"test"`
}

// PreCmd
type PreCmd struct {
	// Cmds []Cmd `yaml:"cmds"`
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

// PreInit
type PreInit struct {
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

// PostInit
type PostInit struct {
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

// PostFini
type PostFini struct {
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

// Node
type Node struct {
	Name       string      `yaml:"name"`
	Type       string      `yaml:"type"`
	NetBase    string      `yaml:"net_base"`
	Image      string      `yaml:"image"`
	Interfaces []Interface `yaml:"interfaces" mapstructure:"interfaces"`
	Sysctls    []Sysctl    `yaml:"sysctls" mapstructure:"sysctls"`
	Mounts     []string    `yaml:"mounts,flow"`
}

// Interface
type Interface struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Args string `yaml:"args"`
	Addr string `yaml:"addr"`
}

// Sysctl
type Sysctl struct {
	Sysctl string `yaml:"string"`
}

// Switch
type Switch struct {
	Name       string      `yaml:"name"`
	Interfaces []Interface `yaml:"interfaces" mapstructure:"interfaces"`
}

// NodeConfig
type NodeConfig struct {
	Name string `yaml:"name"`
	Cmds []Cmd  `yaml:"cmds" mapstructure:"cmds"`
}

// Cmd
type Cmd struct {
	Cmd string `yaml:"cmd"`
}

// Test
type Test struct {
	Name string
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

func ReadTnConfig(fileName string) (tnconfig Tn, err error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return tnconfig, err
	}

	err = yaml.Unmarshal(buf, &tnconfig)
	if err != nil {
		return tnconfig, err
	}

	return tnconfig, nil
}
