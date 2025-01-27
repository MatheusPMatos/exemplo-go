package messagesender

import "cycleexemplo/serviceinterfaces"

type messagesender struct {
	WhatsAppService serviceinterfaces.WhatsAppService
}

type MessageSender interface {
	SendMessage(content string)
}

func NewMessageSender() *messagesender {
	return &messagesender{}
}

func (ms *messagesender) SetWhatsAppService(whatsApp serviceinterfaces.WhatsAppService) {
	ms.WhatsAppService = whatsApp
}

func (ms *messagesender) SendMessage(content string) {
	if ms.WhatsAppService != nil {
		ms.WhatsAppService.SendToClient(1, content)
	}
	println("Mensagem enviada:", content)
}
