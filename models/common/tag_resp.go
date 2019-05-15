package common

type Tag struct {
	TagName string `json:"tag_name"`
	Count   int    `json:"count"`
}

type Key struct {
	Keywords string `json:"keywords"`
}