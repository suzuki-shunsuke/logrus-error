/*
Package logerr provides small API to embed logrus.Fields into error.

fmt.Errorf is very useful, but you can add only a string to error as context. You can't add structured data to error.
If you use logrus, you may want to add structured data to error.

This package provides API to add structured data to error and get structured data from error for logging.
*/
package logerr
