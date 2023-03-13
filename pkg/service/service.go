package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	zlog "github.com/Bhargav-InfraCloud/zerolog-wrapper"
)

type Repository struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"html_url"`
}

type Operator struct {
	logger zlog.Logger
	client http.Client
}

func NewOperator(ctx context.Context, timeout time.Duration) *Operator {
	return &Operator{
		logger: zlog.FromContext(ctx),
		client: http.Client{
			Timeout: timeout,
		},
	}
}

func (o *Operator) ListRepositories(organization string) ([]Repository, error) {
	url := fmt.Sprintf(`https://api.github.com/orgs/%s/repos`, organization)

	respBody, err := o.fetchResponse(url)
	if err != nil {
		return nil, err
	}

	repos, err := o.parseResponse(respBody)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func (o *Operator) fetchResponse(url string) ([]byte, error) {
	resp, err := o.client.Get(url)
	if err != nil {
		o.logger.Error().Err(err).Msg("Failed to get the repos list")

		return nil, fmt.Errorf("failed to get the repos list: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		o.logger.Error().Int("status-code", resp.StatusCode).
			Msg("Invalid response status from repo listing")

		return nil, fmt.Errorf("invalid response status from repo listing")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		o.logger.Error().Err(err).Msg("Failed to read response body")

		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}

func (o *Operator) parseResponse(data []byte) ([]Repository, error) {
	var repos []Repository

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&repos)
	if err != nil {
		o.logger.Error().Err(err).Msg("Failed to JSON decode")

		return nil, fmt.Errorf("failed to JSON decode: %w", err)
	}

	return repos, nil
}

// curl -L \
//   -H "Accept: application/vnd.github+json" \
//   -H "Authorization: Bearer <YOUR-TOKEN>"\
//   -H "X-GitHub-Api-Version: 2022-11-28" \
//   https://api.github.com/orgs/ORG/repos
