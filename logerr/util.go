package logerr

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// getFields extracts fields from err.
// If err wraps other errors, getFields merges these fields.
func getFields(err error) logrus.Fields {
	var e LogrusError
	if errors.As(err, &e) {
		return AppendFields(getFields(errors.Unwrap(e)), e.GetLogrusFields())
	}
	return nil
}
