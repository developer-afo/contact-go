package auth

import (
	"errors"
	"strings"

	databaseModule "github.com/afolabiolayinka/contact-go/database"
	authModel "github.com/afolabiolayinka/contact-go/database/model/auth"
	"github.com/afolabiolayinka/contact-go/database/repository"
	"github.com/afolabiolayinka/contact-go/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(user authModel.User) (authModel.User, error)
	Read(uid uuid.UUID) (authModel.User, error)
	ReadAll(pageable repository.Pageable) ([]authModel.User, repository.Pagination, error)
	Update(user authModel.User) (authModel.User, error)
	Delete(uid uuid.UUID) error
	FindByEmail(email string) (authModel.User, error)
	Authenticate(user authModel.User) (authModel.User, error)
}

type userRepository struct {
	database databaseModule.DatabaseInterface
}

// Constructor : function
func NewUserRepostiory(database databaseModule.DatabaseInterface) UserRepositoryInterface {
	return &userRepository{database: database}
}

// Create : function
func (repository *userRepository) Create(user authModel.User) (authModel.User, error) {
	user.Model.Prepare()

	err := repository.database.Connection().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Read : function
func (repository *userRepository) Read(uid uuid.UUID) (user authModel.User, err error) {
	err = repository.database.Connection().Model(&authModel.User{}).Where("uuid = ?", uid).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// ReadAll : function
func (repository *userRepository) ReadAll(pageable repository.Pageable) (rows []authModel.User, pagination repository.Pagination, err error) {
	var user authModel.User
	pagination.TotalPages = 1
	pagination.TotalItems = 0
	pagination.CurrentPage = int64(pageable.Page)

	var result *gorm.DB
	var errCount error

	offset := (pageable.Page - 1) * pageable.Size
	searchQuery := repository.database.Connection().Model(&authModel.User{})

	if len(strings.TrimSpace(pageable.Search)) > 0 {
		searchQuery.Where("email LIKE ?", "%"+strings.ToLower(pageable.Search)+"%")
	}
	errCount = searchQuery.Count(&pagination.TotalItems).Error
	paginationQuery := searchQuery.Limit(pageable.Size).Offset(offset).Order(pageable.SortBy + " " + pageable.SortDirection)

	result = paginationQuery.Model(&authModel.User{}).Where(user).Find(&rows)

	if result.Error != nil {
		msg := result.Error
		return nil, pagination, msg
	}

	if errCount != nil {
		return nil, pagination, errCount
	}

	pagination.TotalPages = pagination.TotalItems / int64(pageable.Size)

	return rows, pagination, nil

}

// Update : function
func (repository *userRepository) Update(user authModel.User) (authModel.User, error) {

	var checkRow authModel.User

	err := repository.database.Connection().Model(&authModel.User{}).Where("uuid = ? ", user.UUID.String()).First(&checkRow).Error

	if err != nil {
		return checkRow, err
	}

	err = repository.database.Connection().Model(&checkRow).Updates(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// Delete : function
func (repository *userRepository) Delete(uuid uuid.UUID) (err error) {
	var user authModel.User
	err = repository.database.Connection().Model(&authModel.User{}).Where("uuid = ? ", uuid).First(&user).Error

	if err != nil {
		return err
	}

	err = repository.database.Connection().Delete(user).Error
	if err != nil {
		return err
	}

	return nil
}

// FindByEmail : function
func (repository *userRepository) FindByEmail(email string) (row authModel.User, err error) {
	err = repository.database.Connection().Model(&authModel.User{}).Where("email = ?", email).First(&row).Error
	if err != nil {
		return row, err
	}
	return row, nil
}

// Authenticate : function
func (repository *userRepository) Authenticate(user authModel.User) (checkUser authModel.User, err error) {
	var rtn bool

	//checking username
	err = repository.database.Connection().Model(&authModel.User{}).Where("email = ?", user.Email).First(&checkUser).Error

	if err != nil {
		return checkUser, errors.New("user not found")
	}

	//checking password hash
	rtn, err = helper.ComparePassword(user.Password, checkUser.Password)

	if !rtn || err != nil {
		return user, errors.New("password do not match")
	}

	//loadUser
	repository.database.Connection().Where("uuid = ?", checkUser.UUID).First(&user)

	return user, nil
}
