package speech

import (
	"fmt"
	"io/ioutil"
	"log"
	"mirei-tts/config"
	"mirei-tts/web/request"
	"net/http"
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	"github.com/ikawaha/kagome/tokenizer"
	"github.com/labstack/echo/v4"
)

type (
	Req struct {
		Text string `query:"text" validate:"required,max=140"`
	}
)

func Get(c echo.Context) error {
	var req Req
	if err := request.Read(c, &req); err != nil {
		return fmt.Errorf("read request: %w", err)
	}
	log.Printf("req: %+v", req)

	filePath, err := convertTTS(req.Text)
	if err != nil {
		return fmt.Errorf("convert TTS: %v", err)
	}
	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Println(err)
		}
	}()

	http.ServeFile(c.Response(), c.Request(), filePath)
	return nil
}

func convertTTS(text string) (string, error) {
	prononce := generatePronounce(text)

	sourcePaths := []string{}
	for _, p := range prononce {
		soundPath := filepath.Join(config.GetVoiceDirectory(), fmt.Sprintf("%c.mp3", p))

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
	log.Printf("dest file path: %v", destFilePath)

	if err := mergeSounds(destFilePath, sourcePaths); err != nil {
		return "", err
	}

	return destFilePath, nil
}

func generatePronounce(text string) string {
	prononce := ""

	tokens := tokenizer.New().Tokenize(text)
	for _, t := range tokens {
		p := ""

		features := t.Features()
		if len(features) > 0 {
			p = features[len(features)-1]
		}
		if p == "*" {
			p = t.Surface
		}

		if p == "" {
			continue
		}

		prononce += p
	}

	return prononce
}

func mergeSounds(destPath string, sourcePaths []string) error {
	files := make([]*os.File, 0)
	streams := make([]beep.Streamer, 0)
	formats := make([]beep.Format, 0)

	for _, path := range sourcePaths {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		streamer, format, err := mp3.Decode(f)
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
