package serviceinterfaces

type Message interface {
	SaveMessage(msg string)
	CreateMessage(content string) string
}

type MessageSender interface {
	SendMessage(content string)
	SetWhatsAppService(whatsApp WhatsAppservice)
}

type WhatsAppservice interface {
	InitConnections()
	SendToClient(clientID uint, message string)
}

type EventHandler interface {
	HandleEvents(evt interface{})
}
