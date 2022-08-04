package repository

import (
	"context"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/user"

	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &userRepository{Conn}
}

func (m *userRepository) Get(ctx context.Context, params *domain.Request) (res []domain.User, total int64, err error) {
	var totalRows int64

	resUser := []domain.User{}
	bk := []user.User{}

	query := m.Conn.Model(&bk).Debug()

	query.Scopes(helpers.PaginateQuery(params)).Find(&bk).Count(&totalRows)
	if query.Error != nil {
		return []domain.User{}, 0, query.Error
	}

	for _, data := range bk {
		resUser = append(resUser, domain.User{
			ID:          int64(data.ID),
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
		})
	}

	return resUser, totalRows, nil
}
