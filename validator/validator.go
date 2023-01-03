package validator

import (
	"errors"
	"reflect"

	"github.com/afolabiolayinka/contact-go/database"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Validator[T any] struct {
}

func (validator *Validator[T]) ValidateDBUnique(structure T, tableName string, uniqueField string, parameters map[string]interface{}) validation.RuleFunc {
	//db := database.StartDatabaseClient().Connection()
	db := database.DatabaseFacade

	result := map[string]interface{}{}

	// Reflect the structure interface
	e := reflect.ValueOf(&structure).Elem()

	parentID := e.FieldByName("UUID").Interface().(uuid.UUID)

	//log.Println("Structure ID :", parentID)

	return func(value interface{}) error {
		query := db.Table(tableName).Where(uniqueField+" = ?", value)

		// For update; check if the structure contain have an ID
		if parentID != uuid.Nil {
			query = query.Where("uuid != ?", parentID.String())
		}

		for key, parameter := range parameters {
			param := e.FieldByName(key).Interface()
			query = query.Where(parameter.(string)+" = ?", param)
		}

		rows := query.Take(&result)

		// Count the number of rows affected
		if rows.RowsAffected > 0 {
			return errors.New("value already exist")
		}
		return nil
	}
}
