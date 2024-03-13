package logerr_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/logrus-error/logerr"
)

func TestAppendFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		title   string
		fields1 logrus.Fields
		fields2 logrus.Fields
		exp     logrus.Fields
	}{
		{
			title:   "nil",
			fields1: nil,
			fields2: nil,
			exp:     logrus.Fields{},
		},
		{
			title: "fields2 overwrites fields1",
			fields1: logrus.Fields{
				"foo":  "foo_value",
				"name": "foo",
			},
			fields2: logrus.Fields{
				"name": "bar",
				"bar":  "bar_value",
			},
			exp: logrus.Fields{
				"foo":  "foo_value",
				"name": "bar",
				"bar":  "bar_value",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			fields := logerr.AppendFields(d.fields1, d.fields2)
			if !reflect.DeepEqual(d.exp, fields) {
				t.Fatalf("wanted %v, got %v", d.exp, fields)
			}
		})
	}
}

func TestWithError(t *testing.T) {
	t.Parallel()
	data := []struct {
		title     string
		entry     *logrus.Entry
		err       error
		expErr    string
		expFields logrus.Fields
	}{
		{
			title:     "nil",
			expFields: logrus.Fields{},
		},
		{
			title: "entry is nil",
			err:   errors.New("foo"),
			expFields: logrus.Fields{
				"error": errors.New("foo"),
			},
			expErr: "foo",
		},
		{
			title:     "err is nil",
			entry:     logrus.NewEntry(logrus.New()),
			expFields: logrus.Fields{},
		},
		{
			title: "simple",
			err:   errors.New("foo"),
			entry: logrus.WithFields(logrus.Fields{
				"bar": "bar",
			}),
			expErr: "foo",
			expFields: logrus.Fields{
				"error": errors.New("foo"),
				"bar":   "bar",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			entry := logerr.WithError(d.entry, d.err)
			if !reflect.DeepEqual(d.expFields, entry.Data) {
				t.Fatalf("wanted %+v, got %+v", d.expFields, entry.Data)
			}
			err, ok := entry.Data["error"]
			if ok {
				s := err.(error).Error() //nolint:forcetypeassert
				if !reflect.DeepEqual(d.expErr, s) {
					t.Fatalf("wanted %s, got %s", d.expErr, s)
				}
				return
			}
			if d.expErr != "" {
				t.Fatalf(`wanted "", got %+v`, d.expErr)
			}
		})
	}
}

func TestWithFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		title  string
		fields logrus.Fields
		err    error
		exp    error
	}{
		{
			title: "nil",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			err := logerr.WithFields(d.err, d.fields)
			if !reflect.DeepEqual(d.exp, err) {
				t.Fatalf("wanted %+v, got %+v", d.exp, err)
			}
		})
	}
}

func TestWithText(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		text  string
		err   error
		exp   error
	}{
		{
			title: "nil",
		},
		{
			title: "hello",
			err:   errors.New("foo"),
			text:  "hello",
			exp:   fmt.Errorf("hello: %w", errors.New("foo")),
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			err := logerr.WithText(d.err, d.text)
			if !reflect.DeepEqual(d.exp, err) {
				t.Fatalf("wanted %+v, got %+v", d.exp, err)
			}
		})
	}
}
