package server

import (
	"testing"
)

func TestUpdateSettings(t *testing.T) {
	updateSettings("dbType", "bob")
}
