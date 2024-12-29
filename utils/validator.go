package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// Universal function to validate and transfer data using reflection
func ValidateAndTransfer(validatorStruct, modelStruct interface{}) error {
	// Initialize the validator
	validate := validator.New()

	// Validate the Validator struct
	err := validate.Struct(validatorStruct)
	if err != nil {
		// If validation fails, return the error
		return err
	}

	// Use reflection to transfer values if validation is successful
	validatorVal := reflect.ValueOf(validatorStruct).Elem()
	modelVal := reflect.ValueOf(modelStruct).Elem()

	// Ensure both structs are pointers
	if validatorVal.Kind() != reflect.Struct || modelVal.Kind() != reflect.Struct {
		return fmt.Errorf("both input parameters must be pointers to structs")
	}

	// Loop through the fields of the Validator struct and transfer to Model struct
	for i := 0; i < validatorVal.NumField(); i++ {
		fieldName := validatorVal.Type().Field(i).Name
		fieldValue := validatorVal.Field(i)

		// Find the corresponding field in the Model struct by name
		modelField := modelVal.FieldByName(fieldName)

		// Check if the field exists in the Model struct
		if modelField.IsValid() && modelField.CanSet() {
			// Set the value in the Model struct
			modelField.Set(fieldValue)
		}
	}

	return nil
}