package services

import (
	"github.com/personal_proj/hotel_management/domain/offer"
	"github.com/personal_proj/hotel_management/utils/errors"
)

var (
	//UsersService available for other to call the methods
	OfferService offerServiceInterface = &offerService{}
)

type offerService struct {
}

type offerServiceInterface interface {
	InsertOffer(map[string]interface{}) (map[string]interface{}, *errors.RestErr)
	// CreateUser(users.User) (*users.User, *errors.RestErr)
	// GetUser(int64) (*users.User, *errors.RestErr)
	// UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	// DeleteUser(int64) *errors.RestErr
	// Search(string) (users.Users, *errors.RestErr)
}

//CreateUser ...
func (s *offerService) InsertOffer(off map[string]interface{}) (map[string]interface{}, *errors.RestErr) {
	err := offer.Save(off)
	if err != nil {
		return nil, err
	}
	return off, nil

}
