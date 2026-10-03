package check

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fabiodrneles/sdd-kit-demo/internal/markdown"
)

// ReasonNetwork prefixes problems where the request itself failed.
const ReasonNetwork = "erro de rede"

// ExternalOptions configures External (spec 002 FR-1).
type ExternalOptions struct {
	Timeout     time.Duration // per request
	Concurrency int           // simultaneous requests
	Client      *http.Client  // nil: a client that does not follow redirects
}

// External checks the http(s) links of each file: 2xx and 3xx pass, 4xx, 5xx
// and network errors are problems (spec 002 FR-1, FR-2). Each URL is requested
// once, with HEAD, falling back to GET when the server refuses HEAD.
func External(files []string, opt ExternalOptions) ([]Problem, error) {
	if opt.Concurrency < 1 {
		opt.Concurrency = 8
	}
	if opt.Timeout <= 0 {
		opt.Timeout = 10 * time.Second
	}
	client := opt.Client
	if client == nil {
		client = &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	}

	type use struct {
		file string
		line int
	}
	uses := map[string][]use{}
	var urls []string
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		for _, link := range markdown.Links(string(data)) {
			lower := strings.ToLower(link.Target)
			if !strings.HasPrefix(lower, "http://") && !strings.HasPrefix(lower, "https://") {
				continue
			}
			if _, seen := uses[link.Target]; !seen {
				urls = append(urls, link.Target)
			}
			uses[link.Target] = append(uses[link.Target], use{file, link.Line})
		}
	}

	reasons := make([]string, len(urls))
	sem := make(chan struct{}, opt.Concurrency)
	var wg sync.WaitGroup
	for i, u := range urls {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			reasons[i] = probe(client, u, opt.Timeout)
		}()
	}
	wg.Wait()

	var problems []Problem
	for i, u := range urls {
		if reasons[i] == "" {
			continue
		}
		for _, at := range uses[u] {
			problems = append(problems, Problem{at.file, at.line, reasons[i], u})
		}
	}
	Sort(problems)
	return problems, nil
}

// probe returns "" when the URL answers 2xx or 3xx, else the reason.
func probe(client *http.Client, url string, timeout time.Duration) string {
	status, err := request(client, http.MethodHead, url, timeout)
	if err == nil && (status == http.StatusMethodNotAllowed || status == http.StatusNotImplemented) {
		status, err = request(client, http.MethodGet, url, timeout)
	}
	switch {
	case err != nil:
		return fmt.Sprintf("%s: %v", ReasonNetwork, err)
	case status >= 200 && status < 400:
		return ""
	default:
		return fmt.Sprintf("HTTP %d", status)
	}
}

func request(client *http.Client, method, url string, timeout time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", "linkcheck (+https://github.com/fabiodrneles/sdd-kit-demo)")
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	_ = resp.Body.Close()
	return resp.StatusCode, nil
}
