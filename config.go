package main

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnv() (envMap map[string]string) {
	envMap = make(map[string]string)

	file, err := os.Open(".env")
	Check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// skip comments and empty lines
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		envMap[parts[0]] = parts[1]
	}
	return
}
