package art

import (
	"art-api/src/utils"
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"github.com/julienschmidt/httprouter"
)

func Random(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var response *RandomArt
	githubresp := GithubTree{}
	utils.RequestImages("https://api.github.com/repos/artmoe/art/git/trees/master?recursive=1", &githubresp)
	for {
		random := githubresp.Tree[rand.Intn(len(githubresp.Tree))]
		if (random.Type == "blob") && (strings.HasSuffix(random.Path, ".jpg") || strings.HasSuffix(random.Path, ".png") || strings.HasSuffix(random.Path, ".gif")) {
			url := "https://raw.githubusercontent.com/artmoe/art/master/" + random.Path
			url = strings.ReplaceAll(url, " ", "%20")
			sha := random.Sha
			response = &RandomArt{Url: url, Status: 200, Sha: sha}
			break
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}