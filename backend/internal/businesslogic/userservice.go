package businesslogic

import (
	"backend/internal/database/dao"
	"backend/internal/database/models"
	"backend/pkg/logger"
	"errors"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	instance *UserService
	once     sync.Once
	mutex    sync.Mutex
)

const Secret = "mychatapp"

type RegisterUser struct {
	USERNAME string `json:"username" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	PHONENO  string `json:"phoneno" binding:"omitempty"`
}

type LOGIN struct {
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	PHONENO  string `json:"phoneno" binding:"omitempty"`
}

type UserService struct {
	log logger.Logger
}

// GetInstance provides access to the singleton UserService instance
func GetUserServiceInstance() *UserService {
	once.Do(func() {
		instance = &UserService{log: logger.GetLogrusLogger()}
	})
	return instance
}
func (us UserService) RegisterUserForApp(toBeRegisteredUser RegisterUser) error {

	userDao := dao.GetUserDaoInstance()
	user, err := userDao.GetByEmail(toBeRegisteredUser.EMAIL)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		us.log.Error("error getting user", err)
		return err
	}

	if user != nil && user.UserID != 0 {
		return errors.New("user already exist")
	}

	// Create Password Hash

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(toBeRegisteredUser.PASSWORD), bcrypt.DefaultCost)

	if err != nil {
		us.log.Error("error encrypting password", err)
		return err
	}

	user = &models.User{Username: toBeRegisteredUser.USERNAME,
		Email:        toBeRegisteredUser.EMAIL,
		PasswordHash: string(passwordHash),
		PhoneNumber:  &toBeRegisteredUser.PHONENO,
	}

	err = userDao.Create(user)

	if err != nil {
		us.log.Error("error creating user", err)
		return err
	}
	return nil

}
func (us UserService) LoginUserForApp(toBeLoggedinUser LOGIN) (string, error) {

	userDao := dao.GetUserDaoInstance()

	user, err := userDao.GetByEmail(toBeLoggedinUser.EMAIL)

	if err != nil {
		us.log.Error("error getting user", err)
		return "", err
	}

	if user.UserID == 0 {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(toBeLoggedinUser.PASSWORD))

	if err != nil {
		us.log.Error("username or password invalid", err)
		return "", errors.New("username or password invalid")
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(Secret))

	if err != nil {
		us.log.Error("error generating jwt token", err)
		return "", errors.New("error generating jwt token")
	}

	return token, nil

}
