package css

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// TODO - this is to just hash css file as then I can make the docker image much smaller
func HashCSSFile(outputDir, tempCSSFile string) (string, error) {
	tempOutput := filepath.Join(outputDir, tempCSSFile)
	file, err := os.Open(tempOutput)
	if err != nil {
		return "", fmt.Errorf("Failed to open temporary output file %s: %v", tempOutput, err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("Failed to copy contents of %s to hasher: %v", tempOutput, err)
	}
	hash := fmt.Sprintf("%x", hasher.Sum(nil))[:8]

	hashedFilename := fmt.Sprintf("style-%s.css", hash)
	finalOutput := filepath.Join(outputDir, hashedFilename)

	if err := os.Rename(tempOutput, finalOutput); err != nil {
		os.Remove(tempOutput)
		return "", fmt.Errorf("Failed to rename temporary file %s to final output %s: %v", tempOutput, finalOutput, err)
	}

	return hashedFilename, nil
}
