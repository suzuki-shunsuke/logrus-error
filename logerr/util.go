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
		return appendFields(getFields(errors.Unwrap(e)), e.GetLogrusFields())
	}
	return nil
}

func appendFields(fields1, fields2 logrus.Fields) logrus.Fields {
	fields := make(logrus.Fields, len(fields1)+len(fields2))
	for k, v := range fields1 {
		fields[k] = v
	}
	for k, v := range fields2 {
		fields[k] = v
	}
	return fields
}
