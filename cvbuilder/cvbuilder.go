package cvbuilder

import (
	"fmt"
)

const (
	TemplatesDir               = "templates"
	CVTemplateName             = "cv.html.tmpl"
	JobApplicationTemplateName = "job_application.html.tmpl"

	ContentDir                = "content"
	JobApplicationsContentDir = "content/applications"
	CVYAMLFileName            = "cv"
	PersonalInfoYAMLFileName  = "personal_info"

	OutputDir                = "output"
	JobApplicationsOutputDir = "output/applications"

	CSSFileName = "styles.css"
)

func BuildCV(language string) (outputPath string, err error) {
	cv, err := parseCVFile(language)
	if err != nil {
		return "", fmt.Errorf("failed to parse CV: %w", err)
	}

	personalInfo, err := parsePersonalInfoFile(language)
	if err != nil {
		return "", fmt.Errorf("failed to parse personal info: %w", err)
	}

	template := CVTemplate{CV: cv, PersonalInfo: personalInfo}
	outputPath, err = renderTemplate("cv", false, template)
	if err != nil {
		return "", fmt.Errorf("failed to render CV template: %w", err)
	}

	if err := generateTailwindCSS(); err != nil {
		return "", fmt.Errorf("failed to generate styles for rendered CV: %w", err)
	}

	return outputPath, nil
}

func BuildJobApplication(applicationName string, language string) (outputPath string, err error) {
	filePath := fmt.Sprintf("%s/%s.md", JobApplicationsContentDir, applicationName)
	content, err := parseMarkdownFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read job application: %w", err)
	}

	personalInfo, err := parsePersonalInfoFile(language)
	if err != nil {
		return "", fmt.Errorf("failed to parse personal info: %w", err)
	}

	template := JobApplicationTemplate{Application: content, PersonalInfo: personalInfo}
	outputPath, err = renderTemplate(applicationName, true, template)
	if err != nil {
		return "", fmt.Errorf("failed to render job application template: %w", err)
	}

	if err := generateTailwindCSS(); err != nil {
		return "", fmt.Errorf("failed to generate styles for rendered job application: %w", err)
	}

	return outputPath, nil
}
