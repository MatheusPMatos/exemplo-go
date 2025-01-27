package messagesender

import "cycleexemplo/serviceinterfaces"

type messagesender struct {
	WhatsSV serviceinterfaces.WhatsAppservice
}

type MessageSender interface {
	SendMessage(content string)
	SetWhatsAppService(whatsApp serviceinterfaces.WhatsAppservice)
}

func NewMessageSender(whatsApp serviceinterfaces.WhatsAppservice) MessageSender {
	return &messagesender{
		WhatsSV: whatsApp,
	}
}

func (ms *messagesender) SendMessage(content string) {
	if ms.WhatsSV != nil {
		ms.WhatsSV.SendToClient(1, content) // Envia mensagem ao cliente 1
	}
	println("Mensagem enviada:", content)
}

func (ms *messagesender) SetWhatsAppService(whatsApp serviceinterfaces.WhatsAppservice) {
	ms.WhatsSV = whatsApp
}
