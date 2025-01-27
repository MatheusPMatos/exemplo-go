package whatsapp

import "cycleexemplo/serviceinterfaces"

// Cliente WhatsApp
type Client struct {
	Handler serviceinterfaces.EventHandler
}

func (c *Client) AddEventHandler(handler serviceinterfaces.EventHandler) {
	// Atribui o handler ao cliente
	c.Handler = handler
}

// Serviço do WhatsApp (agora exportado)
type WhatsAppservice struct {
	Clts    map[uint]*Client
	Handler serviceinterfaces.EventHandler
}

// Construtor exportado para WhatsAppservice
func NewWhatsAppservice(handler serviceinterfaces.EventHandler) *WhatsAppservice {
	return &WhatsAppservice{
		Clts:    make(map[uint]*Client),
		Handler: handler,
	}
}

// Inicializa conexões do WhatsApp
func (ws *WhatsAppservice) InitConnections() {
	client := &Client{}
	client.AddEventHandler(ws.Handler)

	ws.Clts[1] = client // Adiciona o cliente
	println("Conexões iniciadas")
}

// Envia mensagem para um cliente específico
func (ws *WhatsAppservice) SendToClient(clientID uint, message string) {
	client, exists := ws.Clts[clientID]
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
