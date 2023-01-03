package auth

import (
	"encoding/json"
	"log"

	authDtos "github.com/afolabiolayinka/contact-go/database/dto/auth"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserValidator struct {
	Validator[authDtos.UserDTO]
}

func (validator *UserValidator) Validate(userDto authDtos.UserDTO) (map[string]interface{}, error) {

	err := validation.ValidateStruct(&userDto,
		validation.Field(&userDto.Email, validation.Required, validation.Length(2, 128), validation.By(validator.ValidateDBUnique(userDto, "auth_users", "email", map[string]interface{}{}))),
		validation.Field(&userDto.Password, validation.Required, validation.Length(6, 128)),
	)

	if err != nil {
		if e, ok := err.(validation.InternalError); ok {
			log.Println(e.InternalError())
			return nil, nil
		}

		var dat map[string]interface{}
		m, _ := json.Marshal(err)

		json.Unmarshal(m, &dat)
		return dat, err
	}

	return nil, nil
}
