package main

import (
	"mirei-tts/web"
)

func main() {
	if err := web.Start(); err != nil {
		panic(err)
	}
}
