package service

import (
	"context"
	"errors"
	"gologin/model"
)

var (
	ErrClientNotEXist = errors.New("ClientID is not exist ")
	ErrClientSecret   = errors.New("invalid clientSecret")
)

type ClientDetialService interface {
	//Get details by clientId
	GetClientDetailByClientId(ctx context.Context, clientId string, clientSecret string) (*model.ClientDetails, error)
}

type InMemoryClientDeatilService struct {
	clientDetailsDict map[string]*model.ClientDetails
}

func (service *InMemoryClientDeatilService) GetClientDetailByClientId(ctx context.Context, clientId string, clientSecret string) (*model.ClientDetails, error) {

	clientDetails, ok := service.clientDetailsDict[clientId]
	if ok {
		if clientDetails.ClientSecret == clientSecret {
			return clientDetails, nil
		} else {
			return nil, ErrClientSecret
		}

	} else {
		return nil, ErrClientNotEXist
	}
}

func NewInMemoryClientDetailsService(clientDetailslist []*model.ClientDetails) *InMemoryClientDeatilService {
	clientDetailsDict := make(map[string]*model.ClientDetails)

	if len(clientDetailsDict) != 0 {
		for _, value := range clientDetailslist {
			clientDetailsDict[value.ClientId] = value
		}
	}

	return &InMemoryClientDeatilService{
		clientDetailsDict: clientDetailsDict,
	}
}
