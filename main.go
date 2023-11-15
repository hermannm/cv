package main

import (
	"flag"
	"log/slog"
	"os"

	"hermannm.dev/cv/cvbuilder"
	"hermannm.dev/devlog"
	"hermannm.dev/devlog/log"
)

func main() {
	logger := slog.New(devlog.NewHandler(os.Stdout, &devlog.Options{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	flags := parseCommandLineFlags()

	if flags.Application == "" {
		log.Info("building CV...")

		outputPath, err := cvbuilder.BuildCV(flags.Language)
		if err != nil {
			log.ErrorCause(err, "failed to build CV")
			os.Exit(1)
		}

		log.Info("CV built successfully!", slog.String("path", outputPath))
	} else {
		log.Info("building job application...", slog.String("name", flags.Application))

		outputPath, err := cvbuilder.BuildJobApplication(flags.Application, flags.Language)
		if err != nil {
			log.ErrorCause(err, "failed to build job application")
			os.Exit(1)
		}

		log.Info("job application built successfully!", slog.String("path", outputPath))
	}
}

type CommandLineFlags struct {
	Application string
	Language    string
}

func parseCommandLineFlags() CommandLineFlags {
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

	return flags
}
