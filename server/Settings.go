package server

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"taal-client/settings"
)

var (
	lineRegex = regexp.MustCompile(`^(\s*)(\S+)(\s*=\s*)(\S+){0,1}(\s*)(#.*)*$`)
)

const (
	settingsFile     = "settings.conf"
	settingsTempFile = "settings.conf.tmp"
	settingsOldFile  = "settings.conf.old"
)

func (s Server) getSettings(c echo.Context) error {
	b, err := settings.GetJSON()
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errPutSettingsGetJson, err)
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, b)
}

type setting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s Server) putSetting(c echo.Context) error {
	set := new(setting)
	if err := c.Bind(set); err != nil {
		return s.sendError(c, http.StatusInternalServerError, errPutSettingsBind, err)
	}

	switch set.Key {
	case "taalTimeout":
		timeout, err := time.ParseDuration(set.Value)
		if err != nil {
			return s.sendError(c, http.StatusInternalServerError, errPutSettingsGetDuration, errors.Wrapf(err, "taal_timeout of %s is invalid", set.Value))
		}

		s.taal.SetTimeout(timeout)
	case "taalUrl":
		taalURL := set.Value
		s.taal.SetUrl(taalURL)
	}

	if err := updateSettings(set.Key, set.Value, settingsFile, settingsTempFile, settingsOldFile); err != nil {
		return s.sendError(c, http.StatusInternalServerError, errPutSettingsUpdateSettings, err)
	}

	settings.Set(set.Key, set.Value)

	return c.String(http.StatusOK, "OK")
}

func updateSettings(key, value string, settingsFile string, settingsTempFile string, settingsOldFile string) error {
	w, err := os.Create(settingsTempFile)
	if err != nil {
		return errors.Wrapf(err, "could not create settings.conf.tmp")
	}

	r, err := os.Open(settingsFile)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return errors.Wrapf(err, "could not open settings.conf")
		}

		if _, err := w.WriteString(fmt.Sprintf("%s=%s\n", key, value)); err != nil {
			return errors.Wrapf(err, "could not write setting")
		}

		if err := w.Close(); err != nil {
			return errors.Wrapf(err, "could not close temp file")
		}

		if err := os.Rename(settingsTempFile, settingsFile); err != nil {
			return errors.Wrapf(err, "could not rename temp file")
		}

		return nil
	}

	written := false
	_ = written

	value = strings.TrimSpace(value)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if written {
			if _, err := w.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
				return errors.Wrapf(err, "could not write to temp file")
			}

			continue
		}

		parts := lineRegex.FindAllStringSubmatch(line, -1)

		if len(parts) != 1 {
			if _, err := w.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
				return errors.Wrapf(err, "could not write to temp file")
			}

			continue
		}

		k := parts[0][2]

		if k == key {
			var sb strings.Builder

			sb.WriteString(parts[0][1]) // Leading space
			sb.WriteString(k)
			sb.WriteString(parts[0][3]) // Equals
			sb.WriteString(value)
			sb.WriteString(parts[0][5]) // Spaces to comment
			sb.WriteString(parts[0][6]) // Comment
			sb.WriteByte('\n')

			if _, err := w.WriteString(sb.String()); err != nil {
				return errors.Wrapf(err, "could not write to temp file")
			}

			written = true
		} else {
			if _, err := w.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
				return errors.Wrapf(err, "could not write to temp file")
			}
		}
	}

	if !written {
		if _, err := w.WriteString(fmt.Sprintf("%s=%s\n", key, value)); err != nil {
			return errors.Wrapf(err, "could not write to temp file")
		}
	}

	if err := w.Close(); err != nil {
		return errors.Wrapf(err, "could not close temp file")
	}

	if err := os.Rename(settingsFile, settingsOldFile); err != nil {
		return errors.Wrapf(err, "could not rename original file")
	}

	if err := os.Rename(settingsTempFile, settingsFile); err != nil {
		return errors.Wrapf(err, "could not rename temp file")
	}

	return nil
}
