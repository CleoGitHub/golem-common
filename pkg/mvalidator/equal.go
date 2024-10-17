package mvalidator

import (
	"errors"
)

// Check if the field is unique in DB
func Equal(fieldOne string, fieldTwo string, valueOne string, valueTwo string) error {
	if valueOne != valueTwo {
		out := make(ApiErrors, 2)
		out[0] = ApiError{
			Msg:   "Ce champ doit être égal à " + fieldTwo,
			Field: fieldOne,
		}
		out[1] = ApiError{
			Msg:   "Ce champ doit être égal à " + fieldOne,
			Field: fieldTwo,
		}
		return errors.New(out.ToJson())
	}

	return nil
}
