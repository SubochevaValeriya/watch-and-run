package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"time"
	"watchAndRun/internal/app/watch-and-run/model"
)

type globalConfig struct {
	DBConfig             DBConfig          `yaml:"db"`
	DBTables             DBTables          `yaml:"db_tables"`
	PathAndCommands      []PathAndCommands `yaml:"path_and_commands"`
	ChangeCheckFrequency time.Duration     `yaml:"change_check_frequency"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type DBTables struct {
	Event  string `yaml:"event"`
	Launch string `yaml:"launch"`
}

type PathAndCommands struct {
	Path          string   `yaml:"path"`
	Commands      []string `yaml:"commands"`
	IncludeRegexp []string `yaml:"include_regexp"`
	ExcludeRegexp []string `yaml:"exclude_regexp"`
	LogFile       string   `yaml:"log_file"`
}

func ParseConfig(path string) (globalConfig, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return globalConfig{}, fmt.Errorf("can't parse config: '%w'", err)
	}
	conf := globalConfig{}
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return globalConfig{}, fmt.Errorf("can't parse config: '%w'", err)
	}
	return conf, nil
}

func ImplementDirectoryStructure(PathAndCommands PathAndCommands) model.Directory {
	directory := model.Directory{
		Path:          PathAndCommands.Path,
		Commands:      PathAndCommands.Commands,
		IncludeRegexp: nil,
		ExcludeRegexp: nil,
		LogFile:       PathAndCommands.LogFile,
	}

	for _, regex := range PathAndCommands.IncludeRegexp {
		directory.IncludeRegexp = append(directory.IncludeRegexp, regexp.MustCompile(regex))

	}
	for _, regex := range PathAndCommands.ExcludeRegexp {
		directory.ExcludeRegexp = append(directory.ExcludeRegexp, regexp.MustCompile(regex))
	}

	return directory
}
