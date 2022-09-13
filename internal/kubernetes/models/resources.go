package models

type Resources struct {
	ResourceMass []*Resource `yaml:"Resources"`
}

type Resource struct {
	Name        string   `yaml:"Name"`
	Description string   `yaml:"Description"`
	Fields      []*Field `yaml:"Fields"`
}

type Field struct {
	Name          string   `yaml:"Name"`
	Description   string   `yaml:"Description"`
	Type          string   `yaml:"Type"`
	PatchStrategy string   `yaml:"PatchStrategy"`
	PatchMergeKey string   `yaml:"PatchMergeKey"`
	Fields        []*Field `yaml:"Fields"`
}
