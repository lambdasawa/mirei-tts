package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		OriginalFilePath    string    `yaml:"OriginalFilePath"`
		OutputDirectoryPath string    `yaml:"OutputDirectoryPath"`
		FileDefList         []FileDef `yaml:"FileDefList"`
	}

	FileDef struct {
		Name     string        `yaml:"Name"`
		Start    time.Duration `yaml:"Start"`
		Duration time.Duration `yaml:"Duration"`
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
		if err := trim(
			conf.OriginalFilePath,
			filepath.Join(conf.OutputDirectoryPath, fmt.Sprintf("%s.wav", fileDef.Name)),
			fileDef.Start,
			fileDef.Duration,
		); err != nil {
			return err
		}
	}

	log.Println("Success.")

	return nil
}

func trim(inFileName, outFileName string, startDuration, lengthDuration time.Duration) error {
	file, err := os.Open(inFileName)
	if err != nil {
		return err
	}

	streamer, format, err := wav.Decode(file)
	if err != nil {
		return err
	}
	defer streamer.Close()

	if err := streamer.Seek(format.SampleRate.N(startDuration)); err != nil {
		return err
	}

	s := beep.Take(format.SampleRate.N(lengthDuration), streamer)

	outFile, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := wav.Encode(outFile, s, format); err != nil {
		return err
	}

	return nil
}
