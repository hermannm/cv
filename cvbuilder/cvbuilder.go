package cvbuilder

import (
	"fmt"

	"hermannm.dev/wrap"
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
		return "", wrap.Error(err, "failed to parse CV")
	}

	personalInfo, err := parsePersonalInfoFile(language)
	if err != nil {
		return "", wrap.Error(err, "failed to parse personal info")
	}

	template := CVTemplate{CV: cv, PersonalInfo: personalInfo}
	outputPath, err = renderTemplate("cv", false, template)
	if err != nil {
		return "", wrap.Error(err, "failed to render CV template")
	}

	if err := generateTailwindCSS(); err != nil {
		return "", wrap.Error(err, "failed to generate styles for rendered CV")
	}

	return outputPath, nil
}

func BuildJobApplication(applicationName string, language string) (outputPath string, err error) {
	filePath := fmt.Sprintf("%s/%s.md", JobApplicationsContentDir, applicationName)
	content, err := parseMarkdownFile(filePath)
	if err != nil {
		return "", wrap.Error(err, "failed to read job application")
	}

	personalInfo, err := parsePersonalInfoFile(language)
	if err != nil {
		return "", wrap.Error(err, "failed to parse personal info")
	}

	template := JobApplicationTemplate{Application: content, PersonalInfo: personalInfo}
	outputPath, err = renderTemplate(applicationName, true, template)
	if err != nil {
		return "", wrap.Error(err, "failed to render job application template")
	}

	if err := generateTailwindCSS(); err != nil {
		return "", wrap.Error(err, "failed to generate styles for rendered job application")
	}

	return outputPath, nil
}
