package components

import (
	"bytes"
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

func RenderRSSTemplate(path, tmplName string, data any) ([]byte, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML template %s: %w", path, err)
	}

	var buf bytes.Buffer
	buf.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	if err := tmpl.ExecuteTemplate(&buf, tmplName, data); err != nil {
		return nil, fmt.Errorf("failed to execute XML template %s: %w", tmplName, err)
	}

	return buf.Bytes(), nil
}
