package messagehandler

import "cycleexemplo/serviceinterfaces"

type messagehandler struct {
	MsgSV     serviceinterfaces.Message
	WebSocket string
}

type Messagehandler interface {
	serviceinterfaces.EventHandler
}

func NewMessageHandler(msgService serviceinterfaces.Message) Messagehandler {
	return &messagehandler{
		MsgSV: msgService,
	}
}

func (mh *messagehandler) HandleEvents(evt interface{}) {
	if msg, ok := evt.(string); ok {
		mh.MsgSV.SaveMessage(msg)
		println("Evento tratado:", msg)
	}
}
