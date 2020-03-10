package config

import "os"

func GetAddress() string {
	value := os.Getenv("MTTS_ADDRESS")
	if value == "" {
		return ":5000"
	}
	return value
}

func GetVoiceDirectory() string {
	value := os.Getenv("MTTS_VOICE_DIRECTORY")
	if value == "" {
		return "voice"
	}
	return value
}

func GetTextSeedPath() string {
	value := os.Getenv("MTTS_TEXT_SEED_PATH")
	if value == "" {
		return "text-seed.json"
	}
	return value
}

func GetDictionaryPath() string {
	value := os.Getenv("MTTS_DICTIONARY_PATH")
	if value == "" {
		return "ipa.dic"
	}
	return value
}
