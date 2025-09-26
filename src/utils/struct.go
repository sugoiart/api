package utils

type ArtDetail struct {
	URL         string `json:"url"`
	Sha         string `json:"sha"`
	Path        string `json:"path"`
	Orientation string `json:"orientation"`
}

type ArtStats struct {
	TotalImages   int            `json:"totalImages"`
	LastUpdateSha string         `json:"lastUpdateSha"`
	Orientations  map[string]int `json:"orientations"`
}

type ProcessedArtCache struct {
	ImagesBySHA         map[string]ArtDetail   `json:"imagesBySHA"`
	ImagesByOrientation map[string][]ArtDetail `json:"imagesByOrientation"`
	Stats               ArtStats               `json:"stats"`
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
