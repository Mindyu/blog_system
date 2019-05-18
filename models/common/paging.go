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
	BlogTypeId int    `json:"blog_type_id"`
	SortType   int    `json:"sort_type"`
	Author     string `json:"author"`
}

type CommentPageRequest struct {
	PageRequest
	BlogId int `json:"blog_id"`
	Author     string `json:"author"`
}

type ReplyPageRequest struct {
	PageRequest
	CommentId int `json:"comment_id"`
}

type LogPageRequest struct {
	PageRequest
	UserName  string `json:"user_name"`
	CallApi   string `json:"call_api"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type RelationPageRequest struct {
	PageRequest
	Username string `json:"username"`
}
