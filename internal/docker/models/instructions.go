package models

type Instructions struct {
	InstructMass []*Instruction `yaml:"instructions"`
}

type Instruction struct {
	Name        string `yaml:"Name"`
	Description string `yaml:"Description"`
	Value       string
	IsMany      bool   `yaml:"IsMany"`
	PlaceHolder string `yaml:"PlaceHolder"`
	HintText    string `yaml:"HintText"`
	IsAdded     bool
}
