package speech

import (
	"fmt"
	"mirei-tts/prononce"
	"mirei-tts/sound"
	"mirei-tts/web/request"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type (
	Req struct {
		Text string `query:"text" validate:"required,max=140"`
	}
)

func (h *handler) get(c echo.Context) error {
	var req Req
	if err := request.Read(c, &req); err != nil {
		return fmt.Errorf("read request: %w", err)
	}

	prononceText, err := prononce.Generate(req.Text)
	if err != nil {
		return fmt.Errorf("generate prononce: %v", err)
	}
	h.Server.Log.Info("prononce text", map[string]interface{}{
		"value": prononceText,
	})

	filePath, err := sound.Generate(prononceText)
	if err != nil {
		return fmt.Errorf("convert TTS: %v", err)
	}
	defer func() {
		if err := os.Remove(filePath); err != nil {
			h.Server.Log.Warn(fmt.Sprintf("delete sound file: %v", err), nil)
		}
	}()

	http.ServeFile(c.Response(), c.Request(), filePath)
	return nil
}
