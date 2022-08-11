package repository

import (
	"context"
	"crm-sebagian-team/domain"
	"crm-sebagian-team/modules/user"
	"gorm.io/gorm"
)

type positionRepository struct {
	Conn *gorm.DB
}

func NewPositionRepository(Conn *gorm.DB) domain.PositionRepository {
	return &positionRepository{Conn}
}

func (m *positionRepository) GetByID(ctx context.Context, id int64) (domain.Position, error) {
	role := user.Position{}

	query := m.Conn.Model(&role).Where("id = ?", id).Take(&role)

	if query.Error != nil {
		return domain.Position{}, query.Error
	}

	response := domain.Position{
		ID:          int64(role.ID),
		Name:        role.Name,
		Description: role.Description,
	}

	return response, nil
}
