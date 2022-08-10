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
			ID:         int64(data.ID),
			Name:       data.Name,
			Email:      data.Email,
			Password:   data.Password,
			Address:    data.Address,
			IdPosition: data.IdPosition,
			CreatedAt:  data.CreatedAt,
			CreatedBy:  data.CreatedBy,
		})
	}

	return resUser, totalRows, nil
}

func (m *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	usr := user.User{}

	query := m.Conn.Model(&usr).Where("email = ?", email).Find(&usr)

	if query.Error != nil {
		return domain.User{}, query.Error
	}

	resUser := domain.User{
		ID:       int64(usr.ID),
		Name:     usr.Name,
		Email:    usr.Email,
		Password: usr.Password,
	}

	return resUser, nil
}

func (m *userRepository) Store(ctx context.Context, usr *domain.User) (domain.User, error) {
	usrEntity := user.User{
		Name:       usr.Name,
		Email:      usr.Email,
		Password:   usr.Password,
		Address:    usr.Address,
		IdPosition: usr.IdPosition,
		CreatedAt:  usr.CreatedAt,
		CreatedBy:  usr.CreatedBy,
		UpdatedAt:  usr.UpdatedAt,
	}

	query := m.Conn.Create(&usrEntity)

	if query.Error != nil {
		return domain.User{}, query.Error
	}

	usr.ID = int64(usrEntity.ID)

	return *usr, nil
}
