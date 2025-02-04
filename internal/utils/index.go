package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func FormatDate(dateStr, inputFormat, outputFormat string) (string, error) {
	t, err := time.Parse(inputFormat, dateStr)

	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Failed to parse time: %w", err)
	}

	return t.Format(outputFormat), nil
}

func SetEnv(envFile string) {
	file, err := os.Open(envFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		env := strings.Split(scanner.Text(), "=")

		if strings.Contains(env[1], "'") {
			env[1] = strings.Replace(env[1], "'", "", -1)
		}

		os.Setenv(env[0], env[1])
	}
}

func Assert(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
