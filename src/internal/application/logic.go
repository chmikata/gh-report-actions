package application

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	UrlBase = "https://api.github.com/repos/%s/%s/issues"
)

type Issue struct {
	title string
	body  string
	label string
}

type Reporter struct {
	org    string
	repo   string
	token  string
	client *http.Client
	issue  *Issue
}

func NewReporter(org, repo, token string) *Reporter {
	return &Reporter{
		org:   org,
		repo:  repo,
		token: token,
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   5 * time.Second,
				ResponseHeaderTimeout: 5 * time.Second,
				IdleConnTimeout:       5 * time.Second,
				MaxIdleConns:          100,
				MaxConnsPerHost:       100,
				MaxIdleConnsPerHost:   100,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (r *Reporter) Report(title, input, label string) error {
	body, err := r.readBodyText(input)
	if err != nil {
		return err
	}
	r.setIssue(title, body, label)
	return nil
}

func (r *Reporter) setIssue(title, body, label string) {
	r.issue = &Issue{
		title: title,
		body:  body,
		label: label,
	}
}

func (r *Reporter) readBodyText(input string) (string, error) {
	b, err := os.ReadFile(input)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *Reporter) isExistIssue() (bool, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(UrlBase, r.org, r.repo),
		nil)
	if err != nil {
		return nil, err
	}
	return ioutil.WriteFile(output, []byte(r.issue.body), 0644)
}

func (r *Reporter) execHttpReq(req *http.Request) (http.Header, []byte, error) {
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	header := resp.Header
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("HTTP status code: %d, error: %s", resp.StatusCode, body)
	}
	return header, body, nil
}
