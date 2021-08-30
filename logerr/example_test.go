package logerr_test

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/logrus-error/logerr"
)

func Example() {
	logrus.SetOutput(os.Stdout)
	logE := logrus.WithFields(logrus.Fields{
		"program": "example",
	}).WithTime(time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC))
	if err := core(); err != nil {
		// Log error
		logerr.WithError(logE, err).Error("program exits")
		// Output:
		// time="2020-01-01T12:00:00Z" level=error msg="program exits" error="foo4: foo2: foo" foo2=foo2 foo4=foo4 name=foo4 program=example
	}
}

func core() error {
	return fmt.Errorf("foo4: %w", foo4())
}

func foo1() error {
	return errors.New("foo")
}

func foo2() error {
	// Add fields to error
	return logerr.WithFields(foo1(), logrus.Fields{ //nolint:wrapcheck
		"name": "foo2",
		"foo2": "foo2",
	})
}

func foo3() error {
	return fmt.Errorf("foo2: %w", foo2())
}

func foo4() error {
	// Add fields to error
	return logerr.WithFields(foo3(), logrus.Fields{ //nolint:wrapcheck
		"name": "foo4",
		"foo4": "foo4",
	})
}
