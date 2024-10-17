package mvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/cleogithub/golem-common/pkg/merror"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string
	Msg   string
}
type ApiErrors []ApiError

var mvalidator *validator.Validate

func initValitor() {
	if mvalidator == nil {
		mvalidator = validator.New()
	}
}

func Struct(s interface{}) error {
	initValitor()

	err := mvalidator.Struct(s)
	if err != nil {
		return customMessage(err, s)
	}

	return nil
}

func customMessage(err error, s interface{}) error {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return ErrNotTypeValidationErrors
	}

	var keyToJson map[string]string = make(map[string]string)
	e := addStructToJsonMap("", "", &keyToJson, s)
	if e != nil {
		return merror.Stack(e)
	}

	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	//Create errors
	out := make(ApiErrors, len(ve))
	for i, fe := range ve {
		fieldName := keyToJson[fe.Namespace()]
		if fieldName == "" {
			fieldName = fe.Namespace()
		}
		out[i] = ApiError{fieldName, msgForTag(fe)}
	}

	return errors.New(out.ToJson())
}

func msgForTag(field validator.FieldError) string {
	switch field.Tag() {
	case "required":
		return "Ce champ est obligatoire"
	case "email":
		return "Email invalide"
	case "hexcolor":
		return "Cette valeur doit être une couleur en hexadécimal"
	case "gt":
		if field.Type().String() == "string" {
			return "La longuer doit être supérieure à " + field.Param()
		} else {
			return "Ce champ doit être supérieure à " + field.Param()
		}
	case "gte":
		if field.Type().String() == "string" {
			return "La longuer doit être supérieure ou égal à " + field.Param()
		} else {
			return "Ce champ doit être supérieure ou égal à " + field.Param()
		}
	case "lte":
		if field.Type().String() == "string" {
			return "La longuer doit être inférieure ou égal à " + field.Param()
		} else {
			return "Ce champ doit être inférieure ou égal à " + field.Param()
		}
	}
	return field.Error()
}

func (a ApiError) ToJson() string {
	return fmt.Sprintf("{\"Field\":\"%s\",\"Msg\":\"%s\"}", a.Field, a.Msg)
}

func (as ApiErrors) ToJson() string {
	message := "["
	for idx, a := range as {
		if idx != 0 {
			message += ","
		}
		message += a.ToJson()
	}

	message += "]"

	return message
}

func IsValidatorErrors(err error) bool {
	return json.Unmarshal([]byte(err.Error()), &ApiErrors{}) == nil
}

func addStructToJsonMap(keyPrefix string, jsonPrefix string, keyToJson *map[string]string, s interface{}) error {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	vt := reflect.ValueOf(s)

	if keyPrefix == "" {
		keyPrefix = rt.Name()
	}

	if jsonPrefix != "" {
		jsonPrefix += "."
	}

	for i := 0; i < rt.NumField(); i++ {
		typeField := rt.Field(i)
		valueField := vt.Field(i)

		currentKeyPrefix := keyPrefix + "." + typeField.Name

		jsonName := typeField.Name
		v := strings.Split(typeField.Tag.Get("json"), ",")[0]
		if v != "" && v != "-" {
			jsonName = v
		}
		currentJsonPrefix := jsonPrefix + jsonName

		i := valueField.Interface()
		k := reflect.TypeOf(i).Kind()

		switch k {
		case reflect.Slice:
			for idx := 0; idx < valueField.Len(); idx++ {
				currentKeyPrefix += fmt.Sprintf("[%d]", idx)
				currentJsonPrefix += fmt.Sprintf("[%d]", idx)

				v := valueField.Index(idx)
				vrt := reflect.TypeOf(v.Interface())
				if vrt.Kind() == reflect.Struct {
					addStructToJsonMap(currentKeyPrefix, currentJsonPrefix, keyToJson, v.Interface())
				} else {
					(*keyToJson)[currentKeyPrefix] = currentJsonPrefix
				}
			}
		case reflect.Struct:
			if err := addStructToJsonMap(currentKeyPrefix, currentJsonPrefix, keyToJson, i); err != nil {
				return merror.Stack(err)
			}
		default:
			(*keyToJson)[currentKeyPrefix] = currentJsonPrefix
		}
	}

	return nil
}

func AsMvalidatorError(field string, msg string) error {
	err := ApiError{
		Msg:   msg,
		Field: field,
	}
	return errors.New(ApiErrors{err}.ToJson())
}
