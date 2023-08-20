package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"hermannm.dev/cv/cvbuilder"
)

func main() {
	flags, err := parseCommandLineFlags()
	if err != nil {
		fmt.Printf("invalid args: %v\n", err)
		os.Exit(1)
	}

	if flags.CV {
		outputPath, err := cvbuilder.BuildCV(flags.Language)
		if err != nil {
			fmt.Printf("failed to build CV: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("CV built successfully! Output in %s\n", outputPath)
	} else {
		outputPath, err := cvbuilder.BuildJobApplication(flags.Application, flags.Language)
		if err != nil {
			fmt.Printf("failed to build job application: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Job application built successfully! Output in %s\n", outputPath)
	}
}

type CommandLineFlags struct {
	CV          bool
	Application string
	Language    string
}

func parseCommandLineFlags() (CommandLineFlags, error) {
	var flags CommandLineFlags

	flag.BoolVar(&flags.CV, "cv", false, "Generate CV from content/cv.yml file.")
	flag.StringVar(
		&flags.Application,
		"application",
		"",
		"Generate job application from content/applications/[arg value].md file.",
	)
	flag.StringVar(
		&flags.Language,
		"lang",
		"",
		"Use content files with the given language code as a suffix. E.g. if lang=no is passed, then the file personal_info_no.yml will be used instead of personal_info.yml. Does not apply to job applications.",
	)
	flag.Parse()

	if !flags.CV && flags.Application == "" {
		return flags, errors.New("must pass either -cv or -application=[application name]")
	}
	if flags.CV && flags.Application != "" {
		return flags, errors.New("cannot pass both -cv and -application args")
	}

	return flags, nil
}
