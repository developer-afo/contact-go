package app

import (
	"strings"

	databaseModule "github.com/afolabiolayinka/contact-go/database"
	appModel "github.com/afolabiolayinka/contact-go/database/model/app"
	"github.com/afolabiolayinka/contact-go/database/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContactRepositoryInterface interface {
	Create(contact appModel.Contact) (appModel.Contact, error)
	Read(uid uuid.UUID) (appModel.Contact, error)
	ReadAll(pageable repository.Pageable) ([]appModel.Contact, repository.Pagination, error)
	Update(contact appModel.Contact) (appModel.Contact, error)
	Delete(uid uuid.UUID) error
}

type contactRepository struct {
	database databaseModule.DatabaseInterface
}

// Constructor : function
func NewContactRepostiory(database databaseModule.DatabaseInterface) ContactRepositoryInterface {
	return &contactRepository{database: database}
}

// Create : function
func (repository *contactRepository) Create(contact appModel.Contact) (appModel.Contact, error) {
	contact.Model.Prepare()

	err := repository.database.Connection().Create(&contact).Error
	if err != nil {
		return contact, err
	}

	return contact, nil
}

// Read : function
func (repository *contactRepository) Read(uid uuid.UUID) (contact appModel.Contact, err error) {
	err = repository.database.Connection().Model(&appModel.Contact{}).Where("uuid = ?", uid).First(&contact).Error

	if err != nil {
		return contact, err
	}

	return contact, nil
}

// ReadAll : function
func (repository *contactRepository) ReadAll(pageable repository.Pageable) (rows []appModel.Contact, pagination repository.Pagination, err error) {
	var contact appModel.Contact
	pagination.TotalPages = 1
	pagination.TotalItems = 0
	pagination.CurrentPage = int64(pageable.Page)

	var result *gorm.DB
	var errCount error

	offset := (pageable.Page - 1) * pageable.Size
	searchQuery := repository.database.Connection().Model(&appModel.Contact{})

	if len(strings.TrimSpace(pageable.Search)) > 0 {
		searchQuery.Where("first_name LIKE ?", "%"+strings.ToLower(pageable.Search)+"%")
	}
	errCount = searchQuery.Count(&pagination.TotalItems).Error
	paginationQuery := searchQuery.Limit(pageable.Size).Offset(offset).Order(pageable.SortBy + " " + pageable.SortDirection)

	result = paginationQuery.Model(&appModel.Contact{}).Where(contact).Find(&rows)

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
func (repository *contactRepository) Update(contact appModel.Contact) (appModel.Contact, error) {

	var checkRow appModel.Contact

	err := repository.database.Connection().Model(&appModel.Contact{}).Where("uuid = ? ", contact.UUID.String()).First(&checkRow).Error

	if err != nil {
		return checkRow, err
	}

	err = repository.database.Connection().Model(&checkRow).Updates(contact).Error
	if err != nil {
		return contact, err
	}
	return contact, nil
}

// Delete : function
func (repository *contactRepository) Delete(uuid uuid.UUID) (err error) {
	var contact appModel.Contact
	err = repository.database.Connection().Model(&appModel.Contact{}).Where("uuid = ? ", uuid).First(&contact).Error

	if err != nil {
		return err
	}

	err = repository.database.Connection().Delete(contact).Error
	if err != nil {
		return err
	}

	return nil
}
