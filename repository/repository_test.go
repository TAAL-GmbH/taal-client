package repository_test

import (
	"taal-client/repository"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestISO8601(t *testing.T) {

	dateString := time.Date(2022, 5, 1, 10, 15, 10, 321000000, time.UTC).Format(repository.ISO8601)

	is := is.New(t)
	is.Equal("2022-05-01T10:15:10.321Z", dateString)

	t.Log(dateString)
}
