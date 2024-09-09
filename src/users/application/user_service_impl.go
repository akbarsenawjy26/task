package application

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"task/common/helper"
	"task/common/jwt"
	"task/src/repository"
)

type UserServiceImpl struct {
	userRepository    *repository.Queries
	transactionHelper *helper.TransactionHelper
	jwtService        jwt.ServiceJwt
}

func NewUserServiceImpl(userRepository *repository.Queries, transactionHelper *helper.TransactionHelper, jwtService jwt.ServiceJwt) UserService {
	return &UserServiceImpl{
		userRepository:    userRepository,
		transactionHelper: transactionHelper,
		jwtService:        jwtService,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, req repository.CreateUserParams) (repository.User, error) {
	var user repository.User
	err := service.transactionHelper.WithTransaction(ctx, func(tx *sql.Tx) error {
		var err error
		user, err = service.userRepository.CreateUser(ctx, req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return repository.User{}, err
	}
	return user, nil
}

func (service *UserServiceImpl) Update(ctx context.Context, req repository.UpdateUserParams) (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (service *UserServiceImpl) Delete(ctx context.Context, username string) (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (service *UserServiceImpl) GetByUsername(ctx context.Context, username, password string) (string, error) {
	user, err := service.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return "Username not found!", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "Password is incorrect!", err
	}
	token, err := service.jwtService.GenerateToken(username)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}
