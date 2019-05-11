package common

type PageResult struct {
	TotalNum int         `json:"totalNum"`
	List     interface{} `json:"list"`
}

type PageRequest struct {
	CurrentPage int    `json:"current_page"`
	PageSize    int    `json:"page_size"`
	SearchWords string `json:"search_words"`
}

type UserPageRequest struct {
	PageRequest
	RoleId int `json:"role_id"`
}

type BlogPageRequest struct {
	PageRequest
	BlogTypeId int `json:"blog_type_id"`
}

type CommentPageRequest struct {
	PageRequest
	BlogId int `json:"blog_id"`
}
