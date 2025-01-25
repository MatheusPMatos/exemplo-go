package messagesender

import (
	whatsapp "cycleexemplo/service/WhatsApp"
)

type messagesender struct {
	WhatsSV whatsapp.WhatsAppservice //TENHO QUE BUSCAR OS CLIENTS WHATSAPP PARA PODER ENVIAR MSGS ETC
}

type MessageSender interface {
	SendMessage()
}
