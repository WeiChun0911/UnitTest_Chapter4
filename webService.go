package chap4

type WebService struct {
	lastError string
}

func (ws *WebService)LogError(message string) {
	ws.lastError = message
}