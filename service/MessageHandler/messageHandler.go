package messagehandler

import "cycleexemplo/service/message"

type messagehandler struct {
	MsgSV     message.Message // USAREI O SERVICO PARA BUSCAR E PERCISTIR MSGS
	WebSocket string          // servico de sockets para notificar front
}

type Messagehandler interface {
	HandleEvents(evt interface{})
	//HANDLE QUE SERA PASSADO PARA O CLIENT WHATSAPP ENVIAR OS EVENTOS
	//Recebe o evento verifica o tipo
	// ex: tipo.message -> ai chama o msg service para salvar a mensagem
	// 					-> notifica o front com o websocket passando a nova msg
}
