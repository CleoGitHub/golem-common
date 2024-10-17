package mvalidator

import (
	"errors"
	"fmt"
	"net/http"
)

// Check if the reference exist in DB
func MIMEType(allowedMIMETypes map[string]bool, bytes []byte, field string) error {
	mimeType := http.DetectContentType(bytes)
	if !allowedMIMETypes[mimeType] {
		out := make(ApiErrors, 1)
		out[0] = ApiError{
			Msg:   fmt.Sprintf("MIME type << %s >> non autorisé", mimeType),
			Field: field,
		}
		return errors.New(out.ToJson())
	}

	return nil
}
