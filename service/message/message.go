package message

import messagesender "cycleexemplo/service/messageSender"

type message struct {
	Repository string
	MsgSender  messagesender.MessageSender // INJETO O SERVICE PARA PODER ENVIAR MSGS RECEBIDAS NOS ENDPOINTS
}

type Message interface {
	SaveMessage()
	CreateMessage()
}
