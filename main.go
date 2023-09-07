package main

import (
	"flag"
	"fmt"
	"os"

	"hermannm.dev/cv/cvbuilder"
	"hermannm.dev/wrap"
)

func main() {
	flags, err := parseCommandLineFlags()
	if err != nil {
		fmt.Printf("invalid args: %v\n", err)
		os.Exit(1)
	}

	if flags.Application == "" {
		outputPath, err := cvbuilder.BuildCV(flags.Language)
		if err != nil {
			fmt.Println(wrap.Error(err, "failed to build CV"))
			os.Exit(1)
		}
		fmt.Printf("CV built successfully! Output in %s\n", outputPath)
	} else {
		outputPath, err := cvbuilder.BuildJobApplication(flags.Application, flags.Language)
		if err != nil {
			fmt.Println(wrap.Error(err, "failed to build job application"))
			os.Exit(1)
		}
		fmt.Printf("Job application built successfully! Output in %s\n", outputPath)
	}
}

type CommandLineFlags struct {
	Application string
	Language    string
}

func parseCommandLineFlags() (CommandLineFlags, error) {
	var flags CommandLineFlags

	flag.StringVar(
		&flags.Application,
		"application",
		"",
		"Set to generate job application instead of a CV. Generates application from content/applications/[arg value].md file.",
	)
	flag.StringVar(
		&flags.Language,
		"lang",
		"",
		"Use content files with the given language code as a suffix. E.g. if lang=no is passed, then the file personal_info_no.yml will be used instead of personal_info.yml. Does not apply to job applications.",
	)
	flag.Parse()

	return flags, nil
}
