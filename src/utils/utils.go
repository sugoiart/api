package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/syumai/workers/cloudflare"
	"github.com/syumai/workers/cloudflare/cache"
	"github.com/syumai/workers/cloudflare/fetch"
	"io"
	"log"
	"net/http"
	"path"
	"time"
)

func RequestImages(url string, target interface{}, r *http.Request) error {
	cacheKeyReq, err := http.NewRequestWithContext(r.Context(), http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create cache key request: %w", err)
	}

	c := cache.New()

	cachedResp, err := c.Match(cacheKeyReq, nil)
	if err == nil && cachedResp != nil {
		defer cachedResp.Body.Close()
		return json.NewDecoder(cachedResp.Body).Decode(target)
	}

	cli := fetch.NewClient()
	backendReq, err := fetch.NewRequest(r.Context(), http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error creating new fetch request: %v", err)
		return err
	}

	backendReq.Header.Add("User-Agent", "jckli-worker")
	backendReq.Header.Add("Authorization", "token "+cloudflare.Getenv("GITHUB_TOKEN"))

	backendResp, err := cli.Do(backendReq, nil)
	if err != nil {
		log.Printf("Error performing fetch request: %v", err)
		return err
	}
	defer backendResp.Body.Close()

	bodyBytes, err := io.ReadAll(backendResp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return err
	}

	if err := json.Unmarshal(bodyBytes, target); err != nil {
		log.Printf("Error unmarshaling response body: %v", err)
		return err
	}

	cacheableResp := &http.Response{
		StatusCode: backendResp.StatusCode,
		Header:     backendResp.Header.Clone(),
		Body:       io.NopCloser(bytes.NewReader(bodyBytes)),
	}

	if cacheableResp.Header.Get("Cache-Control") == "" {
		cacheableResp.Header.Set("Cache-Control", fmt.Sprintf("max-age=%.0f", (1*time.Hour).Seconds()))
	}

	cloudflare.WaitUntil(func() {
		err := c.Put(cacheKeyReq, cacheableResp)
		if err != nil {
			log.Printf("Error putting response into cache: %v", err)
		}
	})

	return nil
}

func IsImageFile(p string) bool {
	ext := path.Ext(p)
	return ext == ".jpg" || ext == ".png" || ext == ".gif"
}
