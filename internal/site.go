package app

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Site []struct {
	Name string `yaml:"site"`
	URI  string `yaml:"link"`
}

func RetrieveSites() (Site, error) {
	var sites Site
	yamlFile, err := os.ReadFile("./internal/site.yaml")
	if err != nil {
		log.Fatalln("yamlFile get err", err)
	}
	err = yaml.Unmarshal(yamlFile, &sites)
	if err != nil {
		log.Fatalln("Unmarshal err", err)
	}
	return sites, err
}
