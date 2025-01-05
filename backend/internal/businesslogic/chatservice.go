package businesslogic

import (
	"backend/internal/database/dao"
	"backend/internal/database/models"
	"backend/pkg/logger"
	"errors"
)

var chatInstance *ChatService

type ChatService struct {
	log logger.Logger
}

func GetChatServiceInstance() *ChatService {
	once.Do(func() {
		chatInstance = &ChatService{log: logger.GetLogrusLogger()}
	})
	return chatInstance
}

func (cs ChatService) AddUserToInteractedListOfCurrentUser(currentUser *models.User, toBeAddedUserEmailId string) error {

	userDao := dao.GetUserDaoInstance()

	// Check for user already signedup

	toBeAddedUserFromDb, err := userDao.GetByEmail(toBeAddedUserEmailId)

	if err != nil {
		return err
	}

	if toBeAddedUserFromDb == nil || toBeAddedUserFromDb.UserID == 0 {
		return errors.New("User does not exist")
	}

	currentUser.AddInteractedUser(int64(toBeAddedUserFromDb.UserID))

	err = userDao.Update(currentUser)
	if err != nil {
		cs.log.Error("error updating current user while adding interacted user ", currentUser.Email, err)
		return err
	}

	return nil

}

func (cs ChatService) RemoveUserFromInteractedListOfCurrentUser(currentUser *models.User, toBeRemovedUserEmailId string) error {

	userDao := dao.GetUserDaoInstance()

	// Check for user already signedup

	toBeRemovedUserFromDb, err := userDao.GetByEmail(toBeRemovedUserEmailId)

	if err != nil {
		return err
	}

	if toBeRemovedUserFromDb == nil || toBeRemovedUserFromDb.UserID == 0 {
		return errors.New("User does not exist")
	}
	currentUser.RemoveInteractedUser(int64(toBeRemovedUserFromDb.UserID))
	err = userDao.Update(currentUser)

	if err != nil {
		cs.log.Error("error updating current user while removing interacted user ", currentUser.Email, err)
		return err
	}

	return nil
}

func (cs ChatService) GetInteractedUsers(currentUser *models.User) ([]*models.User, error) {
	userDao := dao.GetUserDaoInstance()
	interactedUsers, err := currentUser.GetInteractedUsers()
	if err != nil {
		return nil, err
	}
	interactedUsersList := make([]*models.User, len(interactedUsers))
	for _, user := range interactedUsers {

		userFromDb, err := userDao.GetByID(uint(user))
		if err != nil {
			return nil, err
		}
		interactedUsersList = append(interactedUsersList, userFromDb)

	}

	return interactedUsersList, nil
}
