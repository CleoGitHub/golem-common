package mvalidator

import (
	"errors"

	"github.com/cleogithub/golem-common/pkg/merror"

	"gorm.io/gorm"
)

func NewReferenceError(relationgName string) error {
	out := make(ApiErrors, 1)
	out[0] = ApiError{
		Msg:   "Cette ressource n'existe pas",
		Field: relationgName,
	}
	return errors.New(out.ToJson())
}

// Check if the reference exist in DB
func Reference(m interface{}, id interface{}, relationgName string, db *gorm.DB) error {
	dbFilter := map[string]interface{}{}
	dbFilter["id"] = id

	query := db.Model(m).Where(dbFilter)
	err := query.First(nil).Error
	//If does not exist in db, return error
	if err == gorm.ErrRecordNotFound {
		out := make(ApiErrors, 1)
		out[0] = ApiError{
			Msg:   "Cette ressource n'existe pas",
			Field: relationgName,
		}
		return errors.New(out.ToJson())
	}

	//If error is nil, field exists in db with the same value
	if err == nil {
		return nil
	}

	//Otherwise db error
	return merror.Stack(err)
}
