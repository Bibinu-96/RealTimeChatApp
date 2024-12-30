package businesslogic

import (
	"backend/internal/database/dao"
	"backend/internal/database/models"
	"backend/pkg/logger"
	"errors"
	"sync"
)

var (
	instance *UserService
	once     sync.Once
	mutex    sync.Mutex
)

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

	user := models.User{Username: toBeRegisteredUser.USERNAME,
		Email:        toBeRegisteredUser.EMAIL,
		PasswordHash: toBeRegisteredUser.PASSWORD,
		PhoneNumber:  &toBeRegisteredUser.PHONENO,
	}
	err := userDao.Create(&user)

	if err != nil {
		us.log.Error("error getting user", err)
		return err
	}
	return nil

}
func (us UserService) LoginUserForApp(toBeLoggedinUser LOGIN) error {

	userDao := dao.GetUserDaoInstance()

	user, err := userDao.GetByEmail(toBeLoggedinUser.EMAIL)

	if err != nil {
		us.log.Error("error getting user", err)
		return err
	}
	isPasswordCorrect := toBeLoggedinUser.PASSWORD == user.PasswordHash

	if !isPasswordCorrect {
		err = errors.New("userid or password does not match")
		us.log.Error(err.Error())
		return err
	}

	return nil

}
