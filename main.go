package main

import (
	"cycleexemplo/service/message"
	"cycleexemplo/service/messagehandler"
	"cycleexemplo/service/messagesender"
	"cycleexemplo/service/whatsapp"
	"cycleexemplo/serviceinterfaces"
)

type Application struct {
	MessageSender   serviceinterfaces.MessageSender
	WhatsAppService serviceinterfaces.WhatsAppservice
	MessageHandler  serviceinterfaces.EventHandler
	MessageService  serviceinterfaces.Message
}

func BuildApplication() *Application {
	// Criar dependências com inicialização mínima
	msgSender := messagesender.NewMessageSender(nil)
	msgService := message.NewMessage("repo", msgSender)
	msgHandler := messagehandler.NewMessageHandler(msgService)
	whatsAppService := whatsapp.NewWhatsAppservice(msgHandler)

	// Configurar dependências cruzadas
	msgSender.SetWhatsAppService(whatsAppService)

	return &Application{
		MessageSender:   msgSender,
		WhatsAppService: whatsAppService,
		MessageHandler:  msgHandler,
		MessageService:  msgService,
	}
}
