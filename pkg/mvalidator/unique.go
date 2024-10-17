package mvalidator

import (
	"errors"

	"github.com/cleogithub/golem-common/pkg/merror"

	"gorm.io/gorm"
)

const UNIQUE_ERROR_MESSAGE = "Ce champ est unique et une resource possède déjà cette valeur"

type UniqueOpt interface {
	configureUniqueContext(opts *uniqueOpts) error
}

type uniqueOpts struct {
	// If ID set, the entity with ID as id will not be taken in account
	ID string

	//If set the where clause is applied to add a uniqueness condition
	WhereClause string
	WhereValue  string

	//If set, overide default error message
	Msg string
}

// If ID set, the entity with ID as id will not be taken in account
type MessageOpt string

func (msg MessageOpt) configureUniqueContext(opts *uniqueOpts) error {
	opts.Msg = string(msg)
	return nil
}

// If ID set, the entity with ID as id will not be taken in account
type IDOpt string

func (id IDOpt) configureUniqueContext(opts *uniqueOpts) error {
	opts.ID = string(id)
	return nil
}

// If set the where clause is applied to add a uniqueness condition
type WhereOpt struct {
	Clause string
	Value  string
}

func (opt WhereOpt) configureUniqueContext(opts *uniqueOpts) error {
	if opt.Clause == "" {
		return ErrWhereOptClauseCanNotBeNull
	}
	if opt.Value == "" {
		return ErrWhereOptValueCanNotBeNull
	}
	opts.WhereClause = opt.Clause
	opts.WhereValue = opt.Value
	return nil
}

func NewUniqueError(field string, opts ...UniqueOpt) error {
	uniqueContext := &uniqueOpts{}
	for _, opt := range opts {
		err := opt.configureUniqueContext(uniqueContext)
		if err != nil {
			return merror.Stack(err)
		}
	}

	msg := UNIQUE_ERROR_MESSAGE
	if uniqueContext.Msg != "" {
		msg = uniqueContext.Msg
	}
	out := make(ApiErrors, 1)
	out[0] = ApiError{
		Msg:   msg,
		Field: field,
	}
	return errors.New(out.ToJson())
}

// Check if the field is unique in DB
func Unique(m interface{}, field string, value string, db *gorm.DB, opts ...UniqueOpt) error {
	uniqueContext := &uniqueOpts{}
	for _, opt := range opts {
		err := opt.configureUniqueContext(uniqueContext)
		if err != nil {
			return merror.Stack(err)
		}
	}

	dbFilter := map[string]interface{}{}
	dbFilter[field] = value

	query := db.Model(m).Where(dbFilter)
	if uniqueContext.ID != "" {
		query = query.Where("id != ?", uniqueContext.ID)
	}

	if uniqueContext.WhereClause != "" && uniqueContext.WhereValue != "" {
		query = query.Where(uniqueContext.WhereClause, uniqueContext.WhereValue)
	}

	err := query.First(m).Error
	//If does not exist in db, return no error
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	//If error is nil, field exists in db with the same value
	if err == nil {
		msg := UNIQUE_ERROR_MESSAGE
		if uniqueContext.Msg != "" {
			msg = uniqueContext.Msg
		}
		out := make(ApiErrors, 1)
		out[0] = ApiError{
			Msg:   msg,
			Field: field,
		}
		return errors.New(out.ToJson())
	}

	//Otherwise db error
	return merror.Stack(err)
}
