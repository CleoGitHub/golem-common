package mvalidator

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/cleogithub/golem-common/pkg/merror"
)

// Check if the reference exist in DB
func AtLeastOne(o interface{}, fields ...string) error {
	bs, err := json.Marshal(o)
	if err != nil {
		return merror.Stack(err)
	}

	m := map[string]interface{}{}
	err = json.Unmarshal(bs, &m)
	if err != nil {
		return merror.Stack(err)
	}

	fieldsStr := strings.Join(fields, ",")
	out := make(ApiErrors, len(fields))

	for idx, field := range fields {
		if _, exist := m[field]; exist && m[field] != "" {
			return nil
		}
		out[idx] = ApiError{
			Msg:   "Au moins, une des valeurs ne doit pas être nul parmi : " + fieldsStr,
			Field: field,
		}
	}

	return errors.New(out.ToJson())
}
