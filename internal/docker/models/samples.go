package models

type Samples struct {
	SampleMass []*Sample `yaml:"Samples"`
}

type Sample struct {
	Name             string   `yaml:"Name"`
	Description      string   `yaml:"Description"`
	IsSelected       bool     `yaml:"IsSelected"`
	InstructionNames []string `yaml:"InstructionNames"`
}
