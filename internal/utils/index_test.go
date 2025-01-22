package utils

import (
	"os"
	"testing"
)

func TestSetEnv(t *testing.T) {
	// Arrange
	testFile := ".testEnv"
	os.Create(testFile)
	file, _ := os.OpenFile(testFile, os.O_APPEND|os.O_WRONLY, 0644)
	file.WriteString("TEST_ENV='test'")
	file.Close()

	// Act
	SetEnv(testFile)

	// Assert
	Assert(t, "test", os.Getenv("TEST_ENV"))
	os.Remove(testFile)
}

func TestAssert(t *testing.T) {
	// Arrange
	expected := "Test"
	actual := "Test"

	// Act & Assert
	Assert(t, expected, actual)
}
