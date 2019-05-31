package common

type Tag struct {
	TagName string `json:"tag_name"`
	Count   int    `json:"count"`
}

type Key struct {
	Keywords string `json:"keywords"`
}

type SearchKey struct {
	BlogTitle string `json:"blog_title"`
	Keywords  string `json:"keywords"`
	Author    string `json:"author"`
	TypeName  string `json:"type_name"`
}
