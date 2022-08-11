package service

import (
	"context"
	"crm-sebagian-team/config"
	"crm-sebagian-team/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	config         *config.Config
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewAuthService(c *config.Config, usrRepo domain.UserRepository, timeout time.Duration) domain.AuthService {
	return &authService{
		config:         c,
		userRepo:       usrRepo,
		contextTimeout: timeout,
	}
}

func newLoginResponse(token string) domain.LoginResponse {
	return domain.LoginResponse{
		AccessToken: token,
	}
}

func (svc *authService) createAccessToken(user *domain.User) (accessToken string, err error) {
	claims := &domain.JwtCustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * svc.config.JWT.TTL).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(svc.config.JWT.AccessSecret))

	return
}

func (svc *authService) Login(c context.Context, req *domain.LoginRequest) (res domain.LoginResponse, err error) {

	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	user, err := svc.userRepo.GetByEmail(ctx, req.Email)

	if err != nil {
		return domain.LoginResponse{}, err
	}

	if user.Email == "" {
		return domain.LoginResponse{}, domain.ErrInvalidCredentials
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return domain.LoginResponse{}, domain.ErrInvalidCredentials
	}

	accessToken, err := svc.createAccessToken(&user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	res = newLoginResponse(accessToken)

	return
}
