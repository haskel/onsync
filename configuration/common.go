package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ParseConfig() (Config, error) {
	configuration := Config{}
	configData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Can't read config: %v", err)
		return configuration, err
	}

	err = yaml.Unmarshal(configData, &configuration)
	if err != nil {
		log.Fatalf("Parse config error: %v", err)
		return configuration, err
	}

	//pathArg := flag.String("n", "", "path")
	//flag.Parse()
	//
	//path, err := filepath.Abs(*pathArg)
	//if err != nil {
	//	log.Fatalf("Failed to resolve path: %v", err)
	//}
	//
	//fmt.Printf("Watch: %s\n", path)

	return configuration, nil
}
