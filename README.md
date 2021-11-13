# logrus-error

[![Go Reference](https://pkg.go.dev/badge/github.com/suzuki-shunsuke/logrus-error.svg)](https://pkg.go.dev/github.com/suzuki-shunsuke/logrus-error)
[![Build Status](https://github.com/suzuki-shunsuke/logrus-error/workflows/test/badge.svg)](https://github.com/suzuki-shunsuke/logrus-error/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/logrus-error.svg)](https://github.com/suzuki-shunsuke/logrus-error)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/logrus-error/main/LICENSE)

Go small library to embed [logrus](https://github.com/sirupsen/logrus).Fields into error.

## Motivation

With [fmt.Errorf](https://pkg.go.dev/fmt#Errorf), you can add additional context to error.

e.g.

```go
fmt.Errorf("get a user: %w", err)
```

[logrus](https://github.com/sirupsen/logrus) is one of most popular structured logging library.

e.g.

```go
logrus.WithFields(logrus.Fields{
	"username": username,
}).WithError(err).Error("get a user")
```

`fmt.Errorf` is very useful, but you can add only a string to error as context. You can't add structured data to error.
If you use logrus, you may want to add structured data to error.

`logrus-error` is a small library to add structured data to error and get structured data from error for logging.

Mainly logrus-error provides only two simple API.

* [WithFields](https://pkg.go.dev/github.com/suzuki-shunsuke/logrus-error/logerr#WithFields): Add structured data to error
* [WithError](https://pkg.go.dev/github.com/suzuki-shunsuke/logrus-error/logerr#WithError): Get structured data from error and return logrus.Entry

AS IS (without logrus-error)

```go
return fmt.Errorf("get a user (username: %s): %w", username, err)
```

```go
logrus.WithError(err).Error("add a member to a group")
```

TO BE (with logrus-error)

```go
return logerr.WithFields(fmt.Errorf("get a user: %w", err), logrus.Fields{
	"username": username,
})
```

```go
entry := logrus.NewEntry(logrus.New())
logerr.WithError(entry, err).Error("add a member to a group")
```

Using logrus-error, you can add structured data to error as context. You don't have to construct a string with [fmt's format](https://pkg.go.dev/fmt#hdr-Printing).

## Document

Please see https://pkg.go.dev/github.com/suzuki-shunsuke/logrus-error/logerr

## License

[MIT](LICENSE)
