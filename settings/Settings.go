package settings

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	mu       sync.RWMutex
	settings map[string]string
)

func init() {
	settings = make(map[string]string)

	// Start with default values for all items...
	settings["listenAddress"] = ":9500"
	settings["taalUrl"] = "https://api.taal.com"
	settings["taalTimeout"] = "10s"
	settings["dbType"] = "sqlite"
	settings["dbFilename"] = "./taal_client.db"
	settings["dbHost"] = "localhost"
	settings["dbPort"] = "5432"
	settings["dbName"] = "dbname"
	settings["dbUsername"] = "username"
	settings["dbPassword"] = "password"

	// Now overwrite with any settings from settings.conf
	readSettingsFile()

	// Now overwrite settings with any specified environment variables...
	if listen := os.Getenv("LISTEN"); listen != "" {
		settings["listenAddress"] = listen
	}

	if url := os.Getenv("TAAL_URL"); url != "" {
		settings["taalUrl"] = url
	}

	if timeout := os.Getenv("TAAL_TIMEOUT"); timeout != "" {
		settings["taalTimeout"] = timeout
	}

	if debugTransactions := os.Getenv("DEBUG"); debugTransactions != "" {
		settings["debugTransactions"] = debugTransactions
	}
}

func GetJSON() ([]byte, error) {
	mu.RLock()
	defer mu.RUnlock()

	return json.Marshal(settings)
}

func Get(key string) string {
	mu.RLock()
	defer mu.RUnlock()

	return settings[key]
}

func GetInt(key string) int {
	val := Get(key)
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}

	return 0
}

func GetBool(key string) bool {
	val := Get(key)
	if b, err := strconv.ParseBool(val); err == nil {
		return b
	}

	return false
}

func GetDuration(key string) (time.Duration, error) {
	val := Get(key)
	return time.ParseDuration(val)
}

func Set(key string, val string) {
	mu.Lock()
	defer mu.Unlock()

	settings[key] = val
}

func readSettingsFile() {
	file, err := os.Open("settings.conf")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Printf("Could not open settings.conf: %v", err)
		}
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Remove comments
		if comment := strings.Index(line, "#"); comment >= 0 {
			line = line[0:comment]
		}

		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				settings[key] = value
			}
		}
	}

	return
}
