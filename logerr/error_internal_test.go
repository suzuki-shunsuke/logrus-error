package logerr

import (
	"errors"
	"testing"
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
