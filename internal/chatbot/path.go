package chatbot

const (
	pathBase = "/chatbot"
)

// PathWebhook returns the external webhook path used by a service worker
func PathWebhook(id string) string { return pathBase + "/" + id }
