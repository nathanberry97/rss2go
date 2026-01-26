package css

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func HashCSSFile(outputDir, tempCSSFile string) (string, error) {
	tempOutput := filepath.Join(outputDir, tempCSSFile)

	if _, err := os.Stat(tempOutput); err != nil {
		if os.IsNotExist(err) {

			files, err := os.ReadDir(outputDir)
			if err != nil {
				return "", fmt.Errorf("Failed to read CSS output directory %s: %v", outputDir, err)
			}
			for _, f := range files {
				if strings.HasPrefix(f.Name(), "style-") && strings.HasSuffix(f.Name(), ".css") {
					return f.Name(), nil
				}
			}
			return "", fmt.Errorf("No temporary CSS file (%s) or hashed CSS file found in %s", tempCSSFile, outputDir)
		}
		return "", fmt.Errorf("Failed to stat temporary CSS file %s: %v", tempOutput, err)
	}

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
