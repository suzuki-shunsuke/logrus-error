package logerr

import (
	"errors"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_getFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		err   error
		exp   logrus.Fields
	}{
		{
			title: "nil",
			err:   nil,
			exp:   nil,
		},
		{
			title: "errors.New",
			err:   errors.New("foo"),
			exp:   nil,
		},
		{
			title: "single",
			err: &logrusError{
				fields: logrus.Fields{
					"foo": "bar",
				},
			},
			exp: logrus.Fields{
				"foo": "bar",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			fields := getFields(d.err)
			if !reflect.DeepEqual(d.exp, fields) {
				t.Fatalf("wanted %v, got %v", d.exp, fields)
			}
		})
	}
}
