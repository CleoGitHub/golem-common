package mvalidator

import "errors"

var (
	//ErrNotStruct is return when an element is not a struct type when it is supposed to be
	ErrNotStruct = errors.New("not a struct type")

	//ErrNotTypeValidationErrors is return when it is expected to get a validator.ValidationErrors and it is not
	ErrNotTypeValidationErrors = errors.New("error should be of type validator.ValidationErrors")

	//ErrWhereOptClauseCanNotBeNull is return when the WhereOpt is used but the clause is null
	ErrWhereOptClauseCanNotBeNull = errors.New("clause in WhereOpt can not be null")

	//ErrWhereOptValueCanNotBeNull is return when the WhereOpt is used but the value is null
	ErrWhereOptValueCanNotBeNull = errors.New("value in WhereOpt can not be null")
)
