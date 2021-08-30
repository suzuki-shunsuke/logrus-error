package logerr

import (
	"errors"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_logrusError_Error(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		err   *logrusError
		exp   string
	}{
		{
			title: "simple",
			err: &logrusError{
				err: errors.New("foo"),
			},
			exp: "foo",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			act := d.err.Error()
			if d.exp != act {
				t.Fatalf("wanted %s, got %s", d.exp, act)
			}
		})
	}
}

func Test_logrusError_Unwrap(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		err   *logrusError
		exp   error
	}{
		{
			title: "nil",
			err:   nil,
			exp:   nil,
		},
		{
			title: "err.err is nil",
			err: &logrusError{
				err: nil,
			},
			exp: nil,
		},
		{
			title: "normal",
			err: &logrusError{
				err: errors.New("foo"),
			},
			exp: errors.New("foo"),
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			act := d.err.Unwrap()
			if !reflect.DeepEqual(d.exp, act) {
				t.Fatalf("wanted %+v, got %+v", d.exp, act)
			}
		})
	}
}

func Test_logrusError_GetLogrusFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		err   *logrusError
		exp   logrus.Fields
	}{
		{
			title: "nil",
			err:   nil,
			exp:   nil,
		},
		{
			title: "fields is nil",
			err:   &logrusError{},
			exp:   nil,
		},
		{
			title: "normal",
			err: &logrusError{
				fields: logrus.Fields{
					"foo": "foo",
				},
			},
			exp: logrus.Fields{
				"foo": "foo",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			act := d.err.GetLogrusFields()
			if !reflect.DeepEqual(d.exp, act) {
				t.Fatalf("wanted %+v, got %+v", d.exp, act)
			}
		})
	}
}
