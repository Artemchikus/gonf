package models

type Instructions struct {
	InstructMass []*Instruction `yaml:"Instructions"`
}

type Instruction struct {
	Name        string `yaml:"Name"`
	Description string `yaml:"Description"`
	IsMany      bool   `yaml:"IsMany"`
	PlaceHolder string `yaml:"PlaceHolder"`
	HintText    string `yaml:"HintText"`
	IsAdded     bool
	IsSelected  bool
}
