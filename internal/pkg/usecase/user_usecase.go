package usecase

import (
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"backend-food/pkg/share/middleware"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserUsecase struct {
	Repo service.UserRepositoryInterface
}

func (u *UserUsecase) GetUserTokenLogin(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := u.Repo.FirstUser(entity.Users{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return dto.LoginResponse{}, err
	}
	if user.Password != req.Password {
		return dto.LoginResponse{}, errors.New("Login fail")
	}
	timeNow := time.Now()

	timeExpriedAt := timeNow.Add(time.Hour * 2)
	// generate uuid
	uuid := uuid.Must(uuid.NewV4(), nil)
	tokenString, err := middleware.GenerateJWTToken(middleware.JWTParam{
		UUID:       uuid,
		Authorized: true,
		ExpriedAt:  timeExpriedAt,
	})
	if err != nil {
		return dto.LoginResponse{}, err
	}
	newUser := entity.Users{
		Token:          &tokenString,
		TokenExpriedAt: &timeExpriedAt,
	}
	err = u.Repo.UpdateUser(newUser, user)

	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: tokenString,
	}, nil
}

func NewUserUsecase(repo service.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{Repo: repo}
}
