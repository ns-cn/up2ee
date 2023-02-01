package main

type AccessData struct {
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	CreateAt     int    `json:"created_at,omitempty"`
}

type CreateFileData struct {
	Content CreateFileContentData `json:"content,omitempty"`
}

type CreateFileContentData struct {
	DownloadUrl string `json:"download_url,omitempty"`
}
