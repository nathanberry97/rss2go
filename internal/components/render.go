package components

import (
	"fmt"
	"html/template"
	"strings"
)

func renderTemplates(fragmentPath, tmplName string, data any) (template.HTML, error) {
	tmpl, err := template.ParseFiles(fragmentPath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", fragmentPath, err)
	}

	var sb strings.Builder
	err = tmpl.ExecuteTemplate(&sb, tmplName, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", tmplName, err)
	}

	return template.HTML(sb.String()), nil
}
