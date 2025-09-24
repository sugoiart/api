package art

import (
	"art-api/src/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func Stats(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stats, err := getArtStats(r)
	if err != nil {
		log.Printf("Error getting art stats: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func getArtStats(r *http.Request) (*ArtStats, error) {
	var githubResp GithubTree
	if err := utils.RequestImages(GITHUB_API_URL, &githubResp, r); err != nil {
		return nil, err
	}

	orientationCounts := make(map[string]int)
	for orientation := range validOrientations {
		orientationCounts[orientation] = 0
	}
	orientationCounts["other"] = 0

	stats := &ArtStats{
		LastUpdateSha: githubResp.Sha,
		Orientations:  orientationCounts,
	}

	for _, node := range githubResp.Tree {
		if node.Kind != "blob" || !utils.IsImageFile(node.Path) {
			continue
		}

		stats.TotalImages++

		categorized := false
		for orientation := range validOrientations {
			if strings.Contains(node.Path, "/"+orientation+"/") {
				stats.Orientations[orientation]++
				categorized = true
				break
			}
		}

		if !categorized {
			stats.Orientations["other"]++
		}
	}

	return stats, nil
}
