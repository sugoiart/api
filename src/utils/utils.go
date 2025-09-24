package utils

import (
	"encoding/json"
	"github.com/syumai/workers/cloudflare"
	"github.com/syumai/workers/cloudflare/fetch"
	"log"
	"net/http"
)

func RequestImages(url string, target interface{}, r *http.Request) error {
	cli := fetch.NewClient()
	req, err1 := fetch.NewRequest(r.Context(), http.MethodGet, url, nil)
	req.Header.Add("User-Agent", "jckli-worker")
	req.Header.Add("Authorization", "token "+cloudflare.Getenv("GITHUB_TOKEN"))
	if err1 != nil {
		log.Fatalln(err1)
	}
	resp, err2 := cli.Do(req, nil)
	if err2 != nil {
		log.Fatalln(err2)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}
