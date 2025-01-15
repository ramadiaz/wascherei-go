package services

import (
	"net/http"
	"os"
	"time"
	"wascherei-go/api/users/dto"
	"wascherei-go/api/users/repositories"
	"wascherei-go/pkg/exceptions"
	"wascherei-go/pkg/helpers"
	"wascherei-go/pkg/mapper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.UserInput) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		return err
	}

	dataModel := mapper.MapUserInputToModel(data)
	dataModel.UUID = uuid.NewString()
	dataModel.HashedPassword = hashedPassword

	err = s.repo.Create(ctx, s.DB, dataModel)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) Login(ctx *gin.Context, data dto.Login) (*string, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	userData, err := s.repo.FindByUsername(ctx, s.DB, data.Username)
	if err != nil {
		return nil, err
	}

	err = helpers.CheckPasswordHash(data.Password, userData.HashedPassword)
	if err != nil {
		return nil, err
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["uuid"] = userData.UUID
	claims["username"] = userData.Username
	claims["business_name"] = userData.BusinessName
	claims["business_owner"] = userData.BusinessOwner
	claims["business_phone"] = userData.BusinessPhone
	claims["business_address"] = userData.BusinessAddress
	claims["business_logo"] = userData.BusinessLogo
	claims["slogan"] = userData.Slogan

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(JWT_SECRET)
	tokenString, signErr := token.SignedString(secretKey)
	if signErr != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}
