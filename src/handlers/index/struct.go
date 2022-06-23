package index

type IndexResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

type IndexRoute struct {
	Path string `json:"path"`
	Method string `json:"method"`
	Url string `json:"url"`
}