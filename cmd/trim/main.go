package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		OriginalFilePath    string    `yaml:"OriginalFilePath"`
		OutputDirectoryPath string    `yaml:"OutputDirectoryPath"`
		FileDefList         []FileDef `yaml:"FileDefList"`
	}

	FileDef struct {
		Name     string `yaml:"Name"`
		Start    string `yaml:"Start"`
		Duration string `yaml:"Duration"`
	}
)

var (
	configPath = "trim.yaml"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("Failure.", err)
	}
}

func run() error {
	fileBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	conf := Config{}
	if err := yaml.Unmarshal(fileBytes, &conf); err != nil {
		return err
	}

	log.Printf("Config = %+v", conf)

	for _, fileDef := range conf.FileDefList {
		log.Printf("Target = %#v", fileDef)

		cmdName := "sox"
		cmdArgs := []string{
			conf.OriginalFilePath,
			filepath.Join(conf.OutputDirectoryPath, fmt.Sprintf("%s.wav", fileDef.Name)),
			"trim",
			fmt.Sprint(fileDef.Start),
			fmt.Sprint(fileDef.Duration),
		}

		var (
			stdout = new(bytes.Buffer)
			stderr = new(bytes.Buffer)
		)

		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		log.Printf("Command = %s %s", cmdName, strings.Join(cmdArgs, " "))

		if err := cmd.Run(); err != nil {
			log.Printf("sox stdout: %s", stdout.String())
			log.Printf("sox stderr: %s", stderr.String())

			return fmt.Errorf("run cmd: %v", err)
		}
		log.Printf("sox stdout: %s", stdout.String())
		log.Printf("sox stderr: %s", stderr.String())

	}

	log.Println("Success.")

	return nil
}
