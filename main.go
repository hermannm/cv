package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"hermannm.dev/devlog"
	"hermannm.dev/devlog/errlog"
	"hermannm.dev/devlog/sloginit"

	"hermannm.dev/cv/cvbuilder"
)

func main() {
	sloginit.InitPrettyLogHandler(
		os.Stdout,
		&devlog.Options{RenameErrorAttrKey: "cause", Level: slog.LevelDebug},
	)

	ctx := context.Background()

	flags := parseCommandLineFlags()

	if flags.Application == "" {
		slog.InfoContext(ctx, "Building CV...")

		outputPath, err := cvbuilder.BuildCV(ctx, flags.Language)
		if err != nil {
			slog.ErrorContext(ctx, "Failed to build CV", errlog.Cause(err))
			os.Exit(1)
		}

		slog.InfoContext(ctx, "CV built successfully!", "path", outputPath)
	} else {
		slog.InfoContext(ctx, "Building job application...", "name", flags.Application)

		outputPath, err := cvbuilder.BuildJobApplication(ctx, flags.Application, flags.Language)
		if err != nil {
			slog.ErrorContext(ctx, "Failed to build job application", errlog.Cause(err))
			os.Exit(1)
		}

		slog.InfoContext(ctx, "Job application built successfully!", "path", outputPath)
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
