package wechat3rd

import "errors"

type TicketServer interface {
	SetTicket(ticket string) error
	GetTicket() (string, error)
}

type defaultTicketServer struct {
	ComponentTicketCache string // *accessToken
}

var defaultTicketServerHander TicketServer = &defaultTicketServer{}


var _ TicketServer = (*defaultTicketServer)(nil)

func (cts *defaultTicketServer) GetTicket() (string, error) {
	if cts.ComponentTicketCache == "" {
		return "", errors.New("component ticket is null")
	}
	return cts.ComponentTicketCache, nil
}

func (cts *defaultTicketServer) SetTicket(v string) error {
	cts.ComponentTicketCache = v
	return nil
}
