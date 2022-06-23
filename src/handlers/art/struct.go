package art

type RandomArt struct {
	Url string `json:"url"`
	Sha string `json:"sha"`
	Status int `json:"status"`
}

type GithubTree struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
	Tree []GithubTreeNode `json:"tree"`
	Truncated bool `json:"truncated"`
}

type GithubTreeNode struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Size int `json:"size"`
	Sha string `json:"sha"`
	Url string `json:"url"`
}

type AllArt struct {
	Data []AllArtData `json:"data"`
	Status int `json:"status"`
	Sha string `json:"sha"`
}

type AllArtData struct {
	Url string `json:"url"`
	Sha string `json:"sha"`
}