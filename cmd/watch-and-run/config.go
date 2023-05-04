package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"watchAndRun/internal/app/watch-and-run/model"
)

type globalConfig struct {
	DBConfig        DBConfig          `yaml:"db"`
	DBTables        DBTables          `yaml:"db_tables"`
	PathAndCommands []PathAndCommands `yaml:"path_and_commands"`
}

type DBConfig struct {
	Host     string `yaml:"db.host"`
	Port     string `yaml:"db.port"`
	Username string `yaml:"db.username"`
	Password string `yaml:"db.password"`
	DBName   string `yaml:"db.dbname"`
	SSLMode  string `yaml:"db.sslmode"`
}

type DBTables struct {
	Event  string `yaml:"dbTables.event"`
	Launch string `yaml:"dbTables.launch"`
}

type PathAndCommands struct {
	Path          string   `yaml:"path"`
	Commands      []string `yaml:"commands"`
	IncludeRegexp []string `yaml:"include_regexp"`
	ExcludeRegexp []string `yaml:"exclude_regexp"`
	LogFile       string   `yaml:"log_file"`
}

func parseConfig(path string) (globalConfig, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return globalConfig{}, fmt.Errorf("can't parse config: '%w'", err)
	}
	conf := globalConfig{}
	fmt.Printf("content='%#v'\n", string(content))
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return globalConfig{}, fmt.Errorf("can't parse config: '%w'", err)
	}
	fmt.Printf("conf='%#v'\n", conf)
	return conf, nil
}

func implementDirectoryStructure(PathAndCommands PathAndCommands) model.Directory {
	result := model.PathAndCommands{
		Path:          PathAndCommands.Path,
		Commands:      PathAndCommands.Commands,
		IncludeRegexp: nil,
		ExcludeRegexp: nil,
		LogFile:       PathAndCommands.LogFile,
	}

	for _, regexp := range PathAndCommands.IncludeRegexp {
		result.IncludeRegexp[regexp] = struct{}{}
	}
	for _, regexp := range PathAndCommands.ExcludeRegexp {
		result.ExcludeRegexp[regexp] = struct{}{}
	}

	return result
}
