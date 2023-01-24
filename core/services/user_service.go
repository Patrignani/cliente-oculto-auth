package services

import (
	"errors"

	"github.com/Patrignani/cliente-oculto-auth/core/entity"
	repository "github.com/Patrignani/cliente-oculto-auth/core/repository/interfaces"
	"github.com/Patrignani/cliente-oculto-auth/core/repository/specifications"
	"github.com/Patrignani/cliente-oculto-auth/core/services/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) interfaces.IUserService {
	return &UserService{repository: repository}
}

func (c *UserService) Authenticate(username string, password string) (*entity.User, error) {
	msgError := "Not authorized"
	specification := specifications.NewFindOneByUsernameAndActive(username, true,
		map[string]int{"_id": 1, "username": 1, "password": 1, "seed": 1, "roles": 1, "permissions": 1})

	user, err := c.repository.FindOneBySpecification(specification)

	if err != nil {
		return nil, errors.New(msgError)
	}

	if user == nil {
		return user, errors.New(msgError)
	}

	password += user.Seed
	passwordToCheck := []byte(password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), passwordToCheck)

	if err != nil {
		err = errors.New(msgError)
	}

	return user, err
}

func (c *UserService) FindById(userId string) (*entity.User, error) {
	project := map[string]int{
		"_id":         1,
		"username":    1,
		"roles":       1,
		"permissions": 1,
		"active":      1,
	}

	specification := specifications.NewFindByOneUserId(userId, true, project)

	return c.repository.FindOneBySpecification(specification)
}
