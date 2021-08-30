package logerr

import (
	"github.com/sirupsen/logrus"
)

// AppendFields merges fields1 and fields2 and returns a new fields.
// fields of fields2 overwrites fields of fields1.
func AppendFields(fields1, fields2 logrus.Fields) logrus.Fields {
	fields := make(logrus.Fields, len(fields1)+len(fields2))
	for k, v := range fields1 {
		fields[k] = v
	}
	for k, v := range fields2 {
		fields[k] = v
	}
	return fields
}

// WithError appends err to entry and returns new entry.
func WithError(entry *logrus.Entry, err error) *logrus.Entry {
	if entry == nil {
		if err == nil {
			return logrus.NewEntry(logrus.New())
		}
		return logrus.WithError(err).WithFields(getFields(err))
	}
	if err == nil {
		return entry
	}
	return entry.WithError(err).WithFields(getFields(err))
}

// LogrusError is an error which has logrus.Fields.
// GetLogrusFields returns fields but doesn't return unwrapped errors fields.
type LogrusError interface {
	error
	GetLogrusFields() logrus.Fields
}

type logrusError struct {
	err    error
	fields logrus.Fields
}

// WithFields appends fields to err and returns new error.
// If err is nil, nil is returned.
func WithFields(err error, fields logrus.Fields) error {
	if err == nil {
		return nil
	}
	return &logrusError{
		err:    err,
		fields: AppendFields(getFields(err), fields),
	}
}

func (e *logrusError) GetLogrusFields() logrus.Fields {
	if e == nil {
		return nil
	}
	return e.fields
}

// Error returns an error message.
func (e *logrusError) Error() string {
	if e == nil {
		return ""
	}
	return e.err.Error()
}

func (e *logrusError) Unwrap() error {
	if e == nil || e.err == nil {
		return nil
	}
	return e.err
}
