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
type InteractedUser struct {
	UserEmailId string `json:"emailId" binding:"required"`
}

type PaginationInfo struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"size" binding:"required"`
}

func GetChatServiceInstance() *ChatService {
	once.Do(func() {
		chatInstance = &ChatService{log: logger.GetLogrusLogger()}
	})
	return chatInstance
}

func (cs ChatService) AddUserToInteractedListOfCurrentUser(currentUser *models.User, toBeAddedUser InteractedUser) error {

	userDao := dao.GetUserDaoInstance()

	// Check for user already signedup

	toBeAddedUserFromDb, err := userDao.GetByEmail(toBeAddedUser.UserEmailId)

	if err != nil {
		return err
	}

	if toBeAddedUserFromDb == nil || toBeAddedUserFromDb.UserID == 0 {
		return errors.New("User does not exist")
	}

	userInteractionDao := dao.GetUserInteractionDAO()

	yes, err := userInteractionDao.InteractionExists(currentUser.UserID, toBeAddedUserFromDb.UserID)

	if err != nil {
		return err
	}

	if !yes {
		err = userInteractionDao.InsertInteraction(currentUser.UserID, toBeAddedUserFromDb.UserID)
		if err != nil {
			return err
		}
	}
	yes, err = userInteractionDao.InteractionExists(toBeAddedUserFromDb.UserID, currentUser.UserID)

	if err != nil {
		return err
	}

	if !yes {
		err = userInteractionDao.InsertInteraction(toBeAddedUserFromDb.UserID, currentUser.UserID)
		if err != nil {
			return err
		}
	}

	return nil

}

func (cs ChatService) RemoveUserFromInteractedListOfCurrentUser(currentUser *models.User, toBeRemovedUser InteractedUser) error {

	userDao := dao.GetUserDaoInstance()

	// Check for user already signedup

	toBeRemovedUserFromDb, err := userDao.GetByEmail(toBeRemovedUser.UserEmailId)

	if err != nil {
		return err
	}

	if toBeRemovedUserFromDb == nil || toBeRemovedUserFromDb.UserID == 0 {
		return errors.New("User does not exist")
	}

	userInteractionDao := dao.GetUserInteractionDAO()

	yes, err := userInteractionDao.InteractionExists(currentUser.UserID, toBeRemovedUserFromDb.UserID)

	if err != nil {
		return err
	}
	if yes {
		err = userInteractionDao.DeleteInteraction(currentUser.UserID, toBeRemovedUserFromDb.UserID)
		if err != nil {
			return err
		}
	}

	yes, err = userInteractionDao.InteractionExists(toBeRemovedUserFromDb.UserID, currentUser.UserID)

	if err != nil {
		return err
	}

	if yes {
		err = userInteractionDao.DeleteInteraction(toBeRemovedUserFromDb.UserID, currentUser.UserID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs ChatService) GetInteractedUsers(currentUser *models.User, paginationInfo PaginationInfo) ([]models.UserInteraction, int64, error) {
	userInteractionDao := dao.GetUserInteractionDAO()
	interactedUsers, totalInteractedUsers, err := userInteractionDao.GetInteractedUsers(currentUser.UserID, paginationInfo.Page, paginationInfo.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return interactedUsers, totalInteractedUsers, nil
}
