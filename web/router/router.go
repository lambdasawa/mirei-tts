package router

import (
	"mirei-tts/web/endpoint/speech"
	"mirei-tts/web/endpoint/text"
	"mirei-tts/web/server"
)

func Set(s *server.Server) {
	s.Echo.File("/", "public/index.html")
	speech.Set(s)
	text.Set(s)

	s.Log.Info("routing initialized", nil)
}
