package whatsapp

import (
	messagehandler "cycleexemplo/service/MessageHandler"
)

type whatsAppservice struct {
	Clts map[uint]*Client //a lista das conexoes dos diferentes numeros conectados no sistema
	// conforme a necessidade busco o cliente correspondente para enviar a mensagem
	Handler messagehandler.Messagehandler //INJETO O HANDLE SERVICE PARA PASSAR A FUNCAO EVENTHANDLER
}

type WhatsAppservice interface {
	InitConections()
}

type Client struct { // simula o client whatsapp fornecido pelo package whatsmeow
	Handler EventHandler
}

func (c *Client) AddEventHandler(handler EventHandler) {
	// a tal funcao que o client tem onde passamos a funcao handler que
	// ira manejar os eventos

	c.Handler = handler
}

type EventHandler func(evt interface{})
