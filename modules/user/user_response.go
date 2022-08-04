package user

import "crm-sebagian-team/domain"

func NewListUserResponse(datas []domain.User) []domain.User {
	resp := []domain.User{}

	for _, data := range datas {
		resp = append(resp, NewUserResponse(data))
	}

	return resp
}

func NewUserResponse(user domain.User) domain.User {
	return domain.User{
		ID:          user.ID,
		Code:        user.Code,
		Name:        user.Name,
		Description: user.Description,
	}
}
