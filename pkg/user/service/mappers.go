package service

import "github.com/vier21/pc-01-network-be/pkg/user/domain"

func setAuthenticatedUser(user *domain.User) *DataUserAuthenticated {
	return &DataUserAuthenticated{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}
}

func toUserFromNewUser(user domain.NewUser) *domain.User {
	return &domain.User{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		HashPassword: user.Password,
		Role:      user.Role,
	}
}

func setToDataSecAuth(authuser DataUserAuthenticated, datasec DataSecurityAuthenticated) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data:     authuser,
		Security: datasec,
	}
}
