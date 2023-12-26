package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type ProbeConfig struct {
	ListenPort      int    `yaml:"listen_port"`
	ListenAddr      string `yaml:"listen_addr"`
	ListenInterface string `yaml:"listen_interface"`
	LogLevel        string `yaml:"log_level"`
	Workers         int    `yaml:"workers"`
}

func NewProbeConfig() *ProbeConfig {
	config := ProbeConfig{}

	yamlFile, err := os.Open("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err2 := yamlFile.Close(); err2 != nil {
			panic(err2)
		}
	}()
	decoder := yaml.NewDecoder(yamlFile)
	err = decoder.Decode(&config)
	config.getConfigFromFlags()
	return &config
}

func (c *ProbeConfig) getConfigFromFlags() {
	flagSet := flag.NewFlagSet("probe", flag.ContinueOnError)
	flagSet.IntVar(&c.ListenPort, "listen_port", c.ListenPort, "listen port")
	flagSet.StringVar(&c.ListenAddr, "listen_addr", c.ListenAddr, "listen address")
	flagSet.StringVar(&c.ListenInterface, "listen_interface", c.ListenInterface, "listen interface")
	flagSet.StringVar(&c.LogLevel, "log_level", c.LogLevel, "log level")
	flagSet.IntVar(&c.Workers, "workers", c.Workers, "workers")
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
}
