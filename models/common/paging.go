package common

type PageResult struct {
	TotalNum int         `json:"totalNum"`
	List     interface{} `json:"list"`
}

type PageRequest struct {
	CurrentPage int    `json:"current_page"`
	PageSize    int    `json:"page_size"`
	RoleId      int    `json:"role_id"`
	SearchWords string `json:"search_words"`
}
