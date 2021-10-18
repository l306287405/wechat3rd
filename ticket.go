package wechat3rd

import "errors"

type TicketServer interface {
	SetTicket(ticket string) error
	GetTicket() (string, error)
}

type DefaultTicketServer struct {
	ComponentTicketCache string // *accessToken
}

var DefaultTicketServerHandler TicketServer = &DefaultTicketServer{}

var _ TicketServer = (*DefaultTicketServer)(nil)

func (cts *DefaultTicketServer) GetTicket() (string, error) {
	if cts.ComponentTicketCache == "" {
		return "", errors.New("component ticket is null")
	}
	return cts.ComponentTicketCache, nil
}

func (cts *DefaultTicketServer) SetTicket(v string) error {
	cts.ComponentTicketCache = v
	return nil
}
