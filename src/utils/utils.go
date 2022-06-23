package utils

import (
	"encoding/json"
	"net/http"
	"log"
	"os"
)

func RequestImages(url string, target interface{}) error {
	cli := &http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "worker")
	req.Header.Add("Authorization", "token " + os.Getenv("GITHUB_TOKEN"))
	if err1 != nil {
		log.Fatalln(err1)
	}
	resp, err2 := cli.Do(req)
	if err2 != nil {
		log.Fatalln(err2)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}