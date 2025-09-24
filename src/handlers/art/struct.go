package art

const (
	GITHUB_API_URL       = "https://api.github.com/repos/sugoiart/art/git/trees/master?recursive=1"
	RAW_CONTENT_BASE_URL = "https://raw.githubusercontent.com/sugoiart/art/master/"
)

var validOrientations = map[string]struct{}{
	"portrait":  {},
	"landscape": {},
	"square":    {},
}

type IndexResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RandomArt struct {
	Url         string `json:"url"`
	Sha         string `json:"sha"`
	Status      int    `json:"status"`
	Orientation string `json:"orientation"`
}

type GithubTree struct {
	Sha       string           `json:"sha"`
	Url       string           `json:"url"`
	Tree      []GithubTreeNode `json:"tree"`
	Truncated bool             `json:"truncated"`
}

type GithubTreeNode struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Kind string `json:"type"`
	Size int    `json:"size"`
	Sha  string `json:"sha"`
	Url  string `json:"url"`
}

type AllArt struct {
	Data        []AllArtData `json:"data"`
	Status      int          `json:"status"`
	Sha         string       `json:"sha"`
	Orientation string       `json:"orientation"`
}

type AllArtData struct {
	Url string `json:"url"`
	Sha string `json:"sha"`
}

type ArtDetail struct {
	Url    string `json:"url"`
	Sha    string `json:"sha"`
	Path   string `json:"path"`
	Status int    `json:"status"`
}

type ArtStats struct {
	TotalImages   int            `json:"totalImages"`
	LastUpdateSha string         `json:"lastUpdateSha"`
	Orientations  map[string]int `json:"orientations"`
}
