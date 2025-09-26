package art

import (
	"art-api/src/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func All(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	orientation := r.URL.Query().Get("o")
	if _, isValid := validOrientations[orientation]; !isValid {
		orientation = "any"
	}

	response, err := getArt(r, orientation)
	if err != nil {
		log.Printf("Failed to get art: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func getArt(r *http.Request, orientation string) (*AllArt, error) {
	var githubResp GithubTree
	err := utils.RequestArtJson(
		r,
		ART_KV_BINDING,
		GITHUB_ART_TREE_CACHE_KEY,
		GITHUB_API_URL,
		CACHE_TTL,
		&githubResp,
	)
	if err != nil {
		return nil, err
	}

	artData := make([]AllArtData, 0, len(githubResp.Tree)/2)

	for _, node := range githubResp.Tree {
		if node.Kind != "blob" || !utils.IsImageFile(node.Path) {
			continue
		}

		if orientation != "any" && !strings.Contains(node.Path, "/"+orientation+"/") {
			continue
		}

		escapedPath := url.PathEscape(node.Path)
		fullURL := RAW_CONTENT_BASE_URL + escapedPath

		artData = append(artData, AllArtData{Url: fullURL, Sha: node.Sha})
	}

	return &AllArt{
		Data:        artData,
		Status:      http.StatusOK,
		Sha:         githubResp.Sha,
		Orientation: orientation,
	}, nil
}
