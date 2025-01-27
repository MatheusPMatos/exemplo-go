package message

import "cycleexemplo/serviceinterfaces"

type message struct {
	Repository string
	MsgSender  serviceinterfaces.MessageSender
}

type Message interface {
	SaveMessage(msg string)
	CreateMessage(content string) string
}

func NewMessage(repo string, sender serviceinterfaces.MessageSender) Message {
	return &message{
		Repository: repo,
		MsgSender:  sender,
	}
}

func (m *message) SaveMessage(msg string) {
	println("Mensagem salva:", msg)
}

func (m *message) CreateMessage(content string) string {
	return "Criado: " + content
}
