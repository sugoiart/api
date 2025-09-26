package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/syumai/workers/cloudflare"
	"github.com/syumai/workers/cloudflare/fetch"
	"github.com/syumai/workers/cloudflare/kv"
)

func RequestArtJson(r *http.Request, kvBinding, cacheKey, url string, ttl time.Duration, target interface{}) error {
	store, err := kv.NewNamespace(kvBinding)
	if err != nil {
		return fmt.Errorf("failed to open KV namespace '%s': %w", kvBinding, err)
	}

	cachedString, err := store.GetString(cacheKey, nil)
	if err == nil && cachedString != "" {
		return json.Unmarshal([]byte(cachedString), target)
	}

	log.Printf("KV Miss for key '%s'. Fetching fresh data from %s.", cacheKey, url)
	if err := fetchFromGithub(r.Context(), url, target); err != nil {
		return fmt.Errorf("failed to fetch from origin: %w", err)
	}

	freshBytes, err := json.Marshal(target)
	if err != nil {
		log.Printf("Failed to marshal fresh data for KV store (key: %s): %v", cacheKey, err)
		return nil
	}

	cloudflare.WaitUntil(func() {
		opts := &kv.PutOptions{
			ExpirationTTL: int(ttl.Seconds()),
		}
		err := store.PutString(cacheKey, string(freshBytes), opts)
		if err != nil {
			log.Printf("Failed to put data into KV (key: %s): %v", cacheKey, err)
		}
	})

	return nil
}

func fetchFromGithub(ctx context.Context, url string, target interface{}) error {
	cli := fetch.NewClient()
	req, err := fetch.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "jckli-worker")
	req.Header.Add("Authorization", "token "+cloudflare.Getenv("GITHUB_TOKEN"))

	resp, err := cli.Do(req, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("github API returned non-200 status: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func IsImageFile(p string) bool {
	ext := path.Ext(p)
	return ext == ".jpg" || ext == ".png" || ext == ".gif"
}
