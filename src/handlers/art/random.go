package art

import (
	"art-api/src/utils"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

var ErrNoImagesFound = errors.New("no images found matching the criteria")

func Random(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	orientation := r.URL.Query().Get("o")
	if _, isValid := validOrientations[orientation]; !isValid {
		orientation = "any"
	}

	response, err := getRandomArt(r, orientation)
	if err != nil {
		log.Printf("Error getting random art: %v", err)
		if errors.Is(err, ErrNoImagesFound) {
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

func getRandomArt(r *http.Request, orientation string) (*RandomArt, error) {
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

	imageNodes := make([]GithubTreeNode, 0)
	for _, node := range githubResp.Tree {
		if node.Kind != "blob" || !utils.IsImageFile(node.Path) {
			continue
		}
		if orientation != "any" && !strings.Contains(node.Path, "/"+orientation+"/") {
			continue
		}
		imageNodes = append(imageNodes, node)
	}

	if len(imageNodes) == 0 {
		return nil, ErrNoImagesFound
	}

	randomNode := imageNodes[rand.Intn(len(imageNodes))]

	escapedPath := url.PathEscape(randomNode.Path)
	fullURL := RAW_CONTENT_BASE_URL + escapedPath

	return &RandomArt{
		Url:         fullURL,
		Status:      http.StatusOK,
		Sha:         randomNode.Sha,
		Orientation: orientation,
	}, nil
}
