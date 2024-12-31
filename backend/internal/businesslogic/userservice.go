package businesslogic

import (
	"backend/internal/database/dao"
	"backend/internal/database/models"
	"backend/pkg/logger"
	"errors"
	"log"
	"net/smtp"
	"os"
	"sync"
	"time"

	"backend/internal/channels"

	"github.com/golang-jwt/jwt/v5"
	emailpkg "github.com/jordan-wright/email"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	instance *UserService
	once     sync.Once
	mutex    sync.Mutex
)

const Secret = "mychatapp"
const subject = "You are caught!"
const body = "\nWelcome to ChatApp!"
const from = "Vipin K <vipin.kunam123@gmail.com>"

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

type email struct {
	smtpHost    string
	smtpPort    string
	senderEmail string
	password    string
	Subject     string
	Body        string
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
	us.SendEmail(user.Email, user.Username)
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

func (us UserService) SendEmail(to, name string) {

	// send mail asynchronous ,let background service take care of status of job
	var mailJob channels.Job
	mailJob = func(errChan chan error, status chan string) {

		err := sendMailtoRegisterUser(to, name)
		if err != nil {
			errChan <- err
		} else {
			status <- "Email sent successfully"
		}

	}

	taskChannel := channels.GetTaskChannel()
	go func() {
		taskChannel <- mailJob
	}()

}

func sendMailtoRegisterUser(to, name string) error {

	smtpHost := os.Getenv("SMPT_HOST")
	smtpPort := os.Getenv("SMPT_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("PASSWORD")

	// Create a new email
	e := emailpkg.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	message := "Hello, " + name + body
	e.Text = []byte(message)

	// Send the email
	err := e.Send(smtpHost+":"+smtpPort, smtp.PlainAuth("", senderEmail, password, smtpHost))
	if err != nil {
		return err
	}

	log.Println("Email sent successfully!")

	return nil

}
