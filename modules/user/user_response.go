package user

import "crm-sebagian-team/domain"

func NewListUserResponse(datas []domain.User) []domain.UserResponse {
	resp := []domain.UserResponse{}

	for _, data := range datas {
		resp = append(resp, NewUserResponse(data))
	}

	return resp
}

func NewUserResponse(user domain.User) domain.UserResponse {
	return domain.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
