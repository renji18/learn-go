package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type EnvConfig struct {
	MONGO_URI       string
	DB_NAME         string
	COLLECTION_NAME string
}

func ParseConfig(filePath string) (EnvConfig, error) {
	env := EnvConfig{}

	file, err := os.Open(filePath)
	if err != nil {
		return env, fmt.Errorf("Error opening config: %v\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		splitString := strings.SplitN(line, "=", 2)

		key := strings.TrimSpace(splitString[0])
		value := strings.TrimSpace(splitString[1])

		switch key {
		case "MONGO_URI":
			env.MONGO_URI = value

		case "DB_NAME":
			env.DB_NAME = value

		case "COLLECTION_NAME":
			env.COLLECTION_NAME = value
		}
	}

	if err := scanner.Err(); err != nil {
		return env, fmt.Errorf("Error reading config:%v\n", err)
	}

	return env, nil
}
