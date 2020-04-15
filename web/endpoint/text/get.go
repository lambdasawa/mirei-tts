package text

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mirei-tts/config"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mb-14/gomarkov"
)

type (
	// TODO refactor
	TextSeed struct {
		Chain        *gomarkov.Chain
		InitialWords []string
	}

	Res struct {
		Text string `json:"text"`
	}
)

const (
	minTextLength = 5
	maxTextLength = 30
)

func (h *handler) get(c echo.Context) error {
	conf := config.GetConfig()
	bytes, err := ioutil.ReadFile(filepath.Join(conf.DataLocalPrefix, conf.TextSeedName))
	if err != nil {
		return fmt.Errorf("open text-seed.json: %v", err)
	}

	textSeed := TextSeed{}
	if err := json.Unmarshal(bytes, &textSeed); err != nil {
		return fmt.Errorf("parse text-seed.json: %v", err)
	}

	text := generateText(textSeed, rand.Intn(maxTextLength-minTextLength)+minTextLength)

	return c.JSON(http.StatusOK, Res{
		Text: text,
	})
}

// TODO refactor
func generateText(textSeed TextSeed, count int) string {
	current := textSeed.InitialWords[rand.Int()%len(textSeed.InitialWords)]
	resultWords := []string{current}

	for i := 0; i < count; i++ {
		next, _ := textSeed.Chain.Generate([]string{current})

		current = next
		resultWords = append(resultWords, next)
	}

	return strings.Join(resultWords, "")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
