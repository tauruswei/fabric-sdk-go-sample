package model

type FileBasic struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Uploader string `json:"uploader"`
}
