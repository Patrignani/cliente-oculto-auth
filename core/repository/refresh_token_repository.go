package repository

import (
	"context"
	"time"

	"github.com/Patrignani/cliente-oculto-auth/core/data"
	"github.com/Patrignani/cliente-oculto-auth/core/entity"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/interfaces"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/specifications"
)

const (
	refreshTokenCollection = "refresh_tokens"
)

type RefreshTokenRepository struct {
	context data.IMongoContext
}

func NewRefreshTokenRepository(context data.IMongoContext) interfaces.IRefreshToken {
	return &RefreshTokenRepository{context: context}
}

func (c *RefreshTokenRepository) Insert(refreshToken *entity.RefreshToken) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	id, err := c.context.Insert(ctx, refreshTokenCollection, refreshToken)
	if err != nil {
		return err
	}

	refreshToken.ID = id

	return nil
}

func (u *RefreshTokenRepository) FindOneBySpecification(specification specifications.ISpecificationByOne) (*entity.RefreshToken, error) {
	filter, opts := specification.GetSpecification()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var refreshToken entity.RefreshToken
	if mgoErr := u.context.FindOne(ctx, refreshTokenCollection, filter, &refreshToken, opts); mgoErr != nil {
		return nil, mgoErr
	}

	return &refreshToken, nil
}

func (u *RefreshTokenRepository) Update(filter map[string]interface{}, fields interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := u.context.UpdateMany(ctx, refreshTokenCollection, filter, fields)

	return err
}
