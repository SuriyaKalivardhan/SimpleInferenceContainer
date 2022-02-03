package api

type InferenceRequest struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Input string `json:"input"`
}

type InferenceResponse struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Output string `json:"output"`
}
