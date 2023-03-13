package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Bhargav-InfraCloud/github-watch/internal/constants"
	"github.com/Bhargav-InfraCloud/github-watch/internal/flags"
	"github.com/Bhargav-InfraCloud/github-watch/pkg/service"
	zlog "github.com/Bhargav-InfraCloud/zerolog-wrapper"
)

func main() {
	var (
		token   string
		org     string
		timeout time.Duration
	)
	ctx, logger := zlog.NewLogger(context.Background(), os.Stdout, zlog.LevelDebug)

	err := flags.NewManager(constants.Project).
		AddFlagSet(
			flags.NewFlagSet(`secrets`),
			flags.NewFlag(&token, "token", "t", "Bearer token for GitHub authentication", ""),
		).
		AddFlagSet(
			flags.NewFlagSet(`environments`),
			flags.NewFlag(&org, "org", "o", "Organization name to fetch the repos list", ""),
			flags.NewFlag(&timeout, "timeout", "", "Organization name to fetch the repos list", time.Minute),
		).
		Parse(os.Args[1:])
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to parse input arguments")
	}

	logger.Debug().Str("organization", org).Str("token", token).Msg("Logging flag inputs")

	logger = zlog.FromRawLogger(logger.With().Str("organization", org).Logger())

	svcOps := service.NewOperator(ctx, timeout)
	repos, err := svcOps.ListRepositories(org)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to fetch the repo list")
	}

	fmt.Println(repos)
}
