package dto

import "01-quickstart/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"Telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
