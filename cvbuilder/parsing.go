package cvbuilder

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

var validate *validator.Validate = validator.New()

func parseCVFile(language string) (CV, error) {
	var filePath string
	if language == "" {
		filePath = fmt.Sprintf("%s/%s.yml", ContentDir, CVYAMLFileName)
	} else {
		filePath = fmt.Sprintf("%s/%s_%s.yml", ContentDir, CVYAMLFileName, language)
	}

	cv, err := parseYAMLFile[CV](filePath)
	if err != nil {
		return CV{}, fmt.Errorf("failed to parse CV YAML file: %w", err)
	}

	if err := validate.Struct(cv); err != nil {
		return CV{}, fmt.Errorf("invalid CV: %w", err)
	}

	for i, experience := range cv.WorkExperience {
		experience.Organization, err = parseMarkdownField([]byte(experience.Organization), true)
		if err != nil {
			return CV{}, fmt.Errorf(
				"invalid organization in work experience '%s': %w", experience.Title, err,
			)
		}

		experience.Description, err = parseMarkdownField([]byte(experience.Description), false)
		if err != nil {
			return CV{}, fmt.Errorf(
				"invalid description in work experience '%s': %w", experience.Title, err,
			)
		}

		cv.WorkExperience[i] = experience
	}

	return cv, nil
}

func parsePersonalInfoFile(language string) (PersonalInfo, error) {
	var filePath string
	if language == "" {
		filePath = fmt.Sprintf("%s/%s.yml", ContentDir, PersonalInfoYAMLFileName)
	} else {
		filePath = fmt.Sprintf("%s/%s_%s.yml", ContentDir, PersonalInfoYAMLFileName, language)
	}

	info, err := parseYAMLFile[PersonalInfo](filePath)
	if err != nil {
		return PersonalInfo{}, fmt.Errorf("failed to parse personal info YAML file: %w", err)
	}

	if err := validate.Struct(info); err != nil {
		return PersonalInfo{}, fmt.Errorf("invalid personal info: %w", err)
	}

	if err := info.setAge(); err != nil {
		return PersonalInfo{}, fmt.Errorf("failed to set age field on personal info: %w", err)
	}

	return info, nil
}

func parseYAMLFile[Format any](yamlFilePath string) (Format, error) {
	var dest Format

	yamlContent, err := os.ReadFile(yamlFilePath)
	if err != nil {
		return dest, fmt.Errorf("failed to read file '%s': %w", yamlFilePath, err)
	}

	if err := yaml.Unmarshal(yamlContent, &dest); err != nil {
		return dest, fmt.Errorf("failed to parse YAML file '%s': %w", yamlFilePath, err)
	}

	return dest, nil
}

func parseMarkdownFile(markdownFilePath string) (template.HTML, error) {
	rawContent, err := os.ReadFile(markdownFilePath)
	if err != nil {
		return template.HTML(""), fmt.Errorf("failed to read file '%s': %w", markdownFilePath, err)
	}

	markdownParser := goldmark.New(goldmark.WithRendererOptions(html.WithUnsafe()))

	var parsedContent strings.Builder
	if err := markdownParser.Convert(rawContent, &parsedContent); err != nil {
		return template.HTML(""), fmt.Errorf(
			"failed to parse markdown file '%s': %w", markdownFilePath, err,
		)
	}

	return template.HTML(parsedContent.String()), nil
}

func parseMarkdownField(fieldValue []byte, removeParagraphTags bool) (template.HTML, error) {
	var parsedField strings.Builder
	if err := goldmark.Convert([]byte(fieldValue), &parsedField); err != nil {
		return template.HTML(""), fmt.Errorf("failed to parse markdown field: %w", err)
	}

	fieldString := parsedField.String()
	if removeParagraphTags {
		fieldString = removeParagraphTagsAroundHTML(fieldString)
	}

	return template.HTML(fieldString), nil
}

func removeParagraphTagsAroundHTML(html string) string {
	html = strings.TrimSpace(html)
	html, _ = strings.CutPrefix(html, "<p>")
	html, _ = strings.CutSuffix(html, "</p>")
	return html
}
