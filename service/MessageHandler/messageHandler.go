package messagehandler

import "cycleexemplo/service/message"

type messagehandler struct {
	MsgSV message.Message // USAREI O SERVICO PARA BUSCAR E PERCISTIR MSGS
}

type Messagehandler interface {
	HandleEvents(evt interface{}) //HANDLE QUE SERA PASSADO PARA O CLIENT WHATSAPP ENVIAR OS EVENTOS
}
