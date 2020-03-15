package sound

import (
	"fmt"
	"io/ioutil"
	"mirei-tts/config"
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
)

func Generate(prononce string) (string, error) {
	sourcePaths := []string{}
	for _, p := range prononce {
		soundPath := filepath.Join(config.GetVoiceDirectory(), fmt.Sprintf("%c.wav", p))

		if _, err := os.Stat(soundPath); err != nil {
			continue
		}

		sourcePaths = append(sourcePaths, soundPath)
	}

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	destFilePath := filepath.Join(tempFile.Name())

	if err := merge(destFilePath, sourcePaths); err != nil {
		return "", err
	}

	return destFilePath, nil
}

func merge(destPath string, sourcePaths []string) error {
	files := make([]*os.File, 0)
	streams := make([]beep.Streamer, 0)
	formats := make([]beep.Format, 0)

	for _, path := range sourcePaths {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		streamer, format, err := wav.Decode(f)
		if err != nil {
			return err
		}
		defer streamer.Close()

		files = append(files, f)
		streams = append(streams, streamer)
		formats = append(formats, format)
	}

	if len(streams) <= 0 {
		return fmt.Errorf("can not speech")
	}

	stream := beep.Seq(streams...)

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if err := wav.Encode(destFile, stream, formats[0]); err != nil {
		return err
	}

	return nil
}
