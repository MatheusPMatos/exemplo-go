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
	WhatsAppService serviceinterfaces.WhatsAppService
	MessageHandler  serviceinterfaces.EventHandler
	MessageService  serviceinterfaces.Message
}

func BuildApplication() *Application {
	// 1. Criação de instâncias básicas
	whatsAppService := whatsapp.NewWhatsAppService()
	msgSender := messagesender.NewMessageSender()
	msgService := message.NewMessage("repo", msgSender)

	// 2. Configuração das dependências cruzadas
	msgSender.SetWhatsAppService(whatsAppService)

	// Passa o msgService para o handler
	msgHandler := messagehandler.NewMessageHandler(msgService)

	whatsAppService.InitConnections(msgHandler)

	return &Application{
		MessageSender:   msgSender,
		WhatsAppService: whatsAppService,
		MessageHandler:  msgHandler,
		MessageService:  msgService,
	}
}

func main() {
	app := BuildApplication()
	app.MessageHandler.HandleEvents("Olá, mensagem de teste!")
	app.WhatsAppService.SendToClient(1, "Mensagem para cliente!")
}
