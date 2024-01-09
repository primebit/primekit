package config

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/primebit/primekit/utils"
	"os"
)

var (
	Config      = &Configuration{}
	Environment = "dev"
)

func Init(path string) error {
	var err error

	_ = godotenv.Load(".env")
	Environment = os.Getenv("ENV")
	Environment = *flag.String("env", Environment, "set environment name")

	err = readMainConfig(path)
	return err
}

func readMainConfig(path string) error {
	return Reader{}.ReadYaml(Config, utils.RealPath(path+"/"+Environment+"/main.yml"))
}
