package cvbuilder

import (
	"bufio"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

func renderTemplate(
	outputName string, isJobApplication bool, templateData any,
) (outputPath string, err error) {
	outputPath, directories := getRenderOutputPath(outputName, isJobApplication)

	permissions := fs.FileMode(0755)
	if err := os.MkdirAll(directories, permissions); err != nil {
		return "", fmt.Errorf(
			"failed to create render output directories '%s': %w", directories, err,
		)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", fmt.Errorf("failed to create template output file '%s': %w", outputPath, err)
	}
	defer outputFile.Close()

	templates, err := parseTemplates()
	if err != nil {
		return "", fmt.Errorf("failed to parse template files: %w", err)
	}

	var templateName string
	if isJobApplication {
		templateName = JobApplicationTemplateName
	} else {
		templateName = CVTemplateName
	}

	if err := templates.ExecuteTemplate(outputFile, templateName, templateData); err != nil {
		return "", fmt.Errorf("failed to execute template '%s': %w", templateName, err)
	}

	return outputPath, nil
}

func getRenderOutputPath(
	outputName string, isJobApplication bool,
) (outputPath string, directories string) {
	var dirs strings.Builder
	if isJobApplication {
		dirs.WriteString(JobApplicationsOutputDir)
	} else {
		dirs.WriteString(OutputDir)
	}

	var fileName string
	outputNameParts := strings.Split(outputName, "/")
	for i, part := range outputNameParts {
		if i == len(outputNameParts)-1 {
			fileName = part
		} else {
			dirs.WriteRune('/')
			dirs.WriteString(part)
		}
	}

	directories = dirs.String()
	outputPath = fmt.Sprintf("%s/%s.html", directories, fileName)
	return outputPath, directories
}

func parseTemplates() (*template.Template, error) {
	templates := template.New("templates")

	templatesPattern := fmt.Sprintf("%s/*.tmpl", TemplatesDir)
	templates, err := templates.ParseGlob(templatesPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	return templates, nil
}

func generateTailwindCSS() error {
	outputPath := fmt.Sprintf("%s/%s", OutputDir, CSSFileName)
	return execCommand(
		"tailwind", "npx", "tailwindcss", "-i", CSSFileName, "-o", outputPath,
	)
}

func execCommand(displayName string, commandName string, args ...string) error {
	command := exec.Command(commandName, args...)

	stderr, err := command.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get pipe to %s's error output: %w", displayName, err)
	}

	if err := command.Start(); err != nil {
		return fmt.Errorf("failed to start %s command: %w", displayName, err)
	}

	errScanner := bufio.NewScanner(stderr)
	var commandErrs strings.Builder
	for errScanner.Scan() {
		if commandErrs.Len() == 0 {
			fmt.Fprintf(&commandErrs, "errors from %s:", displayName)
		}
		fmt.Fprintf(&commandErrs, "\n%s", errScanner.Text())
	}

	if err := command.Wait(); err != nil {
		if commandErrs.Len() == 0 {
			return fmt.Errorf("%s failed: %w", displayName, err)
		} else {
			return fmt.Errorf("%s failed: %w\n%s", displayName, err, commandErrs.String())
		}
	}

	return nil
}
