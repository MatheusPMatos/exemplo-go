package whatsapp

import "cycleexemplo/serviceinterfaces"

type Client struct {
	Handler serviceinterfaces.EventHandler
}

func (c *Client) AddEventHandler(handler serviceinterfaces.EventHandler) {
	c.Handler = handler
}

type WhatsAppService struct {
	Clients map[uint]*Client
}

func NewWhatsAppService() *WhatsAppService {
	return &WhatsAppService{
		Clients: make(map[uint]*Client),
	}
}

func (ws *WhatsAppService) InitConnections(handler serviceinterfaces.EventHandler) {
	client := &Client{Handler: handler}
	ws.Clients[1] = client
	println("Conexões iniciadas")
}

func (ws *WhatsAppService) SendToClient(clientID uint, message string) {
	client, exists := ws.Clients[clientID]
	if !exists {
		println("Cliente não encontrado:", clientID)
		return
	}

	if client.Handler != nil {
		client.Handler.HandleEvents(message)
	} else {
		println("Nenhum handler associado ao cliente:", clientID)
	}
}
