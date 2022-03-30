package chatbot

import "strconv"

const (
	pathBase = "/chatbot"
)

// PathWebhook returns the external webhook path used by a service worker
func PathWebhook(id int64) string { return pathBase + "/" + strconv.FormatInt(id, 10) }
