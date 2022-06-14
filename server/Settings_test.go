package server

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestUpdateSettings(t *testing.T) {
	t.Run("update settings", func(t *testing.T) {

		is := is.New(t)

		f, err := os.Create("settings_test.conf")
		is.NoErr(err)

		defer f.Close()
		defer os.Remove("settings_test.conf")
		defer os.Remove("settings_test.conf.old")

		_, err = f.WriteString("## this is setting 1\nsetting1=value1")
		is.NoErr(err)

		err = updateSettings("test", "false", "settings_test.conf", "settings_test.conf.tmp", "settings_test.conf.old")
		is.NoErr(err)

		err = updateSettings("setting1", "value2", "settings_test.conf", "settings_test.conf.tmp", "settings_test.conf.old")
		is.NoErr(err)

		content, err := ioutil.ReadFile("settings_test.conf")
		is.NoErr(err)

		resultContent := string(content)
		expectedContent := "## this is setting 1\nsetting1=value2\ntest=false\n"
		is.Equal(expectedContent, resultContent)

	})
}
