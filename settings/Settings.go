package settings

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Settings struct {
	ListenAddress string
	TaalUrl       string
	TaalTimeOut   string
	DBType        string
	DBFile        string
	DBHost        string
	DBPort        int
	DBName        string
	DBUsername    string
	DBPassword    string
}

var (
	settings *Settings
)

func init() {
	// Start with default values for all items...
	settings = &Settings{
		ListenAddress: ":9500",
		TaalUrl:       "https://api.taal.com",
		TaalTimeOut:   "10s",
		DBType:        "sqlite",
		DBName:        "./taal_client.db",
	}

	// Now overwrite with any settings from settings.conf
	if m := readSettingsFile(); len(m) > 0 {
		if listen, found := m["listen"]; found {
			settings.ListenAddress = listen
		}

		if url, found := m["taal_url"]; found {
			settings.TaalUrl = url
		}

		if timeout, found := m["taal_timeout"]; found {
			settings.TaalTimeOut = timeout
		}

		if database, found := m["database"]; found {
			settings.DBType = database

			switch database {
			case "sqlite":
				if file, found := m["db_filename"]; found {
					settings.DBFile = file
				}

			case "postgres":
				dbHost, found := m["db_host"]
				if !found {
					log.Fatalf("postgres database requested but missing db_host setting")
				}
				settings.DBHost = dbHost

				dbPort, found := m["db_port"]
				if !found {
					log.Fatalf("postgres database request but missing db_port setting")
				}
				p, err := strconv.Atoi(dbPort)
				if err != nil {
					log.Fatalf("db_port was present in settings.conf but was an invalid number: %v", err)
				}
				settings.DBPort = p

				dbName, found := m["db_name"]
				if !found {
					log.Fatalf("postgres database requested but missing db_name setting")
				}
				settings.DBHost = dbName

				dbUsername, found := m["db_username"]
				if !found {
					log.Fatalf("postgres database requested but missing db_username setting")
				}
				settings.DBUsername = dbUsername

				dbPassword, found := m["db_password"]
				if !found {
					log.Fatalf("postgres database requested but missing db_password setting")
				}
				settings.DBHost = dbPassword

			default:
				log.Fatalf("database was present in settings.conf but %s is not a supported database", database)
			}
		}
	}

	// Now overwrite settings with any specified environment variables...
	if listen := os.Getenv("LISTEN"); listen != "" {
		settings.ListenAddress = listen
	}

	if url := os.Getenv("TAAL_URL"); url != "" {
		settings.TaalUrl = url
	}

	if timeout := os.Getenv("TAAL_TIMEOUT"); timeout != "" {
		settings.TaalTimeOut = timeout
	}
}

func Get() *Settings {
	return settings
}

func readSettingsFile() map[string]string {
	m := make(map[string]string)

	file, err := os.Open("settings.conf")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Printf("Could not open settings.conf: %v", err)
		}
		return nil
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
				m[key] = value
			}
		}
	}

	return m
}
