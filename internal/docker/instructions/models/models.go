package models

type Instructions struct {
	FROM        *instruction `yaml:"FROM"`
	LABEL       *instruction `yaml:"LABEL"`
	ENV         *instruction `yaml:"ENV"`
	RUN         *instruction `yaml:"RUN"`
	COPY        *instruction `yaml:"COPY"`
	ADD         *instruction `yaml:"ADD"`
	CMD         *instruction `yaml:"CMD"`
	WORKDIR     *instruction `yaml:"WORKDIR"`
	ARG         *instruction `yaml:"ARG"`
	ENTYPOINT   *instruction `yaml:"ENTYPOINT"`
	EXPOSE      *instruction `yaml:"EXPOSE"`
	VOLUME      *instruction `yaml:"VOLUME"`
	MAINTAINER  *instruction `yaml:"MAINTAINER"`
	USER        *instruction `yaml:"USER"`
	ONBUILD     *instruction `yaml:"ONBUILD"`
	STOPSIGNAL  *instruction `yaml:"STOPSIGNAL"`
	HELATHCHECK *instruction `yaml:"HELATHCHECK"`
	SHELL       *instruction `yaml:"SHELL"`
}

type instruction struct {
	Name        string `yaml:"Name"`
	Description string `yaml:"Description"`
	Value       string `yaml:"Value"`
	
}
