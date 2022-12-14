package wechat3rd

import (
	"errors"
	"github.com/l306287405/wechat3rd/cache"
	"time"
)

type TicketServer interface {
	SetTicket(ticket string) error
	GetTicket() (string, error)
}

type DefaultTicketServer struct {
	Cache                cache.Cache
	ComponentTicketCache string // *accessToken
	Cfg                  Config
}

var _ TicketServer = (*DefaultTicketServer)(nil)

func (cts *DefaultTicketServer) GetTicket() (string, error) {
	if cts.ComponentTicketCache == "" {
		if c := cts.Cache.Get(cts.getCacheKey()); c != nil {
			return c.(string), nil
		}
		return "", errors.New("component ticket is null")
	}
	return cts.ComponentTicketCache, nil
}

func (cts *DefaultTicketServer) SetTicket(v string) error {
	cts.ComponentTicketCache = v
	//存缓存
	err := cts.Cache.Set(cts.getCacheKey(), v, 3600*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (cts *DefaultTicketServer) getCacheKey() string {
	return "wechat_open_platform.verify_ticket." + cts.Cfg.AppID
}
