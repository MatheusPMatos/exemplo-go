package whatsapp

import (
	messagehandler "cycleexemplo/service/MessageHandler"
)

type whatsAppservice struct {
	Clts    map[uint]*Client
	Handler messagehandler.Messagehandler //INJETO O HANDLE SERVICE PARA PASSAR A FUNCAO EVENTHANDLER
}

type WhatsAppservice interface {
	InitConections()
}

type Client struct {
	Handler EventHandler
}

func (c *Client) AddEventHandler(handler EventHandler) {
	c.Handler = handler
}

type EventHandler func(evt interface{})
