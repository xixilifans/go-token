package service

import (
	"context"
	"errors"
	"gologin/model"
)

var (
	ErrUserNotExist = errors.New("username is not exist ")
	ErrPassword     = errors.New("invalid password")
)

type UserDetailsService interface {
	// Get UserDetails By username
	GetUserDetailByUsername(ctx context.Context, username, password string) (*model.UserDetails, error)
}

type InMemoryUserDetailsService struct {
	userDetailsDict map[string]*model.UserDetails
}

func (service *InMemoryUserDetailsService) GetUserDetailByUsername(ctx context.Context, username, password string) (*model.UserDetails, error) {
	//Get userDetails by username

	userDetails, ok := service.userDetailsDict[username]
	if ok {
		// match password
		if userDetails.Password == password {
			return userDetails, nil
		} else {
			return nil, ErrPassword
		}

	} else {
		return nil, ErrUserNotExist
	}
}

func NewInMemoryUserDetailsService(userDetailsList []*model.UserDetails) *InMemoryUserDetailsService {

	userDetailsDict := make(map[string]*model.UserDetails)
	if len(userDetailsDict) != 0 {
		for _, value := range userDetailsList {

			userDetailsDict[value.Username] = value
		}
	}

	return &InMemoryUserDetailsService{
		userDetailsDict: userDetailsDict,
	}

}
