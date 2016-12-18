package main

import yaml "gopkg.in/yaml.v2"

type Profile struct {
	Name    string `yaml:"Name"`
	Label   string `yaml:"Label"`
	Image   string `yaml:"Image"`
	Address string `yaml:"Address"`
	Phone   string `yaml:"Phone"`
	Email   string `yaml:"Email"`
	Github  string `yaml:"Github"`
	Summary string `yaml:Summary`
}

type Experience struct {
	Company   string `yaml:"Company"`
	Role      string `yaml:"Role"`
	StartDate string `yaml:"StartDate"`
	EndDate   string `yaml:"EndDate"`
	Summary   string `yaml:"Summary"`
}

type Project struct {
	Name      string `yaml:"Name"`
	Role      string `yaml:"Role"`
	StartDate string `yaml:"StartDate"`
	EndDate   string `yaml:"EndDate"`
	Summary   string `yaml:"Summary"`
}
type Education struct {
	Location  string `yaml:"Location"`
	Major     string `yaml:"Major"`
	StartDate string `yaml:"StartDate"`
	EndDate   string `yaml:"EndDate"`
	Summary   string `yaml:"Summary"`
}

type Resume struct {
	Profile     Profile      `yaml:"Profile"`
	Experiences []Experience `yaml:"Experiences,flow"`
	Projects    []Project    `yaml:"Projects,flow"`
	Educations  []Education  `yaml:"Educations,flow"`
}

func parseYaml(data []byte) (*Resume, error) {
	result := Resume{}
	err := yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
