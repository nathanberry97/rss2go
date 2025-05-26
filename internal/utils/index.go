package utils

import (
	"fmt"
	"time"
)

func FormatDate(dateStr, inputFormat, outputFormat string) (string, error) {
	t, err := time.Parse(inputFormat, dateStr)

	if err != nil {
		return "", fmt.Errorf("Failed to parse time: %w", err)
	}

	return t.Format(outputFormat), nil
}

func FormatTime(dateStr, inputFormat string) (string, error) {
	t, err := time.Parse(inputFormat, dateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse time: %w", err)
	}

	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return "just now", nil
	case duration < time.Hour:
		mins := int(duration.Minutes())
		if mins == 1 {
			return "1 minute ago", nil
		}
		return fmt.Sprintf("%d minutes ago", mins), nil
	case duration < 24*time.Hour:
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago", nil
		}
		return fmt.Sprintf("%d hours ago", hours), nil
	case duration < 30*24*time.Hour:
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago", nil
		}
		return fmt.Sprintf("%d days ago", days), nil
	case duration < 365*24*time.Hour:
		months := int(duration.Hours() / (24 * 30))
		if months == 1 {
			return "1 month ago", nil
		}
		return fmt.Sprintf("%d months ago", months), nil
	default:
		years := int(duration.Hours() / (24 * 365))
		if years == 1 {
			return "1 year ago", nil
		}
		return fmt.Sprintf("%d years ago", years), nil
	}
}
