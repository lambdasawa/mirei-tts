package config

import "os"

type (
	Config struct {
		Address            string
		DataBucket         string
		DataBucketPrefix   string
		DataLocalPrefix    string
		VoiceDirectoryName string
		TextSeedName       string
		DictionaryName     string
	}
)

func getEnvWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig() Config {
	return Config{
		Address:            getEnvWithDefault("MTTS_ADDRESS", ":5000"),
		DataBucket:         getEnvWithDefault("MTTS_DATA_BUCKET", "mireittsstack-databucketxxxx-xxxx"),
		DataBucketPrefix:   getEnvWithDefault("MTTS_DATA_BUCKET_PREFIX", ""),
		DataLocalPrefix:    getEnvWithDefault("MTTS_DATA_LOCAL_PREFIX", "data"),
		VoiceDirectoryName: getEnvWithDefault("DATA_VOICE_DIRECTORY_NAME", "voice"),
		TextSeedName:       getEnvWithDefault("DATA_TEXT_SEED_NAME", "text-seed.json"),
		DictionaryName:     getEnvWithDefault("DATA_DICTIONARY_NAME", "ipa.dic"),
	}
}
