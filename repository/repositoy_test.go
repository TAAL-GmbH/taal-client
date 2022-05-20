package repository

import (
	"testing"
	"time"
)

func TestISO8601(t *testing.T) {
	str := time.Now().UTC().Format(ISO8601)
	t.Log(str)
}
