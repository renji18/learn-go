package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type EnvConfig struct {
	DATABASE_URL        string
	DRIVER              string
	PORT                string
	JWT_SECRET          string
	ACCESS_COOKIE_NAME  string
	REFRESH_COOKIE_NAME string
	EMAIL_FROM          string
	EMAIL_PASSWORD      string
	REDIS_ADDR          string
	REDIS_PORT          string
}

var Config EnvConfig

func ParseConfig() {
	// get current working directory
	cwd, err := os.Getwd()
	Fatal(fmt.Errorf("Error while accessing current working directory: %v", err), err)

	envFilePath := filepath.Join(cwd, ".env")

	// open env file
	file, err := os.Open(envFilePath)
	Fatal(fmt.Errorf("Error opening env file: %v", err), err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read env
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		splitString := strings.SplitN(line, "=", 2)

		key := strings.TrimSpace(splitString[0])
		value := strings.TrimSpace(splitString[1])

		value = strings.TrimPrefix(value, "\"")
		value = strings.TrimSuffix(value, "\"")

		// set env in env config
		switch key {
		case "DATABASE_URL":
			Config.DATABASE_URL = value

		case "DRIVER":
			Config.DRIVER = value

		case "PORT":
			Config.PORT = value

		case "JWT_SECRET":
			Config.JWT_SECRET = value

		case "ACCESS_COOKIE_NAME":
			Config.ACCESS_COOKIE_NAME = value

		case "REFRESH_COOKIE_NAME":
			Config.REFRESH_COOKIE_NAME = value

		case "EMAIL_FROM":
			Config.EMAIL_FROM = value

		case "EMAIL_PASSWORD":
			Config.EMAIL_PASSWORD = value

		case "REDIS_ADDR":
			Config.REDIS_ADDR = value

		case "REDIS_PORT":
			Config.REDIS_PORT = value
		}
	}

	err = scanner.Err()
	Fatal(fmt.Errorf("Error while scanning: %v", err), err)
}
