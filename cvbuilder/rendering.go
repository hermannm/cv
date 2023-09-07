package cvbuilder

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"hermannm.dev/wrap"
)

func renderTemplate(
	outputName string, isJobApplication bool, templateData any,
) (outputPath string, err error) {
	outputPath, directories := getRenderOutputPath(outputName, isJobApplication)

	permissions := fs.FileMode(0755)
	if err := os.MkdirAll(directories, permissions); err != nil {
		return "", wrap.Errorf(
			err, "failed to create render output directories '%s'", directories,
		)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", wrap.Errorf(err, "failed to create template output file '%s'", outputPath)
	}
	defer outputFile.Close()

	templates, err := parseTemplates()
	if err != nil {
		return "", wrap.Error(err, "failed to parse template files")
	}

	var templateName string
	if isJobApplication {
		templateName = JobApplicationTemplateName
	} else {
		templateName = CVTemplateName
	}

	if err := templates.ExecuteTemplate(outputFile, templateName, templateData); err != nil {
		return "", wrap.Errorf(err, "failed to execute template '%s'", templateName)
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
		return nil, wrap.Error(err, "failed to parse templates")
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
		return wrap.Errorf(err, "failed to get pipe to %s command's error output", displayName)
	}

	if err := command.Start(); err != nil {
		return wrap.Errorf(err, "failed to start %s command", displayName)
	}

	errScanner := bufio.NewScanner(stderr)
	var commandErrs strings.Builder
	for errScanner.Scan() {
		if commandErrs.Len() != 0 {
			commandErrs.WriteRune('\n')
		}
		commandErrs.WriteString(errScanner.Text())
	}

	if err := command.Wait(); err != nil {
		err = fmt.Errorf("%s command failed: %w", displayName, err)
		if commandErrs.Len() == 0 {
			return err
		} else {
			return wrap.Error(errors.New(commandErrs.String()), err.Error())
		}
	}

	return nil
}
