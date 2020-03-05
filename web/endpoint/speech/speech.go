package speech

import (
	"mirei-tts/web/server"
)

type handler struct {
	*server.Server
}

func Set(s *server.Server) {
	h := &handler{Server: s}

	h.Echo.GET("/speech", h.get)
}
