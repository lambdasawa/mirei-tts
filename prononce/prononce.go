package prononce

import (
	"fmt"
	"mirei-tts/config"
	"path/filepath"
	"sync"

	"github.com/ikawaha/kagome/tokenizer"
)

var (
	globalTokenizer *tokenizer.Tokenizer
	mutex           = new(sync.Mutex)
)

func init() {
	go func() {
		_, _ = getTokenizer() // initialize asynchronously
	}()
}

func getTokenizer() (*tokenizer.Tokenizer, error) {
	mutex.Lock() // initialized asynchronously
	defer mutex.Unlock()

	if globalTokenizer != nil {
		return globalTokenizer, nil
	}

	conf := config.GetConfig()
	dic, err := tokenizer.NewDic(filepath.Join(conf.DataLocalPrefix, conf.DictionaryName))
	if err != nil {
		return nil, err
	}

	t := tokenizer.NewWithDic(dic)

	globalTokenizer = &t

	return globalTokenizer, nil
}

func Generate(text string) (string, error) {
	prononce := ""

	tokenizer, err := getTokenizer()
	if err != nil {
		return "", fmt.Errorf("get tokenizer: %v", err)
	}

	tokens := tokenizer.Tokenize(text) // take long time...
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

	return prononce, nil
}
