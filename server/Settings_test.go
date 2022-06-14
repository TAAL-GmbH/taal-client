package server

import (
	"testing"

	"github.com/matryer/is"
)

func TestUpdateSettings(t *testing.T) {

	t.Run("update settings", func(t *testing.T) {

		is := is.New(t)
		err := updateSettings("test", "false")

		is.NoErr(err)

	})
}
