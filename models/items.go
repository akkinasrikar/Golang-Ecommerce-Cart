package models

type ImageResize struct {
	Id         int    `json:"id"`
	ImageBytes []byte `json:"image_bytes"`
}
