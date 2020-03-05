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