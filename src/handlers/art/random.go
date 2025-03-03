package art

import (
	"art-api/src/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"net/http"
	"strings"
)

func Random(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	orien := r.URL.Query()["o"]
	var response *RandomArt
	if len(orien) > 0 {
		if orien[0] == "portrait" || orien[0] == "landscape" || orien[0] == "square" {
			response = RandomArtwork(orien[0])
		} else {
			response = RandomArtwork("")
		}
	} else {
		response = RandomArtwork("")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func RandomArtwork(orientation string) *RandomArt {
	var response *RandomArt
	githubresp := GithubTree{}
	utils.RequestImages(
		"https://api.github.com/repos/artmoe/art/git/trees/master?recursive=1",
		&githubresp,
	)

	var list []GithubTreeNode
	if orientation != "" {
		list = getArtOfOrientation(orientation)
	} else {
		list = githubresp.Tree
		orientation = "any"
	}

	for {
		random := list[rand.Intn(len(list))]
		if (random.Kind == "blob") &&
			(strings.HasSuffix(random.Path, ".jpg") || strings.HasSuffix(random.Path, ".png") || strings.HasSuffix(random.Path, ".gif")) {
			url := "https://raw.githubusercontent.com/artmoe/art/master/" + random.Path
			url = strings.ReplaceAll(url, " ", "%20")
			sha := random.Sha
			response = &RandomArt{Url: url, Status: 200, Sha: sha, Orientation: orientation}
			break
		}
	}

	return response
}

func getArtOfOrientation(orientation string) []GithubTreeNode {
	var list []GithubTreeNode
	githubresp := GithubTree{}
	utils.RequestImages(
		"https://api.github.com/repos/artmoe/art/git/trees/master?recursive=1",
		&githubresp,
	)

	for _, s := range githubresp.Tree {
		if strings.Contains(s.Path, "/"+orientation+"/") {
			if (s.Kind == "blob") &&
				(strings.HasSuffix(s.Path, ".jpg") || strings.HasSuffix(s.Path, ".png") || strings.HasSuffix(s.Path, ".gif")) {
				url := "https://raw.githubusercontent.com/artmoe/art/master/" + s.Path
				url = strings.ReplaceAll(url, " ", "%20")
				path := s.Path
				sha := s.Sha
				mode := s.Mode
				kind := s.Kind
				size := s.Size
				list = append(
					list,
					GithubTreeNode{
						Path: path,
						Mode: mode,
						Kind: kind,
						Sha:  sha,
						Size: size,
						Url:  url,
					},
				)
			}
		}
	}

	return list
}

