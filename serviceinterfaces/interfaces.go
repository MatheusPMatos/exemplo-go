package serviceinterfaces

type Message interface {
	SaveMessage(msg string)
	CreateMessage(content string) string
}

type MessageSender interface {
	SendMessage(content string)
}

type WhatsAppService interface {
	InitConnections(handler EventHandler)
	SendToClient(clientID uint, message string)
}

type EventHandler interface {
	HandleEvents(evt interface{})
}
