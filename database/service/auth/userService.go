package service

import (
	"errors"

	authDtos "github.com/afolabiolayinka/contact-go/database/dto/auth"
	authModels "github.com/afolabiolayinka/contact-go/database/model/auth"
	"github.com/afolabiolayinka/contact-go/database/repository"
	authRepos "github.com/afolabiolayinka/contact-go/database/repository/auth"
	"github.com/google/uuid"
)

type userService struct {
	userRepository authRepos.UserRepositoryInterface
}

type UserServiceInterface interface {
	Create(userDTO authDtos.UserDTO) (authDtos.UserDTO, error)
	Read(uid uuid.UUID) (authDtos.UserDTO, error)
	ReadAll(pageable repository.Pageable) ([]authDtos.UserDTO, repository.Pagination, error)
	Update(userDTO authDtos.UserDTO) (authDtos.UserDTO, error)
	Delete(userUUID uuid.UUID, uid uuid.UUID) error
	FindByEmail(email string) (authDtos.UserDTO, error)
	Authenticate(user authDtos.UserDTO) (authDtos.UserDTO, error)
}

func NewUserService(userRepository authRepos.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

// ConvertToDTO : converts to DTO
func (service *userService) ConvertToDTO(user authModels.User) (userDTO authDtos.UserDTO) {
	userDTO.UUID = user.UUID
	userDTO.CreatedAt = user.CreatedAt
	userDTO.UpdatedAt = user.UpdatedAt
	userDTO.DeletedAt = user.DeletedAt.Time
	userDTO.Email = user.Email
	userDTO.CreatedByID = user.CreatedByID
	userDTO.UpdatedByID = user.UpdatedByID
	userDTO.DeletedByID = user.DeletedByID

	return userDTO
}

// ConvertToModel : converts to Model
func (service *userService) ConvertToModel(userDTO authDtos.UserDTO) (user authModels.User) {
	user.CreatedAt = userDTO.CreatedAt
	user.UpdatedAt = userDTO.UpdatedAt
	user.DeletedAt.Time = userDTO.DeletedAt
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	user.CreatedByID = userDTO.CreatedByID
	user.UpdatedByID = userDTO.UpdatedByID
	user.DeletedByID = userDTO.DeletedByID

	return user
}

// Create
func (service *userService) Create(userDTO authDtos.UserDTO) (authDtos.UserDTO, error) {

	user := service.ConvertToModel(userDTO)

	newRecord, err := service.userRepository.Create(user)

	return service.ConvertToDTO(newRecord), err

}

// Read
func (service *userService) Read(uid uuid.UUID) (authDtos.UserDTO, error) {

	record, err := service.userRepository.Read(uid)

	return service.ConvertToDTO(record), err
}

// ReadAll
func (service *userService) ReadAll(pageable repository.Pageable) (recordsDto []authDtos.UserDTO, pagination repository.Pagination, err error) {

	records, pagination, err := service.userRepository.ReadAll(pageable)

	for _, record := range records {
		recordsDto = append(recordsDto, service.ConvertToDTO(record))
	}

	return recordsDto, pagination, err
}

// Update
func (service *userService) Update(userDTO authDtos.UserDTO) (authDtos.UserDTO, error) {

	user := service.ConvertToModel(userDTO)

	newRecord, err := service.userRepository.Update(user)

	return service.ConvertToDTO(newRecord), err
}

// Delete
func (service *userService) Delete(userUUID uuid.UUID, uid uuid.UUID) (err error) {
	rtn := service.userRepository.Delete(uid)

	record, _ := service.userRepository.Read(uid)
	record.DeletedByID = userUUID

	service.userRepository.Update(record)

	return rtn
}

// FindByEmail : function
func (service *userService) FindByEmail(email string) (authDtos.UserDTO, error) {
	record, err := service.userRepository.FindByEmail(email)
	if err != nil {
		err = errors.New("email address not found")
	}
	return service.ConvertToDTO(record), err
}

// Authenticate : function for authentication
func (service *userService) Authenticate(userDTO authDtos.UserDTO) (authDtos.UserDTO, error) {
	user := service.ConvertToModel(userDTO)
	newRecord, err := service.userRepository.Authenticate(user)
	return service.ConvertToDTO(newRecord), err
}
