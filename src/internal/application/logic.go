package application

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	UrlBase = "https://api.github.com/repos/%s/%s/issues"
)

type Issue struct {
	Id     int    `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Reporter struct {
	org     string
	repo    string
	token   string
	request Request
}

type Option func(*Reporter)

func WithRequest(request Request) Option {
	return func(r *Reporter) {
		r.request = request
	}
}

type Request interface {
	ExecHttpReq(req *http.Request, token string) (http.Header, []byte, error)
}

type HttpRequest struct {
	Client *http.Client
}

var _ = Request(&HttpRequest{})

func NewReporter(org, repo, token string, options ...Option) *Reporter {
	reporter := &Reporter{
		org:   org,
		repo:  repo,
		token: token,
		request: &HttpRequest{
			Client: &http.Client{
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
		},
	}
	for _, option := range options {
		option(reporter)
	}
	return reporter
}

func (r *Reporter) Report(title, input, search string, labels []string) (*Issue, error) {
	exist, number, err := r.isExistIssue(title, search)
	if err != nil {
		return nil, err
	}
	if exist {
		return r.updateIssue(title, input, labels, number)
	} else {
		return r.createIssue(title, input, labels)
	}
}

func (r *Reporter) createIssue(title, input string, labels []string) (*Issue, error) {
	body, err := r.readBodyText(input)
	if err != nil {
		return nil, err
	}
	data := struct {
		Title  string   `json:"title"`
		Body   string   `json:"body"`
		Labels []string `json:"labels"`
	}{
		Title:  title,
		Body:   body,
		Labels: labels,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(UrlBase, r.org, r.repo),
		bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	_, retBody, err := r.request.ExecHttpReq(req, r.token)
	if err != nil {
		return nil, err
	}
	issue := &Issue{}
	json.Unmarshal(retBody, issue)
	return issue, nil
}

func (r *Reporter) updateIssue(title, input string, labels []string, number int) (*Issue, error) {
	body, err := r.readBodyText(input)
	if err != nil {
		return nil, err
	}
	data := struct {
		Title  string   `json:"title"`
		Body   string   `json:"body"`
		Labels []string `json:"labels"`
	}{
		Title:  title,
		Body:   body,
		Labels: labels,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%d", fmt.Sprintf(UrlBase, r.org, r.repo), number),
		bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	_, retBody, err := r.request.ExecHttpReq(req, r.token)
	if err != nil {
		return nil, err
	}
	issue := &Issue{}
	json.Unmarshal(retBody, issue)
	return issue, nil
}

func (r *Reporter) readBodyText(input string) (string, error) {
	b, err := os.ReadFile(input)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *Reporter) isExistIssue(title, search string) (bool, int, error) {
	for i := 1; ; i++ {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf(UrlBase, r.org, r.repo),
			nil)
		if err != nil {
			return false, 0, err
		}
		q := req.URL.Query()
		q.Add("labels", search)
		q.Add("per_page", "100")
		q.Add("page", fmt.Sprintf("%d", i))
		req.URL.RawQuery = q.Encode()

		header, body, err := r.request.ExecHttpReq(req, r.token)
		if err != nil {
			return false, 0, err
		}
		tmpArr := make([]Issue, 0, 100)
		json.Unmarshal(body, &tmpArr)
		for _, i := range tmpArr {
			if i.Title == title {
				return true, i.Number, nil
			}
		}
		linkVal := header.Get("Link")
		links := strings.Split(linkVal, ",")
		for _, link := range links {
			if strings.Contains(link, "rel=\"next\"") {
				goto loop
			}
		}
		break
	loop:
	}
	return false, 0, nil
}

func (r *HttpRequest) ExecHttpReq(req *http.Request, token string) (http.Header, []byte, error) {
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	header := resp.Header
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, nil, fmt.Errorf("HTTP status code: %d, error: %s", resp.StatusCode, body)
	}
	return header, body, nil
}
