package service

import (
	appDtos "github.com/afolabiolayinka/contact-go/database/dto/app"
	appModels "github.com/afolabiolayinka/contact-go/database/model/app"
	"github.com/afolabiolayinka/contact-go/database/repository"
	appRepositories "github.com/afolabiolayinka/contact-go/database/repository/app"
	"github.com/google/uuid"
)

type contactService struct {
	contactRepository appRepositories.ContactRepositoryInterface
}

type ContactServiceInterface interface {
	Create(contactDTO appDtos.ContactDTO) (appDtos.ContactDTO, error)
	Read(uid uuid.UUID) (appDtos.ContactDTO, error)
	ReadAll(pageable repository.Pageable) ([]appDtos.ContactDTO, repository.Pagination, error)
	Update(contactDTO appDtos.ContactDTO) (appDtos.ContactDTO, error)
	Delete(userUUID uuid.UUID, uid uuid.UUID) error
}

func NewContactService(contactRepository appRepositories.ContactRepositoryInterface) ContactServiceInterface {
	return &contactService{
		contactRepository: contactRepository,
	}
}

// ConvertToDTO : converts to DTO
func (service *contactService) ConvertToDTO(contact appModels.Contact) (contactDTO appDtos.ContactDTO) {
	contactDTO.UUID = contact.UUID
	contactDTO.CreatedAt = contact.CreatedAt
	contactDTO.UpdatedAt = contact.UpdatedAt
	contactDTO.DeletedAt = contact.DeletedAt.Time
	contactDTO.FirstName = contact.FirstName
	contactDTO.LastName = contact.LastName
	contactDTO.CreatedByID = contact.CreatedByID
	contactDTO.UpdatedByID = contact.UpdatedByID
	contactDTO.DeletedByID = contact.DeletedByID

	return contactDTO
}

// ConvertToModel : converts to Model
func (service *contactService) ConvertToModel(contactDTO appDtos.ContactDTO) (contact appModels.Contact) {
	contact.CreatedAt = contactDTO.CreatedAt
	contact.UpdatedAt = contactDTO.UpdatedAt
	contact.DeletedAt.Time = contactDTO.DeletedAt
	contact.FirstName = contactDTO.FirstName
	contact.LastName = contactDTO.LastName
	contact.CreatedByID = contactDTO.CreatedByID
	contact.UpdatedByID = contactDTO.UpdatedByID
	contact.DeletedByID = contactDTO.DeletedByID

	return contact
}

// Create
func (service *contactService) Create(contactDTO appDtos.ContactDTO) (appDtos.ContactDTO, error) {

	contact := service.ConvertToModel(contactDTO)

	newRecord, err := service.contactRepository.Create(contact)

	return service.ConvertToDTO(newRecord), err

}

// Read
func (service *contactService) Read(uid uuid.UUID) (appDtos.ContactDTO, error) {

	record, err := service.contactRepository.Read(uid)

	return service.ConvertToDTO(record), err
}

// ReadAll
func (service *contactService) ReadAll(pageable repository.Pageable) (recordsDto []appDtos.ContactDTO, pagination repository.Pagination, err error) {

	records, pagination, err := service.contactRepository.ReadAll(pageable)

	for _, record := range records {
		recordsDto = append(recordsDto, service.ConvertToDTO(record))
	}

	return recordsDto, pagination, err
}

// Update
func (service *contactService) Update(contactDTO appDtos.ContactDTO) (appDtos.ContactDTO, error) {

	user := service.ConvertToModel(contactDTO)

	newRecord, err := service.contactRepository.Update(user)

	return service.ConvertToDTO(newRecord), err
}

// Delete
func (service *contactService) Delete(userUUID uuid.UUID, uid uuid.UUID) (err error) {
	rtn := service.contactRepository.Delete(uid)

	record, _ := service.contactRepository.Read(uid)
	record.DeletedByID = userUUID

	service.contactRepository.Update(record)

	return rtn
}
