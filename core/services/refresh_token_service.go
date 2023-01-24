package services

import (
	"strconv"
	"sync"
	"time"

	common "github.com/Patrignani/cliente-oculto-auth/core/config"
	"github.com/Patrignani/cliente-oculto-auth/core/entity"
	repository "github.com/Patrignani/cliente-oculto-auth/core/repository/interfaces"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/specifications"
	"github.com/Patrignani/cliente-oculto-auth/core/services/interfaces"
	"go.mongodb.org/mongo-driver/bson"
)

type RefreshTokenService struct {
	repository repository.IRefreshToken
}

func NewRefreshTokenService(repository repository.IRefreshToken) interfaces.IRefreshTokenSerice {
	return &RefreshTokenService{repository: repository}
}

func (r *RefreshTokenService) CreateRefreshToken(userID string) (*entity.RefreshToken, error) {
	var refresh *entity.RefreshToken = nil
	var erro error = nil

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		expireTimeMinutes, err := strconv.Atoi(common.RefreshToeknExpireTimeMinutes)

		if err != nil {
			erro = err
		}

		refreshToken := &entity.RefreshToken{
			UserID:         userID,
			Active:         true,
			CreateAt:       time.Now().UTC(),
			ExpirationDate: time.Now().Add(time.Minute * time.Duration(expireTimeMinutes)),
		}

		r.DisableAllWithUserId(userID)
		erro = r.repository.Insert(refreshToken)
		refresh = refreshToken
	}()

	go func() {
		defer wg.Done()
		r.DisableAllWithExpiredDate()
	}()

	wg.Wait()

	return refresh, erro

}

func (r *RefreshTokenService) FindById(refreshTokenId string) (*entity.RefreshToken, error) {
	project := map[string]int{
		"_id":             1,
		"expiration_date": 1,
		"user_id":         1,
		"active":          1,
	}

	specification := specifications.NewFindByOneRefreshTokenId(refreshTokenId, time.Now().UTC(), true, project)

	return r.repository.FindOneBySpecification(specification)
}

func (r *RefreshTokenService) DisableAllWithExpiredDate() error {
	now := time.Now().UTC()

	filter := bson.M{
		"$and": []bson.M{
			{"expiration_date": bson.M{"$lt": now}},
			{"active": true},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"active":   false,
			"UpdateAt": now,
		},
	}

	return r.repository.Update(filter, update)
}

func (r *RefreshTokenService) DisableAllWithUserId(userID string) error {
	now := time.Now().UTC()

	filter := bson.M{
		"$and": []bson.M{
			{"user_id": userID},
			{"active": true},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"active":   false,
			"UpdateAt": now,
		},
	}

	return r.repository.Update(filter, update)
}
