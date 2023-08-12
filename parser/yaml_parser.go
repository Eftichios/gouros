package parser

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Route struct {
	Endpoint    string `yaml:"endpoint"`
	Method      string `yaml:"method"` // change to enum
	Description string `yaml:"description"`
	Return      string `yaml:"return"`
}

type Entity struct {
	Resource    string  `yaml:"resource"`
	Description string  `yaml:"description"`
	Base        string  `yaml:"base"`
	Routes      []Route `yaml:"routes"`
}

type Attribute struct {
	Column string `yaml:"column"`
	Type   string `yaml:"type"` // ??
	Serial bool   `yaml:"serial"`
}

type Model struct {
	Table      string      `yaml:"table"`
	Attributes []Attribute `yaml:"attributes"`
	// Relationships []Relationship `yaml:"relationships"`
}

type ApiResource struct {
	Entity Entity  `yaml:"entity"`
	Models []Model `yaml:"models"`
}

func ParseYML(filename string) ApiResource {
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error reading provided yaml file: %v", err)
	}

	var apiResource ApiResource
	err = yaml.Unmarshal(yamlFile, &apiResource)

	if err != nil {
		log.Fatalf("Error unmarshalling yaml file: %v", err)
	}

	return apiResource
}
