package prononce

import "github.com/ikawaha/kagome/tokenizer"

func Generate(text string) string {
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
