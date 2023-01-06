package app

import (
	"encoding/json"
	"log"

	appDtos "github.com/afolabiolayinka/contact-go/database/dto/app"
	"github.com/afolabiolayinka/contact-go/validator"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ContactValidator struct {
	validator.Validator[appDtos.ContactDTO]
}

func (validator *ContactValidator) Validate(contactDto appDtos.ContactDTO) (map[string]interface{}, error) {

	err := validation.ValidateStruct(&contactDto,
		validation.Field(&contactDto.FirstName, validation.Required, validation.Length(1, 32)),
		validation.Field(&contactDto.LastName, validation.Required, validation.Length(1, 32)),
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
