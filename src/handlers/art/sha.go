package art

import (
	"art-api/src/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

var ErrArtNotFound = errors.New("art with the specified SHA was not found")

func BySHA(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sha := params.ByName("sha")
	if sha == "" {
		http.Error(w, "Bad Request: SHA parameter is missing.", http.StatusBadRequest)
		return
	}

	response, err := getArtBySHA(r, sha)
	if err != nil {
		log.Printf("Error getting art by SHA %s: %v", sha, err)
		if errors.Is(err, ErrArtNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func getArtBySHA(r *http.Request, sha string) (*ArtDetail, error) {
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

	for _, node := range githubResp.Tree {
		if node.Sha == sha && node.Kind == "blob" && utils.IsImageFile(node.Path) {
			escapedPath := url.PathEscape(node.Path)
			fullURL := RAW_CONTENT_BASE_URL + escapedPath

			return &ArtDetail{
				Url:    fullURL,
				Sha:    node.Sha,
				Path:   node.Path,
				Status: http.StatusOK,
			}, nil
		}
	}

	return nil, ErrArtNotFound
}
