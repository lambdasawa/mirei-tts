package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/ikawaha/kagome/tokenizer"
	"github.com/mb-14/gomarkov"
)

type (
	Sentence struct {
		Original string
		Words    []string
	}

	TextSeed struct {
		Chain        *gomarkov.Chain
		InitialWords []string
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	println("get tweet")
	tweetDataPath := os.Getenv("TWEET_DATA_PATH")
	tweets := make([]twitter.Tweet, 0)
	if err := getOrFetch(tweetDataPath, &tweets, func() (interface{}, error) {
		println("fetch tweet")
		ts, err := fetchTweets()
		if err != nil {
			return nil, err
		}

		println("save tweet")
		if err := save(tweetDataPath, ts); err != nil {
			return nil, err
		}

		tweets = ts

		return ts, nil
	}); err != nil {
		panic(err)
	}

	println("get sentence")
	sentenceDataPath := os.Getenv("SENTENCE_DATA_PATH")
	sentenceList := make([]Sentence, 0)
	if err := getOrFetch(sentenceDataPath, &sentenceList, func() (interface{}, error) {
		println("find sentence")
		list, err := createSentenceList(tweets)
		if err != nil {
			return nil, err
		}

		println("save sentence")
		if err := save(sentenceDataPath, list); err != nil {
			return nil, err
		}

		sentenceList = list

		return list, nil
	}); err != nil {
		panic(err)
	}

	println("get text seed")
	textSeedDataPath := os.Getenv("TEXT_SEED_PATH")
	textSeed := TextSeed{}
	if err := getOrFetch(textSeedDataPath, &textSeed, func() (interface{}, error) {
		println("find text seed")
		ts := generateTextSeed(sentenceList, 1)

		println("save text seed")
		if err := save(textSeedDataPath, ts); err != nil {
			return nil, err
		}

		textSeed = ts

		return ts, nil
	}); err != nil {
		panic(err)
	}

	println(generateText(textSeed, rand.Intn(20)+5))
}

func getOrFetch(dataPath string, value interface{}, fetch func() (interface{}, error)) error {
	if _, err := os.Stat(dataPath); err != nil {
		// not found
		if _, err := fetch(); err != nil {
			return err
		}
		return nil
	}

	// found
	bytes, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, value); err != nil {
		return err
	}

	return nil
}

func save(dataPath string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dataPath, bytes, 0644); err != nil {
		return err
	}

	return nil
}

func fetchTweets() ([]twitter.Tweet, error) {
	config := oauth1.NewConfig(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
	)
	token := oauth1.NewToken(
		os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_ACCESS_SECRET"),
	)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	pageSize, _ := strconv.Atoi(os.Getenv("TWEET_PAGE_SIZE"))
	if pageSize == 0 {
		pageSize = 200
	}

	pageCount, _ := strconv.Atoi(os.Getenv("TWEET_PAGE_COUNT"))
	if pageCount == 0 {
		pageCount = 10
	}

	maxID := int64(0)
	tweets := make([]twitter.Tweet, 0)
	for i := 0; i < pageCount; i++ {
		ts, resp, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			ScreenName:      os.Getenv("TARGET_SCREEN_NAME"),
			ExcludeReplies:  twitter.Bool(true),
			IncludeRetweets: twitter.Bool(false),
			TrimUser:        twitter.Bool(true),
			Count:           pageSize,
			MaxID:           maxID,
		})
		if err != nil {
			return nil, err
		}
		resp.Body.Close()

		for _, t := range ts {
			tweets = append(tweets, t)
		}

		if len(ts) == 0 {
			break
		}
		maxID = tweets[len(tweets)-1].ID
	}

	return tweets, nil
}

func createSentenceList(tweets []twitter.Tweet) ([]Sentence, error) {
	list := make([]Sentence, 0)

	t := tokenizer.New()

	for _, tweet := range tweets {
		tokens := t.Tokenize(tweet.Text)

		words := make([]string, 0)
		for _, token := range tokens {
			w := token.Surface

			if w == "BOS" || w == "EOS" {
				continue
			}

			words = append(words, token.Surface)
		}

		list = append(list, Sentence{
			Original: tweet.Text,
			Words:    words,
		})
	}

	return list, nil
}

func generateTextSeed(sentences []Sentence, order int) TextSeed {
	chain := gomarkov.NewChain(order)

	initialWords := []string{}
	for _, s := range sentences {
		chain.Add(s.Words)
		initialWords = append(initialWords, s.Words[0])
	}

	return TextSeed{
		Chain:        chain,
		InitialWords: initialWords,
	}

}

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
